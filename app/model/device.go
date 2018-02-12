package model

import (
	"bytes"
)

// Device container for logs
type Device struct {
	UUID  string
	Model string
	Logs  []Log
}

func (d *Device) String() string {
	b := bytes.Buffer{}
	b.WriteString("[" + d.Model + ":" + d.UUID + "]")
	for _, l := range d.Logs {
		b.WriteString(l.String())
	}
	return b.String()
}
