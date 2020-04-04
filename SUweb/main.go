package main

import (
	"SUweb/conf"
	"SUweb/routers"
	"fmt"
	"net/http"
)

func main() {
	routers := routers.InitRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: routers,
	}
	fmt.Println(conf.Port)
	server.ListenAndServe()
}
