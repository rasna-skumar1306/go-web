package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, `Rasswanth Senthilkumar S/o Senthilkumar`)
	if err != nil {
		log.Fatalln(err)
	}
}
