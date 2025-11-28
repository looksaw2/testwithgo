package adapters

import (
	"context"
	"log/slog"
	"strconv"
	"sync"
	"time"

	domain "github.com/looksaw2/gorder3/internal/order/domain/order"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

var fakeData = []*domain.Order{
	{
		ID:          "fake-id",
		CustomerID:  "fake-customer-id",
		Status:      "fake-status",
		PaymentLink: "fake-payment-link",
		Items:       nil,
	},
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: fakeData,
	}
}

func (m *MemoryOrderRepository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		PaymentLink: order.PaymentLink,
		Items:       order.Items,
	}
	m.store = append(m.store, newOrder)
	slog.Info("Order Service",
		slog.String("Handler", "Create Handler"),
		slog.Any("input_order", order),
		slog.Any("store_after_create", m.store),
	)
	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(ctx context.Context, id string, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, o := range m.store {
		if o.ID == id && o.CustomerID == customerID {
			slog.Info("Order Service",
				slog.String("Handler", "Get Handler"),
				slog.String("input_id", id),
				slog.String("input customer_id", customerID),
				slog.Any("output Order", &o),
			)
			return o, nil
		}
	}
	slog.Info("Order Service",
		slog.String("Handler" , "Get Handler"),
		slog.String("failed reason","can't found this Order"),
		slog.String("input_id", id),
		slog.String("input customer_id", customerID),
		//
		slog.String("Inmen ID [0] is",m.store[0].ID),
		slog.String("Inmem CustomerID [0] is ",m.store[0].CustomerID),
	)
	return nil, domain.NotFoundError{OrderID: id}
}

func (m *MemoryOrderRepository) Update(
	ctx context.Context,
	order *domain.Order,
	updateFn func(context.Context, *domain.Order) (*domain.Order, error),
) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	found := false
	for i, o := range m.store {
		if o.ID == order.ID && o.CustomerID == order.CustomerID {
			found = true
			updatedOrder, err := updateFn(ctx, o)
			if err != nil {
				return err
			}
			slog.Info("Order Service",
				slog.Any("input_order", order),
				slog.Any("update_order", updatedOrder),
			)
			m.store[i] = updatedOrder
		}
	}
	if !found {
		return domain.NotFoundError{OrderID: order.ID}
	}
	return nil
}
