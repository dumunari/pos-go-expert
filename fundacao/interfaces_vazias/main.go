package main

import "fmt"

func main() {
	// Pode ser substituido na maioria dos casos pelos Generics
	var x interface{} = 10
	var y interface{} = "Hello, world"

	showType(x)
	showType(y)

	res, ok := x.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2, ok := x.(string)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res2, ok)

	// No caso de conversão errada, causa Panic pois não tem o bool de check se a conversão deu certo
	// res3 := x.(string)
	// fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res3, res3)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
