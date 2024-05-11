package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	errorz "github.com/kunitsucom/util.go/errors"
)

// Update Firewall
// https://indigo.arena.ne.jp/userapi/#update_firewall
//
// CURL EXAMPLE
//
//	curl -X PUT \
//	  https://api.customer.jp/webarenaIndigo/v1/nw/updatefirewall \
//	  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjNiZjI5MWU4NmMyMTY4ZTMxN2U2OTViMGFmYjI3MzBjYzUwZWFiYjg4MzIwZmQzMTcwY2FhYjE5ZDRjNzZmMDlkYzAyZWI4YTU1ZTYyN2Q0In0.eyJhdWQiOiJjbGllbnQxIiwianRpIjoiM2JmMjkxZTg2YzIxNjhlMzE3ZTY5NWIwYWZiMjczMGNjNTBlYWJiODgzMjBmZDMxNzBjYWFiMTlkNGM3NmYwOWRjMDJlYjhhNTVlNjI3ZDQiLCJpYXQiOjE1NDE1NjY2MjIsIm5iZiI6MTU0MTU2NjYyMiwiZXhwIjoxNTQxNjUzMDIyLCJzdWIiOiIxNzEiLCJzY29wZXMiOltdfQ.efsYC4sJx1Arms5qwuBf5tbBjGxJ1Ep_p6XYSf9ugO8PiZECmPcl4p3MaqPs_Loxz_EcQODGqIKSF3w0KBpf-gC9vGD4G_4pikGgpsRcmvsvJ64X5K1p_WOI_v_REUo3adW7Gu4QXyDeEjiAGesl23dSi0C4EZHQ2x7L7joGAzkS0cWASWrM7H3cOKIlXDVtqh8fdykSkxXghx504SSiDzObXt28VnFY4fRvqKawe2-vMvG5EmZgAejriyPIS8nbrd2gnlLCaBTD96u4qz05w4MRtgPPWSfQ8e6sYMavXGau8D60KSyTstJyVV3AKX5GY5ESD15Ikt4tSIWlk8BbdNovnkGx8oKv_rkDEY1thWu_40Jgf7plUQ9ox_kf1kbhhj-YXSxfpQFT_tLFIaM3Ntb-FhfFlaw37OhC93B79Bq0J3RbQO7qw3hPKaIb0AdTL_iYjU2IKUUZ6Sr--GuZiqWW5Ls-L4OTkl7NyNc0G8ZJk5Pibavgp91B1Do5jpVS43IsuQGKmVfLRqJc-mR5RR3FmbKaGvj3ULvxN9kYYeP2e98bcHIBmkWhMgkuaZqUFrgTdq6a9fc0vhUZTmNVINGjCA6BLCixTrQfZ_eAtIam2jW4TGZB0yuQuLNa-yKx_Sop6rwVv-y9U7H-gvXE5G-K7lIXd0-Lz-sROhKMwtA' \
//	  -d '{
//	    "templateid":"55",
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
//	    "message": "Firewall template is updated successfully.",
//	    "sucessCode": "F6004",
//	    "firewallId": 55
//	}
func (c *Client) UpdateWebArenaIndigoV1NwFirewall(ctx context.Context, req *UpdateWebArenaIndigoV1NwFirewallRequest) (*UpdateWebArenaIndigoV1NwFirewallResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPut, PathWebArenaIndigoV1NwUpdateFirewall, body) // NOTE: PUT
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &UpdateWebArenaIndigoV1NwFirewallResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type UpdateWebArenaIndigoV1NwFirewallRequest struct {
	TemplateID int64                            `json:"templateid"`
	Name       string                           `json:"name"`
	Inbound    []WebArenaIndigoV1NwFirewallRule `json:"inbound"`
	Outbound   []WebArenaIndigoV1NwFirewallRule `json:"outbound"`
	Instances  []string                         `json:"instances"`
}

type UpdateWebArenaIndigoV1NwFirewallResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	SucessCode string `json:"sucessCode"`
	FirewallID int64  `json:"firewallId"`
}
