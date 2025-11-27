package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/looksaw2/gorder3/internal/common/client/order"
	"github.com/looksaw2/gorder3/internal/common/config"
	"github.com/looksaw2/gorder3/internal/common/genproto/orderpb"
	"github.com/looksaw2/gorder3/internal/common/server"
	"github.com/looksaw2/gorder3/internal/order/port"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//初始化
func init(){
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

//主要函数
func main() {
	//使用GRPC代码
	go server.RunGRPCServer(
		viper.GetString("order.service-name"),
			func(service *grpc.Server) {
				svc := port.NewGRPCHandler()
				orderpb.RegisterOrderServiceServer(service,svc)
			},
	)
	//使用HTTP的代码
	h := &HTTPHandler{}
	server.RunHTTPServer(
		viper.GetString("order.service-name"),
		func(router chi.Router) {
			order.HandlerFromMux(h,router)
		},
	)
}