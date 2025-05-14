package main

import "fmt"

func main() {
	// Variável -> Endereço Memória -> Valor
	a := 10
	println("---Variavel a---")
	fmt.Printf("Valor da var a: %d\n", a)
	fmt.Printf("Endereço da var a: %p\n", &a)
	println("---Variavel a---")
	println()

	// Ponteiro -> Endereço Memória -> Valor
	println("---Variavel ponteiro---")
	var ponteiro *int = &a
	println("Variavel ponteiro recebe endereço da variavel a (&a) ")
	fmt.Printf("Valor da var ponteiro (Endereço de memória da var a): %p\n", ponteiro)
	fmt.Printf("Endereço da var ponteiro (Endereço de memória da variável ponteiro): %p\n", &ponteiro)
	fmt.Printf("Deref da var ponteiro (Valor que o endereço de memória aponta): %d\n", *ponteiro)

	// Valor da var ponteiro
	*ponteiro = 20
	println("Deref da variavel ponteiro recebe 20 ")
	println("---Variavel ponteiro---")
	println()

	// Endereço Memória
	println("---Variavel b---")
	b := &a
	println("Variavel b recebe endereço da variavel a (&a)")
	fmt.Printf("Valor da var b (Endereço de memória da var a): %p\n", b)
	fmt.Printf("Deref da var b (Valor que o endereço de memória aponta): %d\n", *b)

	//Valor do Endereço Memória
	*b = 30
	println("Deref de b (*b) recebe 30")
	fmt.Printf("Deref da var b (Valor que o endereço de memória aponta): %d\n", *b)
	println("---Variavel b---")
	println()

	println("---Variavel a---")
	fmt.Printf("Valor da var a: %d\n", a)
	println("---Variavel a---")

}
