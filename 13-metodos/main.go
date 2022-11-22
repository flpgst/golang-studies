package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado", c.Nome)
}

func main() {
	filipe := Cliente{
		Nome:  "Filipe",
		Idade: 37,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 2000",
		},
	}

	filipe.Cidade = "Itajaí"
	filipe.Estado = "SC"
	filipe.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Endereço: %v", filipe.Nome, filipe.Idade, filipe.Ativo, filipe.Endereco)
	fmt.Printf("Cliente: %s - Ativo: %t", filipe.Nome, filipe.Ativo)
}
