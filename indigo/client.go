package indigo

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"

	"golang.org/x/time/rate"

	"github.com/kunitsucom/util.go/env"
	errorz "github.com/kunitsucom/util.go/errors"
	"github.com/kunitsucom/util.go/retry"
)

type (
	AccessToken struct {
		AccessToken string        `json:"accessToken"`
		TokenType   string        `json:"tokenType"`
		ExpiresIn   time.Duration `json:"expiresIn"`
		Scope       string        `json:"scope"`
		IssuedAt    time.Time     `json:"issuedAt"`
	}

	Client struct {
		debugLog     *log.Logger
		httpClient   *http.Client
		endpoint     string
		clientID     string
		clientSecret string
		// NOTE: The API returns the header `X-Quota-Allowed: 6`, so it is estimated that up to 6 requests per minute (1 request every 10 seconds) are allowed.
		//
		// Too Many Request Error Response Example:
		//
		//	HTTP/2.0 429 Too Many Requests
		//	Content-Length: 259
		//	Access-Control-Allow-Headers: origin, x-requested-with, accept, Authorization, Content-Type
		//	Access-Control-Allow-Methods: GET, PUT, POST, DELETE
		//	Access-Control-Allow-Origin: *
		//	Access-Control-Max-Age: 1728000
		//	Content-Type: application/json
		//	Date: Sun, 12 May 2024 13:41:52 GMT
		//	X-Quota-Allowed: 6
		//	X-Quota-Available: 0
		//	X-Quota-Reset: 1715521320000
		//	X-Request-Id: 00000000-0000-4000-0000-000000000000
		//
		//	{"errorCode": "429", "errorMessage": "Too Many Request.", "developerMessage": "Rate limit quota violation. Quota limit  exceeded. Identifier : ffffffff-ffff-4fff-ffff-ffffffffffff", "moreInfo": null, "requestId": "ffffffff-ffff-ffff-ffff-fffffffffffffffffff"}
		rateLimiter *rate.Limiter
		retryConfig *retry.Config
		accessToken *AccessToken
	}

	ClientOption interface {
		apply(c *Client)
	}
)

type debugLogOption struct{ debugLog *log.Logger }

func (o *debugLogOption) apply(c *Client) { c.debugLog = o.debugLog }

func ClientOptionWithDebugLog(debugLog *log.Logger) ClientOption { //nolint:ireturn
	return &debugLogOption{debugLog: debugLog}
}

type clientIDOption struct{ clientID string }

func (o *clientIDOption) apply(c *Client) {
	c.clientID = o.clientID
}

func ClientOptionWithClientID(clientID string) ClientOption { //nolint:ireturn
	return &clientIDOption{clientID: clientID}
}

type clientSecretOption struct{ clientSecret string }

func (o *clientSecretOption) apply(c *Client) {
	c.clientSecret = o.clientSecret
}

func ClientOptionWithClientSecret(clientSecret string) ClientOption { //nolint:ireturn
	return &clientSecretOption{clientSecret: clientSecret}
}

type httpClientOption struct{ httpClient *http.Client }

func (o *httpClientOption) apply(c *Client) { c.httpClient = o.httpClient }

func ClientOptionWithHTTPClient(httpClient *http.Client) ClientOption { //nolint:ireturn
	return &httpClientOption{httpClient: httpClient}
}

type endpointOption struct{ endpoint string }

func (o *endpointOption) apply(c *Client) { c.endpoint = o.endpoint }

func ClientOptionWithEndpoint(endpoint string) ClientOption { //nolint:ireturn
	return &endpointOption{endpoint: endpoint}
}

type rateLimiterOption struct{ rateLimiter *rate.Limiter }

func (o *rateLimiterOption) apply(c *Client) {
	c.rateLimiter = o.rateLimiter
}

func ClientOptionWithoutRateLimiter() ClientOption { //nolint:ireturn
	return &rateLimiterOption{rateLimiter: rate.NewLimiter(rate.Inf, 0)}
}

func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	// rate limit
	const (
		defaultRateLimitInterval = 10 * time.Second
		defaultRateLimitBurst    = 1
	)

	// retry
	const (
		defaultInitialRetryInterval = 1 * time.Second
		defaultMaxRetryInterval     = 10 * time.Second
	)

	c := &Client{
		debugLog:     log.New(io.Discard, "", log.LstdFlags),
		httpClient:   http.DefaultClient,
		endpoint:     env.StringOrDefault(WEBARENA_INDIGO_ENDPOINT, "https://api.customer.jp"),
		clientID:     os.Getenv(WEBARENA_INDIGO_CLIENT_ID),
		clientSecret: os.Getenv(WEBARENA_INDIGO_CLIENT_SECRET),
		rateLimiter:  rate.NewLimiter(rate.Every(defaultRateLimitInterval), defaultRateLimitBurst),
		retryConfig:  retry.NewConfig(defaultInitialRetryInterval, defaultMaxRetryInterval),
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	if c.clientID == "" || c.clientSecret == "" {
		return nil, errorz.Errorf("clientId or clientSecret is empty: %w", ErrInvalidClientCredentials)
	}

	accessToken, err := c.IssueAccessToken(ctx)
	if err != nil {
		return nil, errorz.Errorf("c.IssueAccessToken: %w", err)
	}
	c.accessToken = accessToken

	return c, nil
}

func (c *Client) newRequest(ctx context.Context, method string, urlSuffix string, reqBody []byte) (*http.Request, error) {
	url := c.endpoint + urlSuffix

	var body io.Reader = http.NoBody
	if len(reqBody) > 0 {
		body = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, errorz.Errorf("http.NewRequest: %w", err)
	}

	return req, nil
}

func (c *Client) IssueAccessToken(ctx context.Context) (*AccessToken, error) {
	ctx, span := start(ctx)
	defer span.End()

	req := &PostOAuthV1AccessTokensRequest{
		GrantType:    "client_credentials", // default
		ClientId:     c.clientID,
		ClientSecret: c.clientSecret,
		Code:         "",
	}

	resp, err := c.PostOAuthV1AccessTokens(ctx, req)
	if err != nil {
		return nil, errorz.Errorf("c.PostOAuthV1AccessTokens: %w", err)
	}

	accessToken, err := c.convertAuthResponseToAccessToken(resp)
	if err != nil {
		return nil, errorz.Errorf("c.convertResponseToAccessToken: %w", err)
	}

	return accessToken, nil
}

func (c *Client) convertAuthResponseToAccessToken(resp *PostOAuthV1AccessTokensResponse) (*AccessToken, error) {
	issuedAtUnixMilli, err := strconv.ParseInt(resp.IssuedAt, 10, 64) //nolint:mnd
	if err != nil {
		return nil, errorz.Errorf("strconv.ParseInt: IssuedAt=%s: %w", resp.IssuedAt, err)
	}
	expiresIn, err := strconv.ParseInt(resp.ExpiresIn, 10, 64) //nolint:mnd
	if err != nil {
		return nil, errorz.Errorf("strconv.ParseInt: ExpiresIn=%s: %w", resp.ExpiresIn, err)
	}

	return &AccessToken{
		AccessToken: resp.AccessToken,
		TokenType:   resp.TokenType,
		ExpiresIn:   time.Duration(expiresIn) * time.Second,
		Scope:       resp.Scope,
		IssuedAt:    time.UnixMilli(issuedAtUnixMilli),
	}, nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	if c.accessToken == nil || time.Now().After(c.accessToken.IssuedAt.Add(c.accessToken.ExpiresIn)) {
		accessToken, err := c.IssueAccessToken(req.Context())
		if err != nil {
			return nil, errorz.Errorf("c.issueAccessToken: %w", err)
		}
		c.accessToken = accessToken
	}
	req.Header.Set("Authorization", "Bearer "+c.accessToken.AccessToken)

	return c.doRequestWithoutAccessToken(req)
}

//nolint:cyclop
func (c *Client) doRequestWithoutAccessToken(req *http.Request) (*http.Response, error) {
	var out *http.Response

	f := func(ctx context.Context) (err error) {
		var dumpReq, dumpResp []byte
		defer func() { c.debugLog.Printf("\n[REQUEST]\n%s\n[RESPONSE]\n%s", dumpReq, dumpResp) }()

		if err := c.rateLimiter.Wait(ctx); err != nil {
			return errorz.Errorf("c.rateLimiter.Wait: %w", err)
		}

		dumpReq, err = httputil.DumpRequest(req, true)
		if err != nil {
			return errorz.Errorf("httputil.DumpRequest: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return errorz.Errorf("c.httpClient.Do: %w", err)
		}
		defer func() {
			if err != nil {
				_ = resp.Body.Close()
			}
		}()

		dumpResp, err = httputil.DumpResponse(resp, true)
		if err != nil {
			return errorz.Errorf("httputil.DumpResponse: %w", err)
		}

		switch {
		case resp.StatusCode == http.StatusTooManyRequests:
			return errorz.Errorf("method=%s url=%s code=%d body=%s: %w", req.Method, req.URL, resp.StatusCode, getLimitedBody(resp.Body), ErrAPIReturnsTooManyRequest)
		case resp.StatusCode == http.StatusUnauthorized:
			return errorz.Errorf("method=%s url=%s code=%d body=%s: %w", req.Method, req.URL, resp.StatusCode, getLimitedBody(resp.Body), ErrAPIReturnsUnauthorized)
		case resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode:
			return errorz.Errorf("method=%s url=%s code=%d body=%s: %w", req.Method, req.URL, resp.StatusCode, getLimitedBody(resp.Body), ErrUnexpectedStatusCode)
		}

		out = resp
		return nil
	}

	retryer := retry.New(req.Context(), c.retryConfig)
	if err := retryer.Do(f, retry.WithRetryableErrors(ErrAPIReturnsTooManyRequest)); err != nil {
		return nil, errorz.Errorf("(*retry.Retryer) Do: %w", err)
	}

	return out, nil
}

func getLimitedBody(r io.Reader) string {
	const limitedBodyLength = 512
	b, _ := io.ReadAll(io.LimitReader(r, limitedBodyLength))
	return string(b)
}
