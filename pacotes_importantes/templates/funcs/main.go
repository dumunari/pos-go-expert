package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func AddStar(str string) string {
	return fmt.Sprintf("%s *", str)
}

func main() {
	curso := []Curso{{"Go", 40}, {"Python", 30}, {"Java", 50}}
	templates := []string{"content.html", "header.html", "footer.html"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.New("content.html")
		tmpl.Funcs(template.FuncMap{"AddStar": AddStar})
		tmpl = template.Must(tmpl.ParseFiles(templates...))
		err := tmpl.Execute(w, curso)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
