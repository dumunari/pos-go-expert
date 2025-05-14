package main

import "fmt"

type ID int

var (
	b bool
	c int
	d string
)

func main() {
	var e ID = 1

	fmt.Printf("O tipo de b é %T\n", b)
	fmt.Printf("O tipo de c é %T\n", c)
	fmt.Printf("O tipo de d é %T\n", d)
	fmt.Printf("O tipo de e é %T\n", e) // mostra onde foi criado o tipo

	fmt.Printf("O valor de b é %v\n", b)
	fmt.Printf("O valor de c é %v\n", c)
	fmt.Printf("O valor de d é %v\n", d)
}
