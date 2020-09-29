package main

import (
	"fmt"
	"harbor/pkg/toml"
	"harbor/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	setting := toml.GetServerConfig()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
