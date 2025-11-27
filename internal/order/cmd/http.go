package main

import (
	"fmt"
	"net/http"
)

type HTTPHandler struct{}


func(h *HTTPHandler)PostCustomerCustomerIdOrders(
	w http.ResponseWriter, 
	r *http.Request, 
	customerId string){
	text := fmt.Sprintf("HTTP GET is %s",customerId)
	w.Write([]byte(text))
}


func(h *HTTPHandler)GetCustomerCustomerIdOrdersOrderId(
	w http.ResponseWriter, 
	r *http.Request, 
	customerId string, 
	orderId string){
	w.Write([]byte("Hello World"))
}