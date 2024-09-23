package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1NwFirewallRule struct {
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Source   string `json:"source"`
}

// Create Firewall
// https://indigo.arena.ne.jp/userapi/#firewall_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/nw/createfirewall \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	    "name":"Example",
//	    "inbound":[
//	        {"type":"HTTP","protocol":"TCP","port":"80","source":"0.0.0.0"},
//	        {"type":"HTTPS","protocol":"TCP","port":"443","source":"0.0.0.0"}
//	    ],
//	    "outbound":[
//	        {"type":"HTTP","protocol":"TCP","port":"80","source":"0.0.0.0"},
//	        {"type":"HTTPS","protocol":"TCP","port":"443","source":"0.0.0.0"}
//	    ],
//	    "instances":["6","5"]
//	}'
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
//	    "message": "Firewall template has been created successfully.",
//	    "sucessCode": "F60002",
//	    "firewallId": 55
//	}
func (c *Client) PostWebArenaIndigoV1NwCreateFirewall(ctx context.Context, req *PostWebArenaIndigoV1NwCreateFirewallRequest) (*PostWebArenaIndigoV1NwCreateFirewallResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1NwCreateFirewall, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &PostWebArenaIndigoV1NwCreateFirewallResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type PostWebArenaIndigoV1NwCreateFirewallRequest struct {
	Name      string                           `json:"name"`
	Inbound   []WebArenaIndigoV1NwFirewallRule `json:"inbound"`
	Outbound  []WebArenaIndigoV1NwFirewallRule `json:"outbound"`
	Instances []string                         `json:"instances"`
}

type PostWebArenaIndigoV1NwCreateFirewallResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	SuccessCode string `json:"sucessCode"`
	FirewallID  int64  `json:"firewallId"`
}
