package http

import (
	"net/http"

	"vergil.com/practice/sawmill/app/handler"
)

// SetupRoute setting routes here
func SetupRoute() {
	http.Handle("/log", &handler.Logger{})
}
