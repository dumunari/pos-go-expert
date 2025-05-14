package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Vendedor struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s, de %d anos, foi desativado\n", c.Nome, c.Idade)
}

func (v Vendedor) Desativar() {
	v.Ativo = false
	fmt.Printf("O vendedor %s, de %d anos, foi desativado\n", v.Nome, v.Idade)
}

func main() {
	cliente := Cliente{
		Nome:  "Cliente Teste",
		Idade: 30,
		Ativo: true,
	}

	vendedor := Vendedor{
		Nome:  "Vendedor Teste",
		Idade: 35,
		Ativo: true,
	}

	cliente.Cidade = "Teste City"
	vendedor.Endereco.Cidade = "Teste City 2"

	cliente.Desativar()
	vendedor.Desativar()

}
