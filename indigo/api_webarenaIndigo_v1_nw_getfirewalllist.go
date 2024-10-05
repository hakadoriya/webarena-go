package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1NwFirewall struct {
	ID        int64  `json:"id"`
	ServiceID string `json:"service_id"` //nolint:tagliatelle // JSON field name is defined by the API
	UserID    int64  `json:"user_id"`    //nolint:tagliatelle // JSON field name is defined by the API
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"created_at"` //nolint:tagliatelle // JSON field name is defined by the API
	UpdatedAt string `json:"updated_at"` //nolint:tagliatelle // JSON field name is defined by the API
}

// Get Firewall list
// https://indigo.arena.ne.jp/userapi/#firewall_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/nw/getfirewalllist \
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
//	[
//	  {
//	     "id": 55,
//	     "service_id": "wsi-00003",
//	     "user_id": 6,
//	     "name": "Example",
//	     "status": 1,
//	     "created_at": "2018-11-14 10:35:33",
//	     "updated_at": "2018-11-14 10:35:33"
//	  },
//	  {
//	     "id": 47,
//	     "service_id": "wsi-00003",
//	     "user_id": 6,
//	     "name": "Example2",
//	     "status": 1,
//	     "created_at": "2018-10-30 10:07:46",
//	     "updated_at": "2018-11-05 04:54:15"
//	  }
//	]
func (c *Client) GetWebArenaIndigoV1NwGetFirewallList(ctx context.Context) (*GetWebArenaIndigoV1NwGetFirewallListResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1NwGetFirewallList, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1NwGetFirewallListResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1NwGetFirewallListResponse []WebArenaIndigoV1NwFirewall
