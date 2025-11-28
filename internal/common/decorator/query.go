package decorator

import (
	"context"
	"log/slog"
)


type QueryHandler[Q any , R any] interface {
	Handle(ctx context.Context , query Q)(R , error)
}

//实现logger的Decorator并且
func ApplyQueryHandler[H any ,R any](handler QueryHandler[H,R] , logger *slog.Logger ,metricsClient MetricsClient)QueryHandler[H,R]{
	return &queryLoggingDecorator[H,R]{
		logger: logger,
		base: &queryMetricsDecorator[H,R]{
			client: metricsClient,
			base: handler,
		},
	}
} 