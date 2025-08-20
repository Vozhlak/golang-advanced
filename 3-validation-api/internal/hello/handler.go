package hello

import (
	"fmt"
	"net/http"
)

type HandlerHello struct{}

func NewHandlerHello(router *http.ServeMux) {
	handler := &HandlerHello{}
	router.HandleFunc("/api/hello", handler.hello())
}

func (handler *HandlerHello) hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Println("Hello")
	}
}
