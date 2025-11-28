package command

import (
	"context"
	"log/slog"

	"github.com/looksaw2/gorder3/internal/common/decorator"
	"github.com/looksaw2/gorder3/internal/common/genproto/orderpb"
	domain "github.com/looksaw2/gorder3/internal/order/domain/order"
)


type CreateOrder struct {
	CustomerID string
	Items []*orderpb.ItemWIthQuantity
}

type CreateOrderResult struct {
	OrderID string
}


type CreateOrderHandler decorator.CommandHandler[CreateOrder,*CreateOrderResult]


type createOrderHandler struct {
	orderRepo domain.Repository
}

func (c *createOrderHandler)Handle(ctx context.Context , cmd CreateOrder)(*CreateOrderResult,error){
	//TODO GRPC
	var stockResponse  []*orderpb.Item
	for _ ,item := range cmd.Items {
		stockResponse = append(stockResponse , &orderpb.Item{
			ID: item.ID,
			Quantity: item.Quantity,
		})
	}
	o ,err := c.orderRepo.Create(ctx,&domain.Order{
		CustomerID: cmd.CustomerID,
		Items: stockResponse,
	})
	if err != nil {
		return nil ,err
	}
	return &CreateOrderResult{OrderID: o.ID} , nil
}

func newcreateOrderHandler(oderRepo domain.Repository) *createOrderHandler {
	return &createOrderHandler{
		orderRepo: oderRepo,
	}
}


func NewCreateOrderHandler(
	orderRepo domain.Repository,
	logger *slog.Logger,
	metricsClient decorator.MetricsClient,
)CreateOrderHandler {
	return decorator.ApplyCommandDecorators[CreateOrder,*CreateOrderResult](
		newcreateOrderHandler(orderRepo),
		logger,
		metricsClient,
	)
}
