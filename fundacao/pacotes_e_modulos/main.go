package main

import (
	"fmt"
	"fundacao/matematica"

	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Println(uuid.New())
	fmt.Println("Resultado: ", soma)
}
