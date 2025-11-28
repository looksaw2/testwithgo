package service

import (
	"log/slog"

	"github.com/looksaw2/gorder3/internal/common/metrics"
	"github.com/looksaw2/gorder3/internal/order/adapters"
	"github.com/looksaw2/gorder3/internal/order/app"
	"github.com/looksaw2/gorder3/internal/order/app/query"
)

func NewApplication() *app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := slog.Default()
	metricsClient := metrics.NewTodoMetrics()
	return &app.Application{
		Command: &app.Command{},
		Queries: &app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(
				orderRepo,
				logger,
				metricsClient,
			),
		},
	}
}
