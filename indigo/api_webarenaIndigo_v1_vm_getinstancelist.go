package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Get Instance list
// https://indigo.arena.ne.jp/userapi/#instance_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/getinstancelist \
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
//		{
//			"id": 20,
//			"instance_name": "Centos Dev Env",
//			"set_no": 10,
//			"vps_kind": 10,
//			"sequence_id": 20,
//			"user_id": 18,
//			"service_id": "wsi-000200",
//			"status": "UNUSED",
//			"sshkey_id": 5,
//			"start_date": "2018-11-07 17:11:21",
//			"host_id": 2,
//			"plan": "2CR2GB",
//			"disk_point": 12,
//			"memsize": 2,
//			"cpus": 2,
//			"os_id": 1,
//			"otherstatus": 10,
//			"uuid": "92fa4815-1ec1-403b-93bc-514a11dfeb56",
//			"uidgid": 100001,
//			"vnc_port": 10001,
//			"vnc_passwd": "6z2d9ngprQuxofs6",
//			"arpaname": "192-168-0-20.pro.static.arena.ne.jp",
//			"arpadate": 0,
//			"status_change_date": "2018-11-07 17:11:21",
//			"updated_at": null,
//			"vm_revert": 0,
//			"VEID": "10100000000020",
//			"os": {
//				"id": 1,
//				"name": "CentOS6.6",
//				"viewname": "CentOS 6.6"
//			},
//			"ip": "192.168.0.20"
//		},
//		{
//			"id": 19,
//			"instance_name": "Centos Dev Env",
//			"set_no": 10,
//			"vps_kind": 10,
//			"sequence_id": 19,
//			"user_id": 18,
//			"service_id": "wsi-000200",
//			"status": "UNUSED",
//			"sshkey_id": 11,
//			"start_date": "2018-11-07 17:11:08",
//			"host_id": 2,
//			"plan": "2CR2GB",
//			"disk_point": 12,
//			"memsize": 2,
//			"cpus": 2,
//			"os_id": 1,
//			"otherstatus": 10,
//			"uuid": "372827ec-0946-494f-911c-fbb0bceba399",
//			"uidgid": 100001,
//			"vnc_port": 10001,
//			"vnc_passwd": "GXr3VkGhGofT3yf8",
//			"arpaname": "192-168-0-19.pro.static.arena.ne.jp",
//			"arpadate": 0,
//			"status_change_date": "2018-11-07 17:11:08",
//			"updated_at": null,
//			"vm_revert": 0,
//			"VEID": "10100000000019",
//			"os": {
//				"id": 1,
//				"name": "CentOS6.6",
//				"viewname": "CentOS 6.6"
//			},
//			"ip": "192.168.0.19"
//		}
//	]
func (c *Client) GetWebArenaIndigoV1VmGetInstanceList(ctx context.Context) (*GetWebArenaIndigoV1VmGetInstanceListResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1VmGetInstanceList, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &GetWebArenaIndigoV1VmGetInstanceListResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type GetWebArenaIndigoV1VmGetInstanceListResponse []WebArenaIndigoV1VmInstance
