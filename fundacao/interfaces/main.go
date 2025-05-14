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

type Empresa struct {
	Nome  string
	CNPJ  string
	Ativa bool
}

type Pessoa interface {
	Desativar()
}

type Negocio interface {
	DesativarCNPJ()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s, de %d anos, foi desativado\n", c.Nome, c.Idade)
}

func (v Vendedor) Desativar() {
	v.Ativo = false
	fmt.Printf("O vendedor %s, de %d anos, foi desativado\n", v.Nome, v.Idade)
}

func (e Empresa) DesativarCNPJ() {
	e.Ativa = false
	fmt.Printf("A empresa %s, com CNPJ %s, foi desativada\n", e.Nome, e.CNPJ)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func DesativacaoNegocio(negocio Negocio) {
	negocio.DesativarCNPJ()
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

	empresa := Empresa{
		Nome:  "Empresta Teste",
		CNPJ:  "12918291829182",
		Ativa: true,
	}

	cliente.Cidade = "Teste City"
	vendedor.Endereco.Cidade = "Teste City 2"

	Desativacao(cliente)
	Desativacao(vendedor)
	DesativacaoNegocio(empresa)
}
