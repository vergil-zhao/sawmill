package server

import (
	"net"

	"vergil.com/practice/sawmill/app/handler"

	"vergil.com/practice/sawmill/app/socket"
)

// StartSocketServer start to receive logs on socket
func StartSocketServer() {
	socket.DefaultUser.StartSocketLogger(func(conn net.Conn) {
		read := true
		data := make([]byte, 4096)
		decoder := new(socket.Decoder)
		handler.NewSocketLogger(decoder)
		for read {
			count, err := conn.Read(data)
			read = (err == nil)
			decoder.Decode(data[:count])
		}
		conn.Close()
	})
}
