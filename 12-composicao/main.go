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

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Endereço: %v", filipe.Nome, filipe.Idade, filipe.Ativo, filipe.Endereco)
}
