package indigo

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/hakadoriya/z.go/errorz"
	"github.com/hakadoriya/z.go/syncz"
	"github.com/hakadoriya/z.go/testingz/requirez"
)

var (
	testClient     *Client
	testClientOnce syncz.Once
)

func NewTestClient(ctx context.Context, tb testing.TB) *Client {
	tb.Helper()

	if err := testClientOnce.Do(func() error {
		client, err := NewClient(ctx, ClientOptionWithDebugLog(log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)))
		if err != nil {
			return errorz.Errorf("NewClient: %w", err)
		}
		testClient = client
		return nil
	}); err != nil {
		if errors.Is(err, ErrInvalidClientCredentials) {
			tb.Skipf("⏸️: testClientOnce.Do: %v", err)
		} else {
			tb.Fatalf("❌: testClientOnce.Do: %v", err)
		}
	}

	return testClient
}

//nolint:tparallel,paralleltest
func TestClient_refreshAccessToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()

		client := NewTestClient(ctx, t)
		accessToken, err := client.IssueAccessToken(ctx)
		requirez.NoError(t, err)
		requirez.NotNil(t, accessToken)
	})
}
