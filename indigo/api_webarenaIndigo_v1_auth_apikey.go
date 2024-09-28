package indigo

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1AuthAPIKey struct {
	ID        int64  `json:"id"`
	APIKey    string `json:"apiKey"`
	CreatedAt string `json:"created_at"` //nolint:tagliatelle // JSON field name is defined by the API
}

// API Key List
// https://indigo.arena.ne.jp/userapi/#api_key_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/auth/apikey \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 Ok
// RESPONSE BODY
//
//	{
//	    "success": true,
//	    "total": 1,
//	    "accesstokens": [
//	        {
//	            "id": 434,
//	            "apiKey": "3PwzRzZyXAmBSi0NiYGQjUGpDfsadf",
//	            "created_at": "2019-10-21 11:17:08"
//	        }
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1AuthAPIKey(ctx context.Context) (*GetWebArenaIndigoV1AuthAPIKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1AuthAPIKey, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1AuthAPIKeyResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1AuthAPIKeyResponse struct {
	Success      bool                         `json:"success"`
	Total        int64                        `json:"total"`
	AccessTokens []WebArenaIndigoV1AuthAPIKey `json:"accesstokens"`
}

// Destroy API Key
// https://indigo.arena.ne.jp/userapi/#api_key_destroy
//
// CURL EXAMPLE
//
//	curl -X DELETE \
//	  https://api.customer.jp/webarenaIndigo/v1/auth/apikey/88 \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 OK
// RESPONSE BODY
//
//	{
//	    "success": true,
//	    "message": "API Key is removed successfully"
//	}
func (c *Client) DeleteWebArenaIndigoV1AuthAPIKey(ctx context.Context, apiKeyID int64) (*DeleteWebArenaIndigoV1AuthAPIKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodDelete, path.Join(PathWebArenaIndigoV1AuthAPIKey, strconv.FormatInt(apiKeyID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp DeleteWebArenaIndigoV1AuthAPIKeyResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type DeleteWebArenaIndigoV1AuthAPIKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
