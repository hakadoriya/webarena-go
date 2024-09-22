package indigo

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

// SSH Key List
// https://indigo.arena.ne.jp/userapi/#list_ssh_keys
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR'
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR.
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
//	            "sshkey": "example1",
//	            "status": "ACTIVE",
//	            "created_at": "2018-11-01 17:18:00",
//	            "updated_at": "2018-11-01 17:18:12"
//	        },
//	        {
//	            "id": 5,
//	            "service_id": "wsi-000001",
//	            "user_id": 431,
//	            "name": "Example",
//	            "sshkey": "example2",
//	            "status": "ACTIVE",
//	            "created_at": "2018-11-01 05:37:03",
//	            "updated_at": "2018-11-01 05:37:03"
//	        }
//
//	    ]
//	}
func (c *Client) GetWebArenaIndigoV1VmSSHKey(ctx context.Context) (*WebArenaIndigoV1VmSSHKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, PathWebArenaIndigoV1VmSSHKey, nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &WebArenaIndigoV1VmSSHKeyResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type WebArenaIndigoV1VmSSHKeyResponse struct {
	Success bool                       `json:"success"`
	Total   int64                      `json:"total"`
	Sshkeys []WebArenaIndigoV1VmSSHKey `json:"sshkeys"`
}

type WebArenaIndigoV1VmSSHKey struct {
	Id        int64  `json:"id"`         //nolint:revive,stylecheck
	ServiceId string `json:"service_id"` //nolint:revive,stylecheck,tagliatelle
	UserId    int64  `json:"user_id"`    //nolint:revive,stylecheck
	Name      string `json:"name"`
	Sshkey    string `json:"sshkey"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"` //nolint:revive,stylecheck,tagliatelle
	UpdatedAt string `json:"updated_at"` //nolint:revive,stylecheck,tagliatelle
}

// Create SSH Key
// https://indigo.arena.ne.jp/userapi/#create_ssh_keys
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	   "sshName":"Example",
//	   "sshKey":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDRGTcjdlRYZ9"
//	  }'
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 201 Created
// RESPONSE BODY
//
//	{
//	    "success": true,
//	    "message": "SSH key has been added successfully",
//	    "sshKey": {
//	        "name": "Example",
//	        "sshkey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDRGTcjdlRYZ9",
//	        "status": "ACTIVE",
//	        "user_id": 431,
//	        "service_id": "wsi-000001",
//	        "updated_at": "2019-10-23 11:24:31",
//	        "created_at": "2019-10-23 11:24:31",
//	        "id": 892
//	    }
//	}
func (c *Client) CreateWebArenaIndigoV1VmSSHKey(ctx context.Context, req *CreateWebArenaIndigoV1VmSSHKeyRequest) (*CreateWebArenaIndigoV1VmSSHKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmSSHKey, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &CreateWebArenaIndigoV1VmSSHKeyResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return resp, nil
}

type CreateWebArenaIndigoV1VmSSHKeyRequest struct {
	SshName string `json:"sshName"` //nolint:revive,stylecheck
	SshKey  string `json:"sshKey"`  //nolint:revive,stylecheck
}

type CreateWebArenaIndigoV1VmSSHKeyResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	SshKey  WebArenaIndigoV1VmSSHKey `json:"sshKey"` //nolint:revive,stylecheck
}

// Retrieve SSH Key
// https://indigo.arena.ne.jp/userapi/#retrive_ssh_keys
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey/34 \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
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
//	    "sshKey": [
//	        {
//	            "id": 5,
//	            "service_id": "wsi-000001",
//	            "user_id": 431,
//	            "name": "Example",
//	            "sshkey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDRGTcjdlRYZ9J4KEaZ3A8FwPSWKHak1UKUusSX",
//	            "status": "ACTIVE",
//	            "created_at": "2018-11-01 17:35:32",
//	            "updated_at": "2018-11-01 17:35:32"
//	        }
//	    ]
//	}
func (c *Client) RetrieveWebArenaIndigoV1VmSSHKey(ctx context.Context, sshKeyID int64) (*RetrieveWebArenaIndigoV1VmSSHKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, path.Join(PathWebArenaIndigoV1VmSSHKey, strconv.FormatInt(sshKeyID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &RetrieveWebArenaIndigoV1VmSSHKeyResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type RetrieveWebArenaIndigoV1VmSSHKeyResponse struct {
	Success bool                       `json:"success"`
	SshKey  []WebArenaIndigoV1VmSSHKey `json:"sshKey"` //nolint:revive,stylecheck
}

// Update SSH Key
// https://indigo.arena.ne.jp/userapi/#update_ssh_keys
//
// CURL EXAMPLE
//
//	curl -X PUT \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey/34 \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	   "sshName": "Example",
//	   "sshKey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDRGTcjdlRYZ9J4KEaZ3A8FwPSWKHak1UKUusSX",
//	   "sshKeyStatus": "ACTIVE"
//	  }'
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
//	    "message": "SSH key has been updated successfully"
//	}
func (c *Client) UpdateWebArenaIndigoV1VmSSHKey(ctx context.Context, id int64, req *UpdateWebArenaIndigoV1VmSSHKeyRequest) (*UpdateWebArenaIndigoV1VmSSHKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPut, path.Join(PathWebArenaIndigoV1VmSSHKey, strconv.FormatInt(id, 10)), body) //nolint:staticcheck // NOTE: PUT
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &UpdateWebArenaIndigoV1VmSSHKeyResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return resp, nil
}

type UpdateWebArenaIndigoV1VmSSHKeyRequest struct {
	SshName     string `json:"sshName"`     //nolint:revive,stylecheck
	SshKey      string `json:"sshKey"`      //nolint:revive,stylecheck
	SshKeyState string `json:"sshKeyState"` //nolint:revive,stylecheck
}

type UpdateWebArenaIndigoV1VmSSHKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Destroy SSH Key
// https://indigo.arena.ne.jp/userapi/#destroy_ssh_keys
//
// CURL EXAMPLE
//
//	curl -X DELETE \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/sshkey/34 \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
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
//	    "message": "SSH key has been removed successfully"
//	}
func (c *Client) DestroyWebArenaIndigoV1VmSSHKey(ctx context.Context, sshKeyID int64) (*DestroyWebArenaIndigoV1VmSSHKeyResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodDelete, path.Join(PathWebArenaIndigoV1VmSSHKey, strconv.FormatInt(sshKeyID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &DestroyWebArenaIndigoV1VmSSHKeyResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type DestroyWebArenaIndigoV1VmSSHKeyResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
