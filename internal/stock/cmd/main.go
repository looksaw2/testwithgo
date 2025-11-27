package main

import (
	"context"
	"log"

	"github.com/looksaw2/gorder3/internal/common/config"
	"github.com/looksaw2/gorder3/internal/common/genproto/stockpb"
	"github.com/looksaw2/gorder3/internal/common/server"
	"github.com/looksaw2/gorder3/internal/stock/port"
	"github.com/looksaw2/gorder3/internal/stock/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err) 
	}
}


func main(){
	serviceName := viper.GetString("stock.service-name")
	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	server.RunGRPCServer(
		serviceName,
		func(service *grpc.Server) {
			svc := port.NewGRPCHandler(application)
			stockpb.RegisterStockServiceServer(service,svc)
		},
	)
}