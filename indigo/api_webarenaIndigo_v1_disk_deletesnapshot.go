package indigo

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/hakadoriya/z.go/errorz"
)

// Delete snapshot
// https://indigo.arena.ne.jp/userapi/#delete_snapshot
//
// CURL EXAMPLE
//
//	curl -X DELETE \
//	  https://api.customer.jp/webarenaIndigo/v1/disk/deletesnapshot/10 \
//	  -H 'Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR' \
//
// REQUEST HEADERS
// Authorization: Bearer HnGma03YnFPIF4DMttywiSOCGUHR
//
// RESPONSE HEADERS
// content-type: application/json;
// status: 200 Created
// RESPONSE BODY
// {"STATUS":0}.
func (c *Client) DeleteWebArenaIndigoV1DiskDeleteSnapshot(ctx context.Context, snapshotID int64) (*DeleteWebArenaIndigoV1DiskDeleteSnapshotResponse, error) {
	ctx, span := start(ctx)
	defer span.End()

	httpReq, err := c.newRequest(ctx, http.MethodDelete, path.Join(PathWebArenaIndigoV1DiskDeleteSnapshot, strconv.FormatInt(snapshotID, 10)), nil) //nolint:staticcheck
	if err != nil {
		return nil, errorz.Errorf("c.newRequest: %w", err)
	}

	httpResp, err := c.doRequest(httpReq)
	if err != nil {
		return nil, errorz.Errorf("c.doRequest: %w", err)
	}
	defer httpResp.Body.Close()

	var resp DeleteWebArenaIndigoV1DiskDeleteSnapshotResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, errorz.Errorf("json.Decode: %w", err)
	}

	return &resp, nil
}

type DeleteWebArenaIndigoV1DiskDeleteSnapshotResponse struct {
	STATUS int64 `json:"STATUS"` //nolint:tagliatelle // JSON field name is defined by the API
}
