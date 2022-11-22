package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err, b := sum(5, 10)

	if err != nil {
		fmt.Println(err)
	}

	if b {
		fmt.Println(b)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error, bool) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50"), false
	}

	return a + b, nil, true

}
