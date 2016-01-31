package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			Name: "deepthi thotakura",
			Age:  23,
		},
		LicenseToKill: false,
	}

	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}