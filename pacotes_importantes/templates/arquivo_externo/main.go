package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := []Curso{{"Go", 40}, {"Python", 30}, {"Java", 50}}
	tmpl := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmpl.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
