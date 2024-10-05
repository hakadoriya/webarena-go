package indigo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1VmInstanceSpec struct {
	ID              int64                          `json:"id"`
	Name            string                         `json:"name"`
	Description     string                         `json:"description"`
	UsePossibleDate string                         `json:"use_possible_date"` //nolint:tagliatelle // JSON field name is defined by the API
	InstanceTypeID  int64                          `json:"instancetype_id"`   //nolint:tagliatelle // JSON field name is defined by the API
	CreatedAt       string                         `json:"created_at"`        //nolint:tagliatelle // JSON field name is defined by the API
	UpdatedAt       string                         `json:"updated_at"`        //nolint:tagliatelle // JSON field name is defined by the API
	InstanceType    WebArenaIndigoV1VmInstanceType `json:"instance_type"`     //nolint:tagliatelle // JSON field name is defined by the API
}

// Get Instance Specification
// https://indigo.arena.ne.jp/userapi/#get_instance_specification
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/getinstancespec?instanceTypeId=1&osId=1 \
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
//	    "speclist": [
//	        {
//	            "id": 1,
//	            "name": "2 CPU & 2 GB RAM plan",
//	            "description": "2 CPU & 2 GB RAM plan",
//	            "use_possible_date": "2019-01-03 08:50:00",
//	            "instancetype_id": 1,
//	            "created_at": "2019-01-04 08:40:57",
//	            "updated_at": "2019-01-04 08:40:57",
//	            "instance_type": {
//	                "id": 1,
//	                "name": "instance",
//	                "display_name": "KVM Instance",
//	                "created_at": "2019-02-12 22:46:35",
//	                "updated_at": "2019-02-12 22:46:35"
//	            }
//	        }
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmInstanceSpec(ctx context.Context, instanceTypeID, osID int64) (*GetWebArenaIndigoV1VmInstanceSpecResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	urlValues := url.Values{}
	urlValues.Add("instanceTypeId", strconv.FormatInt(instanceTypeID, 10)) //nolint:staticcheck
	urlValues.Add("osId", strconv.FormatInt(osID, 10))                     //nolint:staticcheck

	httpReq, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("%s?%s", PathWebArenaIndigoV1VmInstanceSpec, urlValues.Encode()), nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1VmInstanceSpecResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1VmInstanceSpecResponse struct {
	Success  bool                             `json:"success"`
	Total    int64                            `json:"total"`
	SpecList []WebArenaIndigoV1VmInstanceSpec `json:"speclist"`
}
