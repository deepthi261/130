package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// Registering the URL path and binding it to userNameForm handler
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Take this!")
	})

	// Setting the listener for http on port 80
	log.Println("Listening to 80 ...")
	go http.ListenAndServe(":80", http.RedirectHandler("https://localhost:443/", http.StatusMovedPermanently))

	// Setting the listener for https on port 443
	log.Println("Listening to 443 ...")
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	logError(err)
}

// Logs the error given into log
func logError(err error) {

	if err != nil {
		log.Println(err)
	}
}