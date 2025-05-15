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
	templates := []string{"content.html", "header.html", "footer.html"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("content.html").ParseFiles(templates...))
		err := tmpl.Execute(w, curso)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
