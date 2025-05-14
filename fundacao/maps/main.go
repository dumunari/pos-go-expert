package main

import "fmt"

func main() {
	salarios := map[string]int{"Teste": 1000, "Teste2": 2000}
	fmt.Println(salarios["Teste"])
	fmt.Println(salarios)

	delete(salarios, "Teste")
	fmt.Println(salarios)

	salarios["Teste"] = 3000
	fmt.Println(salarios["Teste"])
	fmt.Println(salarios)

	for nome, valor := range salarios {
		fmt.Printf("O salário do %s é %d\n", nome, valor)
	}

	total := 0

	for _, valor := range salarios {
		fmt.Printf("O salário é %d\n", valor)
		total += valor
	}
	fmt.Printf("A soma dos salários é %d\n", total)
}
