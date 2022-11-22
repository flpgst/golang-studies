package main

import (
	"estudosGo/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado:", s)
	uid := uuid.NewString()
	println(uid)
}
