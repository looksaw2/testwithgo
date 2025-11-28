package order

import (
	"context"
	"fmt"
)

//DDD的存储模式

type Repository interface {
	Create(context.Context, *Order) (*Order, error)
	Get(context.Context, string, string) (*Order, error)
	Update(
		context.Context,
		*Order,
		func(context.Context, *Order) (*Order, error),
	) error
}

type NotFoundError struct {
	OrderID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("order ID %s not found", e.OrderID)
}
