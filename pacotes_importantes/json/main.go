package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Nome  string `json:"n"`
	Saldo int    `json:"s"`
}

func main() {
	conta := Conta{
		Nome:  "Teste",
		Saldo: 1000,
	}

	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := `{"n":"Teste2","s":500}`
	var conta2 Conta
	err = json.Unmarshal([]byte(jsonPuro), &conta2)
	if err != nil {
		println(err)
	}
	println(conta2.Nome)
	println(conta2.Saldo)

	// case insensitive
	// jsonPuro2 := `{"nome":"Teste","saldo":1000}`
	// var conta3 Conta
	// err = json.Unmarshal([]byte(jsonPuro2), &conta3)
	// if err != nil {
	// 	println(err)
	// }
	// println(conta3.Nome)
	// println(conta3.Saldo)
}
