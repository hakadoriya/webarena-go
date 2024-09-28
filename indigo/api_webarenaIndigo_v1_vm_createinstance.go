package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1VmInstanceDate struct {
	Date         string `json:"date"`
	TimezoneType int64  `json:"timezone_type"` //nolint:tagliatelle // JSON field name is defined by the API
	Timezone     string `json:"timezone"`
}

func (d WebArenaIndigoV1VmInstanceDate) UnmarshalJSON(b []byte) error {
	type Alias WebArenaIndigoV1VmInstanceDate
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(&d),
	}
	if err := json.Unmarshal(b, &aux); err == nil {
		d = WebArenaIndigoV1VmInstanceDate(*aux.Alias)
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return errorz.Errorf("json.Unmarshal: %w", err)
	}

	d.Date = s
	d.TimezoneType = 3
	d.Timezone = "UTC"

	return nil
}

type WebArenaIndigoV1VmInstanceOS struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"` //nolint:tagliatelle // JSON field name is defined by the API
	ViewName string `json:"viewname"`
}

type WebArenaIndigoV1VmInstance struct {
	ID               int64                          `json:"id"`
	InstanceName     string                         `json:"instance_name"` //nolint:tagliatelle // JSON field name is defined by the API
	SetNo            int64                          `json:"set_no"`        //nolint:tagliatelle // JSON field name is defined by the API
	VpsKind          string                         `json:"vps_kind"`      //nolint:tagliatelle // JSON field name is defined by the API
	SequenceID       int64                          `json:"sequence_id"`   //nolint:tagliatelle // JSON field name is defined by the API
	UserID           int64                          `json:"user_id"`       //nolint:tagliatelle // JSON field name is defined by the API
	ServiceID        string                         `json:"service_id"`    //nolint:tagliatelle // JSON field name is defined by the API
	Status           string                         `json:"status"`
	SshKeyID         int64                          `json:"sshkey_id"`  //nolint:revive,stylecheck,tagliatelle // JSON field name is defined by the API
	StartDate        WebArenaIndigoV1VmInstanceDate `json:"start_date"` //nolint:tagliatelle // JSON field name is defined by the API
	HostID           int64                          `json:"host_id"`    //nolint:tagliatelle // JSON field name is defined by the API
	Plan             string                         `json:"plan"`
	DiskPoint        int64                          `json:"disk_point"` //nolint:tagliatelle // JSON field name is defined by the API
	MemSize          int64                          `json:"memsize"`
	CPUs             int64                          `json:"cpus"`
	OsID             int64                          `json:"os_id"` //nolint:tagliatelle // JSON field name is defined by the API
	OtherStatus      int64                          `json:"otherstatus"`
	UUID             string                         `json:"uuid"`
	UIDGID           int64                          `json:"uidgid"`
	VncPort          int64                          `json:"vnc_port"`   //nolint:tagliatelle // JSON field name is defined by the API
	VncPasswd        string                         `json:"vnc_passwd"` //nolint:tagliatelle // JSON field name is defined by the API
	ArpaName         string                         `json:"arpaname"`
	ArpaDate         string                         `json:"arpadate"`
	StatusChangeDate WebArenaIndigoV1VmInstanceDate `json:"status_change_date"` //nolint:tagliatelle // JSON field name is defined by the API
	UpdatedAt        string                         `json:"updated_at"`         //nolint:tagliatelle // JSON field name is defined by the API
	VMRevert         int64                          `json:"vm_revert"`          //nolint:tagliatelle // JSON field name is defined by the API
	VEID             string                         `json:"VEID"`               //nolint:tagliatelle // JSON field name is defined by the API
	OS               WebArenaIndigoV1VmInstanceOS   `json:"os"`
	IP               string                         `json:"ip"`
}

// Instance Creation
// https://indigo.arena.ne.jp/userapi/#instance_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/createinstance \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	  "sshKeyId": 11,
//	  "regionId": 1,
//	  "osId": 1,
//	  "instancePlan": 1,
//	  "instanceName": "Centos Dev Env"
//	}'
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
//	    "message": "Instance created  successfully",
//	    "vms": {
//	        "id": 3,
//	        "instance_name": "Centos Dev Env",
//	        "set_no": 10,
//	        "vps_kind": "10",
//	        "sequence_id": 3,
//	        "user_id": 1,
//	        "service_id": "wsi-000001",
//	        "status": "UNUSED",
//	        "sshkey_id": 1,
//	        "start_date": {
//	            "date": "2018-11-10 10:03:17.744562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "host_id": 1,
//	        "plan": "2CR2GB",
//	        "disk_point": 12,
//	        "memsize": 2,
//	        "cpus": 2,
//	        "os_id": 1,
//	        "otherstatus": 10,
//	        "uuid": "9868faef-e658-4880-b4ad-fd3078e51b6a",
//	        "uidgid": 100003,
//	        "vnc_port": 10003,
//	        "vnc_passwd": "fHTsl4EoLfMksYKW",
//	        "arpaname": "192-168-0-3.pro.static.arena.ne.jp",
//	        "arpadate": "",
//	        "status_change_date": {
//	            "date": "2018-11-10 10:03:17.746562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "updated_at": null,
//	        "vm_revert": 0
//	    }
//	}
func (c *Client) PostWebArenaIndigoV1VmCreateInstance(ctx context.Context, req *PostWebArenaIndigoV1VmCreateInstanceRequest) (*PostWebArenaIndigoV1VmCreateInstanceResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmCreateInstance, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1VmCreateInstanceResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1VmCreateInstanceRequest struct {
	SshKeyID     int64  `json:"sshKeyId"` //nolint:revive,stylecheck
	RegionID     int64  `json:"regionId"`
	OsID         int64  `json:"osId"`
	InstancePlan int64  `json:"instancePlan"`
	InstanceName string `json:"instanceName"`
}

type PostWebArenaIndigoV1VmCreateInstanceResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Vms     WebArenaIndigoV1VmInstance `json:"vms"`
}

// Windows Instance Creation
// https://indigo.arena.ne.jp/userapi/#windows_instance_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/createinstance \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	  "winPassword": "Test#2345jh",
//	  "regionId": 1,
//	  "osId": 1,
//	  "instancePlan": 1,
//	  "instanceName": "Centos Dev Env"
//	}'
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
//	    "message": "Instance created  successfully",
//	    "vms": {
//	        "id": 3,
//	        "instance_name": "Centos Dev Env",
//	        "set_no": 10,
//	        "vps_kind": "10",
//	        "sequence_id": 3,
//	        "user_id": 1,
//	        "service_id": "wsi-000001",
//	        "status": "UNUSED",
//	        "sshkey_id": 1,
//	        "start_date": {
//	            "date": "2018-11-10 10:03:17.744562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "host_id": 1,
//	        "plan": "2CR2GB",
//	        "disk_point": 12,
//	        "memsize": 2,
//	        "cpus": 2,
//	        "os_id": 1,
//	        "otherstatus": 10,
//	        "uuid": "9868faef-e658-4880-b4ad-fd3078e51b6a",
//	        "uidgid": 100003,
//	        "vnc_port": 10003,
//	        "vnc_passwd": "fHTsl4EoLfMksYKW",
//	        "arpaname": "192-168-0-3.pro.static.arena.ne.jp",
//	        "arpadate": "",
//	        "status_change_date": {
//	            "date": "2018-11-10 10:03:17.746562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "updated_at": null,
//	        "vm_revert": 0
//	    }
//	}
func (c *Client) PostWebArenaIndigoV1VmCreateWindowsInstance(ctx context.Context, req *PostWebArenaIndigoV1VmCreateWindowsInstanceRequest) (*PostWebArenaIndigoV1VmCreateWindowsInstanceResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmCreateInstance, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1VmCreateWindowsInstanceResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1VmCreateWindowsInstanceRequest struct {
	WinPassword  string `json:"winPassword"`
	RegionID     int64  `json:"regionId"`
	OsID         int64  `json:"osId"`
	InstancePlan int64  `json:"instancePlan"`
	InstanceName string `json:"instanceName"`
}

type PostWebArenaIndigoV1VmCreateWindowsInstanceResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Vms     WebArenaIndigoV1VmInstance `json:"vms"`
}

// Import URL Instance Creation
// https://indigo.arena.ne.jp/userapi/#import_instance_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/createinstance \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	  "importUrl": "https://203.138.52.188/cloudn_os/CentOS7.qcow2",
//	  "regionId": 1,
//	  "osId": 1,
//	  "instancePlan": 2,
//	  "instanceName": "ImportURLinstance"
//	}'
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
//	    "message": "Instance created  successfully",
//	    "vms": {
//	        "id": 3,
//	        "instance_name": "Centos Dev Env",
//	        "set_no": 10,
//	        "vps_kind": "10",
//	        "sequence_id": 3,
//	        "user_id": 1,
//	        "service_id": "wsi-000001",
//	        "status": "UNUSED",
//	        "sshkey_id": 1,
//	        "start_date": {
//	            "date": "2018-11-10 10:03:17.744562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "host_id": 1,
//	        "plan": "2CR2GB",
//	        "disk_point": 12,
//	        "memsize": 2,
//	        "cpus": 2,
//	        "os_id": 1,
//	        "otherstatus": 10,
//	        "uuid": "9868faef-e658-4880-b4ad-fd3078e51b6a",
//	        "uidgid": 100003,
//	        "vnc_port": 10003,
//	        "vnc_passwd": "fHTsl4EoLfMksYKW",
//	        "arpaname": "192-168-0-3.pro.static.arena.ne.jp",
//	        "arpadate": "",
//	        "status_change_date": {
//	            "date": "2018-11-10 10:03:17.746562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "updated_at": null,
//	        "vm_revert": 0
//	    }
//	}
func (c *Client) PostWebArenaIndigoV1VmCreateImportURLInstance(ctx context.Context, req *PostWebArenaIndigoV1VmCreateImportURLInstanceRequest) (*PostWebArenaIndigoV1VmCreateImportURLInstanceResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmCreateInstance, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1VmCreateImportURLInstanceResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1VmCreateImportURLInstanceRequest struct {
	ImportURL    string `json:"importUrl"`
	RegionID     int64  `json:"regionId"`
	OsID         int64  `json:"osId"`
	InstancePlan int64  `json:"instancePlan"`
	InstanceName string `json:"instanceName"`
}

type PostWebArenaIndigoV1VmCreateImportURLInstanceResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Vms     WebArenaIndigoV1VmInstance `json:"vms"`
}

// Snapshot Instance Creation
// https://indigo.arena.ne.jp/userapi/#snapshot_instance_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/vm/createinstance \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{
//	  "sshKeyId": 1480,
//	  "snapshotId": "457",
//	  "instancePlan": 5,
//	  "instanceName": "from-api-2"
//	}'
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
//	    "message": "Instance created  successfully",
//	    "vms": {
//	        "id": 3,
//	        "instance_name": "Centos Dev Env",
//	        "set_no": 10,
//	        "vps_kind": "10",
//	        "sequence_id": 3,
//	        "user_id": 1,
//	        "service_id": "wsi-000001",
//	        "status": "UNUSED",
//	        "sshkey_id": 1,
//	        "start_date": {
//	            "date": "2018-11-10 10:03:17.744562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "host_id": 1,
//	        "plan": "2CR2GB",
//	        "disk_point": 12,
//	        "memsize": 2,
//	        "cpus": 2,
//	        "os_id": 1,
//	        "otherstatus": 10,
//	        "uuid": "9868faef-e658-4880-b4ad-fd3078e51b6a",
//	        "uidgid": 100003,
//	        "vnc_port": 10003,
//	        "vnc_passwd": "fHTsl4EoLfMksYKW",
//	        "arpaname": "192-168-0-3.pro.static.arena.ne.jp",
//	        "arpadate": "",
//	        "status_change_date": {
//	            "date": "2018-11-10 10:03:17.746562",
//	            "timezone_type": 3,
//	            "timezone": "UTC"
//	        },
//	        "updated_at": null,
//	        "vm_revert": 0
//	    }
//	}
func (c *Client) PostWebArenaIndigoV1VmCreateSnapshotInstance(ctx context.Context, req *PostWebArenaIndigoV1VmCreateSnapshotInstanceRequest) (*PostWebArenaIndigoV1VmCreateSnapshotInstanceResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1VmCreateInstance, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1VmCreateSnapshotInstanceResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1VmCreateSnapshotInstanceRequest struct {
	SshKeyID     int64  `json:"sshKeyId"` //nolint:revive,stylecheck
	SnapshotID   string `json:"snapshotId"`
	InstancePlan int64  `json:"instancePlan"`
	InstanceName string `json:"instanceName"`
}

type PostWebArenaIndigoV1VmCreateSnapshotInstanceResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Vms     WebArenaIndigoV1VmInstance `json:"vms"`
}
