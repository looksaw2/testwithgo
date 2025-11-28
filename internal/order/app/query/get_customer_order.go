package query

import (
	"context"
	"log/slog"

	"github.com/looksaw2/gorder3/internal/common/decorator"
	domain "github.com/looksaw2/gorder3/internal/order/domain/order"
)

type GetCustomerOrder struct {
	CustomerID string
	OrderID    string
}

type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

type getCustomerOrderHandler struct {
	orderRepo domain.Repository
}

func newgetCustomerOrderHandler(orderRepo domain.Repository) *getCustomerOrderHandler {
	return &getCustomerOrderHandler{
		orderRepo: orderRepo,
	}
}

func (g *getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	o, err := g.orderRepo.Get(ctx, query.OrderID, query.CustomerID)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func NewGetCustomerOrderHandler(
	orderRepo domain.Repository,
	logger *slog.Logger,
	metricClient decorator.MetricsClient,
) GetCustomerOrderHandler {
	if orderRepo == nil {
		panic("error order repository counldn't be nil")
	}
	return decorator.ApplyQueryHandler[GetCustomerOrder, *domain.Order](
		newgetCustomerOrderHandler(orderRepo),
		logger,
		metricClient,
	)
}
