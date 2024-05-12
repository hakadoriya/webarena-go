package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	errorz "github.com/kunitsucom/util.go/errors"
)

type PostOAuthV1AccessTokensRequest struct {
	GrantType    string `json:"grantType"`
	ClientId     string `json:"clientId"` //nolint:revive,stylecheck
	ClientSecret string `json:"clientSecret"`
	Code         string `json:"code"`
}

type PostOAuthV1AccessTokensResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   string `json:"expiresIn"`
	Scope       string `json:"scope"`
	IssuedAt    string `json:"issuedAt"`
}

// Access Token Generation
// https://indigo.arena.ne.jp/userapi/#generate_access_token
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/oauth/v1/accesstokens\
//	  -H 'Content-Type: application/json' \
//	  -d '{
//	    "grantType": "client_credentials",
//	    "clientId": "YNG2yIyLDA3TqODlEXvLRwL7HzBjDsCQ",
//	    "clientSecret": "uDczeZJlxYvuzqcU",
//	    "code": ""
//	}'
//
// REQUEST HEADERS
// Content-Type: application/json.
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 201 Created
// RESPONSE BODY
//
//	{
//	    "accessToken": "HnGma03YnFPIF4DMttywiSOCGUHR",
//	    "tokenType": "BearerToken",
//	    "expiresIn": "3599",
//	    "scope": "",
//	    "issuedAt": "1550570350202"
//	}
func (c *Client) PostOAuthV1AccessTokens(ctx context.Context, req *PostOAuthV1AccessTokensRequest) (*PostOAuthV1AccessTokensResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathOAuthV1AccessTokens, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequestWithoutAccessToken(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequestWithoutAccessToken: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &PostOAuthV1AccessTokensResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}
