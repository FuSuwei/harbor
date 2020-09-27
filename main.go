package main

import (
	"fmt"
	"net/http"

	"harbor/routers"

	"harbor/pkg/toml"
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

