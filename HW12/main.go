package main

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"io"
	"strings"
)

func main(){
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		cookie, err := req.Cookie("session-id")
		if err == http.ErrNoCookie{
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-id",
				Value: id.String(),
				HttpOnly:true,
			}
		}

		if(req.FormValue("name") != "" && !strings.Contains(cookie.Value, "name")){
			cookie.Value = cookie.Value + `name= ` + req.FormValue("name")
		}

		http.SetCookie(res, cookie)
		io.WriteString(res,`<!DOCTYPE html>
		<html>
		  <body>
		    <form method="POST">
		    `+cookie.Value+`
		      <br/>
		      <input type="text" name="name">
		      <input type="submit">
		    </form>
		  </body>
		</html>`)
	})
	http.ListenAndServe(":8080",nil)
}