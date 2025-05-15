package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := []Curso{{"Go", 40}, {"Python", 30}, {"Java", 50}}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmpl.Execute(w, curso)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
