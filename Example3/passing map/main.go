//program to send a map as value to a template
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
	family := map[string]string{ //map variable
		"Dad":             "Senthilkumar",
		"Mom":             "Hema Senthilkumar",
		"Myself":          "Rasswanth",
		"Younger Brother": "Rithish Balaji",
	}
	err := tpl.Execute(os.Stdout, family) //sending map as value
	if err != nil {
		log.Fatalln(err)
	}
}
