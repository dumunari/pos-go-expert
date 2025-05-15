package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHorario int
}

func main() {
	curso := Curso{"Go", 40}
	tmpl := template.New("CursoTemplate")
	tmpl, _ = tmpl.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHorario}} horas")
	err := tmpl.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
