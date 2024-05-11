package indigo

import (
	"context"
	"testing"

	require "github.com/kunitsucom/util.go/testing/require"
)

func TestClient_GetWebArenaIndigoV1VmOSList(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.GetWebArenaIndigoV1VmOSList(ctx, 1)
		require.NoError(t, err)
		require.NotNil(t, resp)
	})
}
