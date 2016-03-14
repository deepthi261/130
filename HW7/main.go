
package main

import (
	"net/http"
	"io"
)

func main(){
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		key := "q"
		val := req.FormValue(key)
		io.WriteString(res, "Query String - Name :"+val)
	})

	http.ListenAndServe(":8080",nil)
}
