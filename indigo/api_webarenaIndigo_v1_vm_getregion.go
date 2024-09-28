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

type WebArenaIndigoV1VmRegion struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	UsePossibleDate string `json:"use_possible_date"` //nolint:tagliatelle // JSON field name is defined by the API
}

// Get region list
// https://indigo.arena.ne.jp/userapi/#get_region_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/getregion?instanceTypeId=1 \
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
//	    "total": 2,
//	    "regionlist": [
//	        {
//	            "id": 1,
//	            "name": "Tokyo",
//	            "use_possible_date": "2018-09-30 12:00:00"
//	        },
//	        {
//	            "id": 2,
//	            "name": "Tokyo1",
//	            "use_possible_date": "2018-12-09 00:00:00"
//	        }
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmGetRegion(ctx context.Context, instanceTypeID int64) (*GetWebArenaIndigoV1VmGetRegionResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	urlValues := url.Values{}
	urlValues.Add("instanceTypeId", strconv.FormatInt(instanceTypeID, 10)) //nolint:staticcheck

	httpReq, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("%s?%s", PathWebArenaIndigoV1VmInstanceType, urlValues.Encode()), nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1VmGetRegionResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1VmGetRegionResponse struct {
	Success    bool                       `json:"success"`
	Total      int64                      `json:"total"`
	RegionList []WebArenaIndigoV1VmRegion `json:"regionlist"`
}
