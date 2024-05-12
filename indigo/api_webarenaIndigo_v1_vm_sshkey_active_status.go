package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	errorz "github.com/kunitsucom/util.go/errors"
)

type WebArenaIndigoV1VmSSHKeyActiveStatusResponse struct {
	Success bool                       `json:"success"`
	Total   int                        `json:"total"`
	Sshkeys []WebArenaIndigoV1VmSSHKey `json:"sshkeys"`
}

// Active SSH Key List
// https://indigo.arena.ne.jp/userapi/#get_active_sshkey_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey/active/status \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR'
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
//	    "total": 2,
//	    "sshkeys": [
//	        {
//	            "id": 5,
//	            "service_id": "wsi-000001",
//	            "user_id": 431,
//	            "name": "Example",
//	            "sshkey": "examplekey1",
//	            "status": "ACTIVE",
//	            "created_at": "2018-11-01 17:18:00",
//	            "updated_at": "2018-11-01 17:18:12"
//	        },
//	        {
//	            "id": 5,
//	            "service_id": "wsi-000001",
//	            "user_id": 431,
//	            "name": "Example",
//	            "sshkey": "examplekey2",
//	            "status": "ACTIVE",
//	            "created_at": "2018-11-01 05:37:03",
//	            "updated_at": "2018-11-01 05:37:03"
//	        }
//
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmSSHKeyActiveStatus(ctx context.Context) (*WebArenaIndigoV1VmSSHKeyActiveStatusResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1VmSSHKeyActiveStatus, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &WebArenaIndigoV1VmSSHKeyActiveStatusResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}
