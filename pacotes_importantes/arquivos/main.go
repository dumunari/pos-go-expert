package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	// se string, pode usar -> tamanho, err := f.WriteString("Escrevendo dados no arquivo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! O tamanho do arquivo é %d\n", tamanho)
	defer f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println((string(arquivo)))

	// leitura com buffer
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
