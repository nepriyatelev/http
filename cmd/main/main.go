package main

import (
	"cmd/cmd/handler"
	"net/http"
)

func main() {

	p := handler.Person{}
	http.HandleFunc("/", p.PostHandler)
	http.ListenAndServe(":8080", nil)
}
