package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Erro ao ler o corpo da resposta: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Printf("Erro ao fazer o Unmarshal: %v\n", err)
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Printf("Erro ao criar o arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("O CEP %s é da cidade %s, estado %s\n", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade:", data.Localidade)
	}
}
