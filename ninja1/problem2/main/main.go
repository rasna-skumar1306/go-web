//problem 2: create a struct that contains details of hotel and region where it is present
package main

import (
	"fmt"
	"os"
	"text/template"
)

type hotel struct { //struct of hotel
	Name    string
	Address string
	City    string
	Zip     int
	Region  region
}

type region struct { //struct of region
	Southern string
	Central  string
	Northern string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	hotels := []hotel{
		{
			Name:    "Ras",
			Address: `1, Main Road, Korat, TN`,
			City:    "Chi",
			Zip:     2313,
			Region:  region{Southern: "TN", Northern: "TN"},
		},
		{
			Name:    "Sen",
			Address: `3, Station Road, Merroot, CN`,
			City:    "Hem",
			Zip:     1234,
			Region:  region{Northern: "CN"},
		},
	}
	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		fmt.Println(err)
	}
}
