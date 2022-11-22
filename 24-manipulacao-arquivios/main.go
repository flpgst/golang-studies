package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Cliente struct {
	id   int
	nome string
}

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hellpe")

	type Conta struct {
		Numero int
		Saldo  int
	}

	conta := Conta{Numero: 1, Saldo: 2332}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	tamanho, err := f.Write([]byte(res))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes", tamanho)

	f.Close()

	// leitura

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura em buffer
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))

	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
