package main

import (
	"fmt"
)

// canal send-only
// ou seja, vc apenas envia para ele, nao recebe nada dele
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// canal receive-only
// ou seja, vc apenas recebe valores dele, nao envia nada para ele
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	go recebe("World", hello)
	ler(hello)
}
