package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 3, 5, 6, 78, 64, 4, 3, 3, 3) * 2
	}()
	fmt.Println(total)

	fmt.Println((sum(1, 3, 5, 6, 78, 64, 4, 3, 3, 3)))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
