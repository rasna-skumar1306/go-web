//program that sends struct as a value to template
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

type s struct {
	Name string
	Role string
}

func main() {
	family := []s{
		{Name: "Senthilkumar",
			Role: "Father"},
		{Name: "Hema Senthilkumar",
			Role: "Mother"},
	}

	err := tpl.Execute(os.Stdout, family) //passing a struct as value
	if err != nil {
		log.Fatalln(err)
	}
}
