package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/looksaw2/gorder3/internal/order/app"
	"github.com/looksaw2/gorder3/internal/order/app/query"
)

type HTTPHandler struct {
	app *app.Application
}


func NewHTTPHandler(app *app.Application) *HTTPHandler{
	return  &HTTPHandler{
		app: app,
	}
}

func (h *HTTPHandler) PostCustomerCustomerIdOrders(
	w http.ResponseWriter,
	r *http.Request,
	customerId string) {
	text := fmt.Sprintf("HTTP GET is %s", customerId)
	w.Write([]byte(text))
}
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
func (h *HTTPHandler) GetCustomerCustomerIdOrdersOrderId(
	w http.ResponseWriter,
	r *http.Request,
	customerId string,
	orderId string) {
	order, err := h.app.Queries.GetCustomerOrder.Handle(r.Context(), query.GetCustomerOrder{
		CustomerID: customerId,
		OrderID:    orderId,
	})
	if err != nil {
		slog.Info("Order Web Controller",
			slog.String("Handler Name", "GetCustomerIdOrdersOrderID"),
			slog.Any("error", err),
		)
		writeJSON(w, http.StatusInternalServerError, order)
		return
	}
	slog.Info("Order Web Controller",
		slog.String("Handler Name", "GetCustomerIdOrdersOrderID"),
		slog.Any("ok", "ok"),
	)
	writeJSON(w, http.StatusOK, order)
}
