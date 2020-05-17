//program that sends slice to 2 different templates
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml") //first template
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, []string{"Senthilkumar", "Hema senthilkumar", "Rasswanth", "Rithish Balaji"}) //passing []string as data
	if err != nil {
		log.Fatalln(err)
	}
	tp, err := template.ParseFiles("index2.gohtml") //second template
	if err != nil {
		log.Fatalln(err)
	}
	err = tp.Execute(os.Stdout, []string{"Senthilkumar", "Hema senthilkumar", "Rasswanth", "Rithish Balaji"})
	if err != nil {
		log.Fatalln(err)
	}
}
