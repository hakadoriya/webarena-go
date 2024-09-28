package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Restore snapshot
// https://indigo.arena.ne.jp/userapi/#restore_snapshot
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/disk/restoresnapshot \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{"instanceid":"12", "snapshotid":"12"}'
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 Created
// RESPONSE BODY
// {"STATUS":0}.
func (c *Client) PostWebArenaIndigoV1DiskRestoreSnapshot(ctx context.Context, req *PostWebArenaIndigoV1DiskRestoreSnapshotRequest) (*PostWebArenaIndigoV1DiskRestoreSnapshotResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1DiskRestoreSnapshot, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp PostWebArenaIndigoV1DiskRestoreSnapshotResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type PostWebArenaIndigoV1DiskRestoreSnapshotRequest struct {
	InstanceID int64  `json:"instanceid"`
	SnapshotID string `json:"snapshotid"`
}

type PostWebArenaIndigoV1DiskRestoreSnapshotResponse struct {
	STATUS int64 `json:"STATUS"` //nolint:tagliatelle // JSON field name is defined by the API
}
