package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/looksaw2/gorder3/internal/common/client/order"
	"github.com/looksaw2/gorder3/internal/common/config"
	"github.com/looksaw2/gorder3/internal/common/server"
	"github.com/spf13/viper"
)

//初始化
func init(){
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

//主要函数
func main() {
	h := &HTTPHandler{}
	server.RunHTTPServer(
		viper.GetString("order.service-name"),
		func(router chi.Router) {
			order.HandlerFromMux(h,router)
		},
	)
}