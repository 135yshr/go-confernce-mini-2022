package handlers

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
