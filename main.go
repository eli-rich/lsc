package main

import (
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"
)

//go:embed MIT.txt
var mit string

//go:embed APACHE.txt
var apache string

//go:embed BSD2.txt
var bsd2 string

//go:embed BSD3.txt
var bsd3 string

//go:embed GPL.txt
var gpl string

//go:embed MPL.txt
var mpl string

type License struct {
	Name string
	Year string
}

var licenses = map[string]string{
	"MIT":    mit,
	"APACHE": apache,
	"BSD2":   bsd2,
	"BSD3":   bsd3,
	"GPL":    gpl,
	"MPL":    mpl,
}

func main() {
	name := os.Args[1]
	year := os.Args[2]
	selectedLicense := licenses[strings.ToUpper(os.Args[3])]
	license := License{
		Name: name,
		Year: year,
	}
	text := template.Must(template.New("license").Parse(selectedLicense))
	file, err := os.Create("LICENSE")
	panicOn(err)
	defer file.Close()
	text.Execute(file, license)
}

func panicOn(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
