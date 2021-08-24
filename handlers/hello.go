package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// simple handler
type Hello struct {
	l *log.Logger
}

// create a new Hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello requests")

	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Unalbe to read request body", http.StatusBadRequest)
		return
	}

	// response
	fmt.Fprintf(rw, "Hello %s", d)
}
