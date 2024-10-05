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

type WebArenaIndigoV1VmOS struct {
	ID             int64  `json:"id"`
	CategoryID     int64  `json:"categoryid"`
	Name           string `json:"name"`
	ViewName       string `json:"viewname"`
	InstanceTypeID int64  `json:"instancetype_id"` //nolint:tagliatelle // JSON field name is defined by the API
}

// Get OS list
// https://indigo.arena.ne.jp/userapi/#get_os_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/oslist?instanceTypeId=1 \
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
//	    "osCategory": [
//	        {
//	            "id": 1,
//	            "name": "Ubuntu",
//	            "logo": "Ubudu.png",
//	            "osLists": [
//	                {
//	                    "id": 1,
//	                    "categoryid": 1,
//	                    "name": "Ubuntu18.04",
//	                    "viewname": "Ubuntu 18.04",
//	                    "instancetype_id": 1
//	                }
//	            ]
//	        }
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmOSList(ctx context.Context, instanceTypeID int64) (*GetWebArenaIndigoV1VmOsListResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	urlValues := url.Values{}
	urlValues.Add("instanceTypeId", strconv.FormatInt(instanceTypeID, 10)) //nolint:staticcheck

	httpReq, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("%s?%s", PathWebArenaIndigoV1VmOSList, urlValues.Encode()), nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1VmOsListResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1VmOsListResponse struct {
	Success    bool                             `json:"success"`
	Total      int64                            `json:"total"`
	OsCategory []WebArenaIndigoV1VmInstanceSpec `json:"osCategory"`
}
