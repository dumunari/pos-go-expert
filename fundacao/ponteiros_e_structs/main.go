package main

import "fmt"

type Conta struct {
	saldo int
}

func (c Conta) simular(valor int) {
	c.saldo += valor
	fmt.Printf("Saldo simulado: %d\n", c.saldo)
}

func (c *Conta) aplicar(valor int) {
	c.saldo += valor
	fmt.Printf("Saldo aplicado: %d\n", c.saldo)
}

func main() {
	conta := Conta{
		saldo: 100,
	}

	conta.simular(300)
	fmt.Printf("Saldo da conta: %d\n", conta.saldo)

	conta.aplicar(400)
	fmt.Printf("Saldo da conta: %d\n", conta.saldo)

}
