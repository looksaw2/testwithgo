package decorator

import (
	"context"
	"log/slog"
)

//使用Command模式
type CommandHandler[C any , R any] interface {
	Handle(ctx context.Context , cmd C)(R,error)
}

func ApplyCommandDecorators[C any , R any](handler CommandHandler[C,R],logger *slog.Logger,metricsClient MetricsClient)CommandHandler[C,R]{
	return &queryLoggingDecorator[C,R]{
		logger: logger,
		base: &queryMetricsDecorator[C,R]{
			client: metricsClient,
			base: handler,
		},
	}
}