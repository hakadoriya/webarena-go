package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	errorz "github.com/kunitsucom/util.go/errors"
)

type WebArenaIndigoV1VmInstanceType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"` //nolint:tagliatelle // JSON field name is defined by the API
	CreatedAt   string `json:"created_at"`   //nolint:tagliatelle // JSON field name is defined by the API
	UpdatedAt   string `json:"updated_at"`   //nolint:tagliatelle // JSON field name is defined by the API
}

// Instance Type List
// https://indigo.arena.ne.jp/userapi/#get_instance_type
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/instancetypes \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR'
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
//	    "total": 1,
//	    "instanceTypes": [
//	        {
//	            "id": 1,
//	            "name": "instance",
//	            "display_name": "KVM Instance",
//	            "created_at": "2019-02-12 22:46:35",
//	            "updated_at": "2019-02-12 22:46:35"
//	        }
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmInstanceTypes(ctx context.Context) (*GetWebArenaIndigoV1VmInstanceTypesResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1VmInstanceTypes, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &GetWebArenaIndigoV1VmInstanceTypesResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type GetWebArenaIndigoV1VmInstanceTypesResponse struct {
	Success       bool                             `json:"success"`
	Total         int                              `json:"total"`
	InstanceTypes []WebArenaIndigoV1VmInstanceType `json:"instanceTypes"`
}
