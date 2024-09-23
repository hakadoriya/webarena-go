package indigo

import (
	"context"
	"testing"

	"github.com/hakadoriya/z.go/testingz/requirez"
)

func TestClient_GetWebArenaIndigoV1NwGetFirewallList(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.GetWebArenaIndigoV1NwGetFirewallList(ctx)
		requirez.NoError(t, err)
		requirez.NotNil(t, resp)
	})
}
