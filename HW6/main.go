package main

import (
	"io"
	"net/http"
	"strings"
	"log"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := strings.Split(req.URL.Path, "/")
		log.Println(name)
		io.WriteString(res, "Name: " + name[1])

	})
	http.ListenAndServe(":8080", nil)
}
