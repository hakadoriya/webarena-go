package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Assign Firewall
// https://indigo.arena.ne.jp/userapi/#assign_firewall
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/nw/assign \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	      "instanceid":"4",
//	      "templateid":"8"
//	      }'
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
//	    "message": "Firewall template is assigned successfully.",
//	    "sucessCode": "F60003"
//	}
func (c *Client) PostWebArenaIndigoV1NwAssign(ctx context.Context, req *PostWebArenaIndigoV1NwAssignRequest) (*PostWebArenaIndigoV1NwAssignResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1NwAssign, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1NwAssignResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1NwAssignRequest struct {
	InstanceID int64 `json:"instanceid"`
	TemplateID int64 `json:"templateid"`
}

type PostWebArenaIndigoV1NwAssignResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	SucessCode string `json:"sucessCode"` // Typo in the API response
}
