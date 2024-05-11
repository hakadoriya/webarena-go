package indigo

import (
	"context"

	"github.com/kunitsucom/webarena-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func start(ctx context.Context, opts ...trace.SpanStartOption) (context.Context, trace.Span) { //nolint:ireturn,unparam
	return otel.GetTracerProvider().Tracer(pkgPath).Start(ctx, util.FuncName(1), opts...) //nolint:spancheck
}
