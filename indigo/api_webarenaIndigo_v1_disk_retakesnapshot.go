package indigo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hakadoriya/z.go/errorz"
)

// Retry to create snapshot
// https://indigo.arena.ne.jp/userapi/#retry_snapshot
//
// CURL EXAMPLE
//
//	curl -X POST \
//	  https://api.customer.jp/webarenaIndigo/v1/disk/retakesnapshot \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//	  -d '{"instanceid":12, "snapshotid":"11"}'
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
func (c *Client) PostWebArenaIndigoV1DiskRetakeSnapshot(ctx context.Context, req *PostWebArenaIndigoV1DiskRetakeSnapshotRequest) (*PostWebArenaIndigoV1DiskRetakeSnapshotResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	body, err := json.Marshal(req)
	if err != nil {
		return nil, errorz.Errorf("json.Marshal: %w", err)
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, PathWebArenaIndigoV1DiskRetakeSnapshot, body)
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	resp := &PostWebArenaIndigoV1DiskRetakeSnapshotResponse{}
	if err := json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return resp, nil
}

type PostWebArenaIndigoV1DiskRetakeSnapshotRequest struct {
	InstanceID int64  `json:"instanceid"`
	SnapshotID string `json:"snapshotid"`
}

type PostWebArenaIndigoV1DiskRetakeSnapshotResponse struct {
	Status int64 `json:"STATUS"` //nolint:tagliatelle // JSON field name is defined by the API
}
