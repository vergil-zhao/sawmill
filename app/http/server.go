package http

import (
	"log"
	"net/http"

	"vergil.com/practice/sawmill/app/config"
)

//StartHTTPServer start http server
func StartHTTPServer() {
	SetupRoute()
	log.Println("Start to listen HTTP on", config.HTTP)
	err := http.ListenAndServe(config.HTTP, nil)
	if err != nil {
		log.Fatal("Serve HTTP failed:", err)
	}
}
