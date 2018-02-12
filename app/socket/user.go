package socket

import (
	"fmt"
	"log"
	"net"

	"vergil.com/practice/sawmill/app/config"
)

// User socket users who send logs
type User struct{}

// DefaultUser for listen to new logs
var DefaultUser = new(User)

// StartSocketLogger start to record logs on socket
func (u *User) StartSocketLogger() {
	l, e := net.Listen("tcp", config.UserSocket)
	if e != nil {
		log.Println("Listen to user socket failed:", e)
	}
	log.Println("Start to receive logs on:", config.UserSocket)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Accept new user failed:", e)
		}
		go u.acceptLog(conn)
	}
}

func (u *User) acceptLog(conn net.Conn) {
	read := true
	data := make([]byte, 4096)
	for read {
		count, err := conn.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}
	conn.Close()
}
