package socket

import (
	"encoding/json"
	"log"
	"net"

	"vergil.com/practice/sawmill/app/config"
)

// Admin admin object contains a tcp connection
type Admin struct {
	connections []net.Conn
}

// DefaultAdmin default admin
var DefaultAdmin = new(Admin)

// StartAdminSocket start to listen
func (a *Admin) StartAdminSocket() {
	l, e := net.Listen("tcp", config.AdminSocket)
	if e != nil {
		log.Println("Listen to", config.AdminSocket, "failed:", e)
		return
	}
	log.Println("Start to listen Admin on", config.AdminSocket)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Accept new admin error:", err)
		}
		a.connections = append(a.connections, conn)
	}
}

// Send data to admin
func (a *Admin) Send(o interface{}) error {
	if len(a.connections) == 0 {
		return nil
	}

	b, e := json.Marshal(o)
	if e != nil {
		return e
	}

	for _, conn := range a.connections {
		conn.Write(b)
	}

	return nil
}
