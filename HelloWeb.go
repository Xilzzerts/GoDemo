package main

import (
	"fmt"
	"log"
	"net/http"
)

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	name := "World"
	if val, ok := request.Form["name"]; ok {
		name = val[0]
	}
	fmt.Fprint(writer, "Hello ", name)
}

func main() {
	http.HandleFunc("/", MainHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
