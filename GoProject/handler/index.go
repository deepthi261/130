package handler

import (
	"github.com/deepthi261/GoProject/log"
	"html/template"
	"net/http"
)

// The main page handler
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/index.html"))
	err := tpl.Execute(res, GetAPlusTemplateHeader(req, nil))
	log.LogError(err)
}
