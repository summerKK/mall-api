package main

import (
	"net/http"

	"github.com/summerKK/mall-api/global"
	boot "github.com/summerKK/mall-api/init"
	"github.com/summerKK/mall-api/internal/router"
)

func init() {
	boot.SetConfig(nil)
	boot.Boot()
}

func main() {
	r := router.NewRouter()

	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		TLSConfig:      nil,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
