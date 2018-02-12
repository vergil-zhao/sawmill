package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
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
