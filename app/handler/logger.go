package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"vergil.com/practice/sawmill/app/config"
	"vergil.com/practice/sawmill/app/connection"
	"vergil.com/practice/sawmill/app/socket"

	"vergil.com/practice/sawmill/app/model"
)

// Logger handle "/log" uri for receiving logs
type Logger struct{}

// ServeHTTP implement http.Handler interface
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Check http request method
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Check if the body is empty
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	// Decode request body
	d := new(model.Device)
	e := json.NewDecoder(r.Body).Decode(d)
	if e != nil {
		http.Error(w, "Invalid json data.", 400)
		log.Println(e)
		return
	}

	// Write to database
	if config.LogSave {
		save(d)
	}

	// Send to admin
	socket.DefaultAdmin.Send(d)

	// Success response
	w.Write([]byte("OK"))
	log.Println(d)

}

type query map[string]interface{}

func save(d *model.Device) error {
	devices := connection.DB().C("devices")
	logs := connection.DB().C("logs")

	// Save device info if not exist
	_, e := devices.Upsert(query{"uuid": d.UUID}, query{
		"uuid":  d.UUID,
		"model": d.Model,
	})
	if e != nil {
		return e
	}

	// Get object id
	var id map[string]interface{}
	e = devices.Find(query{"uuid": d.UUID}).One(&id)
	log.Println(id)
	if e != nil {
		return e
	}

	// Save log and reference to device
	for _, l := range d.Logs {
		e = logs.Insert(query{
			"level":   l.Level,
			"message": l.Message,
			"file":    l.File,
			"line":    l.Line,
			"time":    l.Time,
			"device":  id["_id"],
		})
		if e != nil {
			return e
		}
	}

	return nil
}
