package indigo

import (
	"context"
	"testing"

	"github.com/hakadoriya/z.go/testingz/requirez"
)

//nolint:tparallel,paralleltest
func TestClient_PostOAuthV1AccessTokens(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.PostOAuthV1AccessTokens(ctx, &PostOAuthV1AccessTokensRequest{
			GrantType:    "client_credentials",
			ClientId:     client.clientID,
			ClientSecret: client.clientSecret,
			Code:         "",
		})
		requirez.NoError(t, err)
		requirez.NotNil(t, resp)
	})
}
