//example 2 for passing value into template
package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template //declaring a *template.Template for parsing file

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml")) //using template.Must for handling err and getting tpl *template.Template
}

func main() {
	err := tpl.Execute(os.Stdout, `Rasswanth Senthilkumar S/o Senthilkumar`) //passing a raw string literal as value
	if err != nil {
		log.Fatalln(err)
	}
}
