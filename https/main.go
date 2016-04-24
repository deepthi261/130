package main

import (
	"net/http"
	"io"
)

func httpsReq(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "HEY, WHATS UP MAN ITS HTTPS !!")
}

func main() {
	http.HandleFunc("/", httpsReq)
	http.ListenAndServeTLS("10443", "../cert.pem", "../key.pem", nil)
}