package indigo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1NwGetTemplateFirewall struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
	Type      string `json:"type"`
	Protocol  string `json:"protocol"`
	Port      string `json:"port"`
	Source    string `json:"source"`
}

// Retrieve firewall
// https://indigo.arena.ne.jp/userapi/#retrive_firewall
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/nw/gettemplate/55 \
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
//	    "id": 55,
//	    "name": "Example",
//	    "direction": "in",
//	    "type": "HTTP",
//	    "protocol": "TCP",
//	    "port": "80",
//	    "source": "0.0.0.0"
//	},
//	{
//	    "id": 55,
//	    "name": "Example",
//	    "direction": "in",
//	    "type": "HTTPS",
//	    "protocol": "TCP",
//	    "port": "443",
//	    "source": "0.0.0.0"
//	},
//	{
//	    "id": 55,
//	    "name": "Example",
//	    "direction": "out",
//	    "type": "HTTP",
//	    "protocol": "TCP",
//	    "port": "80",
//	    "source": "0.0.0.0"
//	},
//	{
//	    "id": 55,
//	    "name": "Example",
//	    "direction": "out",
//	    "type": "HTTPS",
//	    "protocol": "TCP",
//	    "port": "443",
//	    "source": "0.0.0.0"
//	}
//
// IMPORTANT: The response body is not a valid JSON, so we need to read it as a byte slice and then append '[' and ']' to make it a valid JSON.
func (c *Client) GetWebArenaIndigoV1NwGetTemplate(ctx context.Context, firewallID int64) (*GetWebArenaIndigoV1NwGetTemplateResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, path.Join(PathWebArenaIndigoV1NwGetTemplate, strconv.FormatInt(firewallID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	// IMPORTANT: The response body is not a valid JSON, so we need to read it as a byte slice and then append '[' and ']' to make it a valid JSON.
	b, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errorz.Errorf("io.ReadAll: %w", err)
	}
	b = append([]byte{'['}, append(b, ']')...)
	resp := &GetWebArenaIndigoV1NwGetTemplateResponse{}
	if err := json.Unmarshal(b, resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type GetWebArenaIndigoV1NwGetTemplateResponse []WebArenaIndigoV1NwGetTemplateFirewall
