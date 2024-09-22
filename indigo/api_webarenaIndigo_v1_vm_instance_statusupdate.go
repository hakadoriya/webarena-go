package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Update Instance Status
// https://indigo.arena.ne.jp/userapi/#instance_status_update
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/instance/statusupdate \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{"instanceId":"16","status":"stop"}'
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
//	    "message": "Instance has stopped successfully ",
//	    "sucessCode": "I20009",
//	    "instanceStatus": "shutoff"
//	}
func (c *Client) PostWebArenaIndigoV1VmInstanceStatusUpdate(ctx context.Context, req *PostWebArenaIndigoV1VmInstanceStatusUpdateRequest) (*PostWebArenaIndigoV1VmInstanceStatusUpdateResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmInstanceStatusUpdate, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &PostWebArenaIndigoV1VmInstanceStatusUpdateResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type PostWebArenaIndigoV1VmInstanceStatusUpdateRequest struct {
	InstanceID string `json:"instanceId"`
	Status     string `json:"status"`
}

type PostWebArenaIndigoV1VmInstanceStatusUpdateResponse struct {
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	SuccessCode    string `json:"sucessCode"`
	InstanceStatus string `json:"instanceStatus"`
}
