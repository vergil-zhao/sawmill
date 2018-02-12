package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	for i := 0; i < 100; i++ {
		go func() {
			conn.Write([]byte(`{"UUID":"123","Model":"iphone8,2","Logs":[{"Level":0,"Message":"Just a fresh log.","File":"log.swift","Line":100,"Time":"2018-02-08T13:41:05+08:00"}]}`))
		}()
	}

	read := true
	data := make([]byte, 4096)
	for read {
		count, err := conn.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}
	conn.Close()
}
