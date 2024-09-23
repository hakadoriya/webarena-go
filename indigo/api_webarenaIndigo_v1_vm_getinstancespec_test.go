package indigo

import (
	"context"
	"testing"

	"github.com/hakadoriya/z.go/testingz/requirez"
)

func TestClient_GetWebArenaIndigoV1VmInstanceSpec(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.GetWebArenaIndigoV1VmInstanceSpec(ctx, 1, 1)
		requirez.NoError(t, err)
		requirez.NotNil(t, resp)
	})
}
