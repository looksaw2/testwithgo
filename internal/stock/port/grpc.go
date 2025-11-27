package port

import (
	"context"

	"github.com/looksaw2/gorder3/internal/common/genproto/stockpb"
	"github.com/looksaw2/gorder3/internal/stock/app"
)


type GRPCHandler struct {
	//依赖注入
	app app.Application
	stockpb.UnimplementedStockServiceServer
}


func NewGRPCHandler(app app.Application) *GRPCHandler {
	return &GRPCHandler{
		app: app,
	}
}

func(h *GRPCHandler)GetItems(context.Context, *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error){
	return  nil ,nil	
}


func(h *GRPCHandler)CheckIfItemInStock(context.Context, *stockpb.CheckIfItemInStockRequest) (*stockpb.CheckIfItemInSTockResponse, error){
	return nil ,nil
}