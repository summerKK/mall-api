package main

import (
	"net/http"

	"github.com/summerKK/go-code-snippet-library/koel-api/global"
	boot "github.com/summerKK/go-code-snippet-library/koel-api/init"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/router"
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
