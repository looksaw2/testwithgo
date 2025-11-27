package server

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

// 一个通用的使用http的代码
func RunHTTPServer(servicename string, wrapper func(router chi.Router)) {
	//得到服务的HTTP的端口
	addr := viper.Sub(servicename).GetString("http-addr")
	if addr == "" {
		slog.Error("service Run http server",
			slog.String("service name", servicename),
			slog.String("on Port ", addr),
		)
		return
	}
	slog.Info("service Run http server",
		slog.String("service name", servicename),
		slog.String("on Port ", addr),
	)
	//运行OnAddr
	RunHTTPServerOnAddr(addr, wrapper)
}

func RunHTTPServerOnAddr(addr string, wrapper func(router chi.Router)) {
	apiRouter := chi.NewMux()
	apiRouter.Route("/api", func(r chi.Router) {
		wrapper(r)
	})
	slog.Info("Start to Run ...........")
	if err := http.ListenAndServe(addr, apiRouter); err != nil {
		log.Fatal(err)
	}
}
