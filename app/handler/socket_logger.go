package handler

import (
	"fmt"

	"vergil.com/practice/sawmill/app/model"
	"vergil.com/practice/sawmill/app/socket"
)

// SocketLogger decode raw log data from socket
type SocketLogger struct {
}

// NewSocketLogger create a new logger with a decoder
func NewSocketLogger(d *socket.Decoder) *SocketLogger {
	logger := new(SocketLogger)
	d.Delegate = logger
	return logger
}

// LogDecoded interface method from DecoderDelegate
func (l *SocketLogger) LogDecoded(log *model.Device) {
	fmt.Println(log)
}
