package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Generate API Key
// https://indigo.arena.ne.jp/userapi/#api_key_generate
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/auth/create/apikey \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 201 Created
// RESPONSE BODY
//
//	{
//	    "apiKey": "m70QyrbMUZWl06SfSAvRBPQO0ofsadf",
//	    "apiSecret": "LmAY0pB1xA1fas"
//	}
func (c *Client) CreateWebArenaIndigoV1AuthCreateAPIKey(ctx context.Context) (*CreateWebArenaIndigoV1AuthCreateAPIKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1AuthCreateAPIKey, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp CreateWebArenaIndigoV1AuthCreateAPIKeyResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type CreateWebArenaIndigoV1AuthCreateAPIKeyResponse struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}
