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

type Empresa struct {
	Nome  string
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado", c.Nome)
}

func (e Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("a empresa %s foi desativada", e.Nome)
}

type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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

	empresa := Empresa{
		Nome: "FGonçalves",
	}

	filipe.Cidade = "Itajaí"
	filipe.Estado = "SC"
	Desativacao(empresa)

	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Endereço: %v", filipe.Nome, filipe.Idade, filipe.Ativo, filipe.Endereco)
	// fmt.Printf("Cliente: %s - Ativo: %t", filipe.Nome, filipe.Ativo)
}
