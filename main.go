package main

import (
	"vergil.com/practice/sawmill/app/server"
	"vergil.com/practice/sawmill/app/socket"
)

func main() {

	go socket.DefaultAdmin.StartAdminSocket()
	go server.StartSocketServer()
	go server.StartHTTPServer()

	// DO NOT USE `for {}`!
	// An infinite loop will simply use 100% of a CPU
	// and not allow the Go runtime to schedule
	// any other work on that core.
	select {}
}
