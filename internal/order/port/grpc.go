package port

import (
	"context"

	"github.com/looksaw2/gorder3/internal/common/genproto/orderpb"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)


type GRPCHandler struct {
	orderpb.UnimplementedOrderServiceServer
}


func NewGRPCHandler() *GRPCHandler {
	return &GRPCHandler{}
}


func(h *GRPCHandler)CreateOrder(ctx context.Context,req *orderpb.CreateOrderRequest) (*emptypb.Empty, error){
	return nil , nil
}


func(h *GRPCHandler)GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.Order, error){
	return nil , nil
}