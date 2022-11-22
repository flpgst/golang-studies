package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	filipe := Cliente{
		Nome:  "Filipe",
		Idade: 37,
		Ativo: true,
	}

	filipe.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", filipe.Nome, filipe.Idade, filipe.Ativo)
}
