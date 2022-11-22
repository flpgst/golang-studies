package main

import "fmt"

func main() {
	salarios := map[string]int{"Filipe": 2000, "João": 2000}
	delete(salarios, "Filipe")
	salarios["Claudio"] = 4000

	// salario := make(map[string]int)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	for nome := range salarios {
		fmt.Printf("O salario é do %v\n", nome)
	}
}
