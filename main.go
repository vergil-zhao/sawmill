package main

import (
	"vergil.com/practice/sawmill/app/http"
	"vergil.com/practice/sawmill/app/socket"
)

func main() {

	go socket.DefaultAdmin.StartAdminSocket()

	go http.StartHTTPServer()

	for {
	}
}
