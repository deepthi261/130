package main

import (
	"io"
	"net/http"
	"os"
	"log"
)

func main() {
	http.HandleFunc("/", fileHandling)
	http.ListenAndServe(":8080", nil)
}

func fileHandling (res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		key := "text"
		_, hdr, err := req.FormFile(key)
		if err != nil {
			log.Println("Error : ", err)
		}
		old,_ := os.Open(hdr.Filename)
		io.Copy(res, old)
	} else {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
		  <input type="file" name="text">
		  <input type="submit">
		</form>`)
	}
}
