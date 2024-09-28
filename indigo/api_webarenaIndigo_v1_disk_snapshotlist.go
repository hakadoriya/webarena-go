package indigo

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

type WebArenaIndigoV1DiskSnapshot struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	ServiceID          string `json:"service_id"` //nolint:tagliatelle // JSON field name is defined by the API
	UserID             string `json:"user_id"`    //nolint:tagliatelle // JSON field name is defined by the API
	DiskID             int64  `json:"disk_id"`    //nolint:tagliatelle // JSON field name is defined by the API
	Volume             int64  `json:"volume"`
	SlotNumber         int64  `json:"slot_number"` //nolint:tagliatelle // JSON field name is defined by the API
	Status             string `json:"status"`
	Size               string `json:"size"`
	Deleted            int64  `json:"deleted"`
	CompletedTimestamp string `json:"completed_timestamp"` //nolint:tagliatelle // JSON field name is defined by the API
	DeletedTimestamp   string `json:"deleted_timestamp"`   //nolint:tagliatelle // JSON field name is defined by the API
}

// Snapshot list
// https://indigo.arena.ne.jp/userapi/#snapshot_list
//
// CURL EXAMPLE
//
//	curl -X GET \
//	  https://api.customer.jp/webarenaIndigo/v1/disk/snapshotlist/{instanceid} \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR'
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 OK
// RESPONSE BODY
// [
//
//	{
//	    "id": 3,
//	    "name": "Example",
//	    "service_id": "wsi-000401",
//	    "user_id": "134",
//	    "disk_id": 29,
//	    "volume": 1,
//	    "slot_number": 0,
//	    "status": "created",
//	    "size": "2000",
//	    "deleted": 0,
//	    "completed_timestamp": "2018-11-27 07:24:05",
//	    "deleted_timestamp": "0000-00-00 00:00:00"
//	},
//	{
//	    "id": 8,
//	    "name": "Example2",
//	    "service_id": "wsi-000401",
//	    "user_id": "134",
//	    "disk_id": 29,
//	    "volume": 2,
//	    "slot_number": 0,
//	    "status": "failed",
//	    "size": "2000",
//	    "deleted": 0,
//	    "completed_timestamp": "2018-11-27 10:43:22",
//	    "deleted_timestamp": "0000-00-00 00:00:00"
//	}
//
// ].
func (c *Client) GetWebArenaIndigoV1DiskSnapshotList(ctx context.Context, instanceID int64) (*GetWebArenaIndigoV1DiskSnapshotListResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodGet, path.Join(PathWebArenaIndigoV1DiskSnapshotList, strconv.FormatInt(instanceID, 10)), nil)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp GetWebArenaIndigoV1DiskSnapshotListResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type GetWebArenaIndigoV1DiskSnapshotListResponse []WebArenaIndigoV1DiskSnapshot
