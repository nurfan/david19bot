package handler

import (
	"fmt"
	"net/http"
)

// Ping is function for check http health
func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pong\n")
}
