package socket

import (
	"log"
	"net"

	"vergil.com/practice/sawmill/app/config"
)

// User socket users who send logs
type User struct{}

// DefaultUser for listen to new logs
var DefaultUser = new(User)

// StartSocketLogger start to record logs on socket
func (u *User) StartSocketLogger(block func(conn net.Conn)) {
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
		go block(conn)
	}
}
