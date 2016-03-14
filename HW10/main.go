package main

import (
	"net/http"
	"strconv"
	"io"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.NotFound(res, req)
			return
		}

		cookie, err := req.Cookie("cookie-counter")
		if err == http.ErrNoCookie{
			cookie = &http.Cookie{
				Name: "cookie-counter",
				Value: "0",
			}
		}

		count, _ := strconv.Atoi(cookie.Value)
		count++
		cookie.Value = strconv.Itoa(count)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)
	})
	http.ListenAndServe(":8080",nil)
}