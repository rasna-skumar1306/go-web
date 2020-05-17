//problem 3: create a datastructure that contains the info of menu card of hotels specified.
package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Name      string
	Breakfast items
	Lunch     items
	Dinner    items
}

type items struct {
	Item map[int]string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		Name: "CrazyRas@in",
		Breakfast: items{
			Item: map[int]string{
				1: "Dosa",
				2: "Idly",
				3: "Poori",
				4: "Pongal",
				5: "Vadai",
			},
		},
		Lunch: items{
			Item: map[int]string{
				1: "Sambhar rice",
				2: "Curd rice",
				3: "Biriyani",
				4: "Naan",
				5: "Fried rice",
			},
		},
		Dinner: items{
			Item: map[int]string{
				1: "Parota",
				2: "Idly",
				3: "Dosa",
				4: "Chapathi",
				5: "Idiyappam",
			},
		},
	}
	err := tpl.Execute(os.Stdout, m) //passing the DS as value
	if err != nil {
		log.Fatal(err)
	}
}
