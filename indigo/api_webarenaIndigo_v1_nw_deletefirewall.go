package indigo

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

// Delete Firewall
// https://indigo.arena.ne.jp/userapi/#delete_firewall
//
// CURL EXAMPLE
// curl -X DELETE \
//   https://api.customer.jp/webarenaIndigo/v1/nw/deletefirewall/30 \
//   -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \

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
//	    "message": "Firewall template has been deleted successfully.",
//	    "sucessCode": "F6005"
//	}
func (c *Client) DeleteWebArenaIndigoV1NwDeleteFirewall(ctx context.Context, firewallID int64) (*DeleteWebArenaIndigoV1NwDeleteFirewallResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodDelete, path.Join(PathWebArenaIndigoV1NwDeleteFirewall, strconv.FormatInt(firewallID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &DeleteWebArenaIndigoV1NwDeleteFirewallResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type DeleteWebArenaIndigoV1NwDeleteFirewallResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	SucessCode string `json:"sucessCode"`
}
