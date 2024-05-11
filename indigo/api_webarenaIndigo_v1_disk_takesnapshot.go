package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	errorz "github.com/kunitsucom/util.go/errors"
)

// Snapshot Creation
// https://indigo.arena.ne.jp/userapi/#snapshot_creation
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/disk/takesnapshot \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{"name":"Example", "instanceid":12, "slotnum":"0"}'
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 Created
// RESPONSE BODY
//
//	{"STATUS":0}.
func (c *Client) PostWebArenaIndigoV1DiskTakeSnapshot(ctx context.Context, req *PostWebArenaIndigoV1DiskTakeSnapshotRequest) (*PostWebArenaIndigoV1DiskTakeSnapshotResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1DiskTakeSnapshot, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &PostWebArenaIndigoV1DiskTakeSnapshotResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type PostWebArenaIndigoV1DiskTakeSnapshotRequest struct {
	Name       string `json:"name"`
	InstanceID int64  `json:"instanceid"`
	SlotNum    string `json:"slotnum"`
}

type PostWebArenaIndigoV1DiskTakeSnapshotResponse struct {
	Status int `json:"STATUS"` //nolint:tagliatelle // JSON field name is defined by the API
}
