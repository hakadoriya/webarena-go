package indigo

import (
	"context"
	"math"
	"testing"

	"github.com/hakadoriya/z.go/testingz/requirez"
)

func TestClient_GetWebArenaIndigoV1DiskSnapshotList(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.GetWebArenaIndigoV1DiskSnapshotList(ctx, math.MaxInt)
		requirez.ErrorIs(t, err, ErrUnexpectedStatusCode)
		requirez.Nil(t, resp)
	})
}
