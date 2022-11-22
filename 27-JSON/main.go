package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Id     int `json:"-"` //oculta propriedade no JSON
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	conta := Conta{Id: 1, Numero: 1, Saldo: 2332}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2,"s":200}`)

	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)

}
