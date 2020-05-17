//program to pass a value into a template
package main

import (
	"log"
	"os"
	"text/template" //template package for parsing files
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml") //parse a gohtml file and getting tpl of type *template.Template
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, "Rasswanth") //passing data into template
	if err != nil {
		log.Fatalln(err)
	}
}
