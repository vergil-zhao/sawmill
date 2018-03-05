package socket

import (
	"encoding/json"
	"log"

	"vergil.com/practice/sawmill/app/model"
)

// ByteDataSeperator Seperator of byte data "\x01"
const ByteDataSeperator byte = 1

// Decoder decode binary log data to string
type Decoder struct {
	data     ByteData
	Delegate DecoderDelegate
}

// DecoderDelegate will be called when a complete log received
type DecoderDelegate interface {
	LogDecoded(*model.Device)
}

// Decode decode raw data
func (d *Decoder) Decode(b []byte) {
	data := ByteData(b)
	copy(d.data, data)
	if data.Contains(1) {
		pieces := d.data.Split(1)
		d.data = []byte{}
		for _, piece := range pieces {
			log.Println(piece)
			log.Println(string(piece))
			log := new(model.Device)
			json.Unmarshal(piece, log)
			d.Delegate.LogDecoded(log)
		}
	}
}

// ByteData customized []byte for split data
type ByteData []byte

// Contains test if a ByteData contains a byte
func (b ByteData) Contains(number byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == number {
			return true
		}
	}
	return false
}

// Split split b into pieces by a byte
func (b ByteData) Split(number byte) []ByteData {
	last := 0
	var result []ByteData
	for i := 0; i < len(b); i++ {
		if b[i] == number {
			result = append(result, b[last:i])
			last = i + 1
		}
	}
	return result
}
