package main

import (
	"vergil.com/practice/sawmill/app/server"
	"vergil.com/practice/sawmill/app/socket"
)

func main() {

	go socket.DefaultAdmin.StartAdminSocket()
	go server.StartSocketServer()
	go server.StartHTTPServer()

	for {
	}
}
