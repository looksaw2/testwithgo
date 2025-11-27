package main

import (
	"log"

	"github.com/looksaw2/gorder3/internal/common/config"
	"github.com/looksaw2/gorder3/internal/common/server"
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
	server.RunGRPCServer(
		serviceName,
		func(service *grpc.Server) {

		},
	)
}