package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (hello *HelloHandler) Hello() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {
		fmt.Println("Hello")
	}
}
