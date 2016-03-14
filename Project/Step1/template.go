package main

import (
	"net/http"
	"html/template"
	"log"
)


//Serves index.html
func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}


func main() {
	http.HandleFunc("/", serve_the_webpage)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}