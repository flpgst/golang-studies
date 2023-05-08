package main

import "fmt"

func main() {
	evento := []string{"Evento 1", "Evento 2", "Evento 3"}
	evento = append(evento[:0], evento[1:]...)
	fmt.Println(evento)
}

// remove(3)

// [1,2,3,4,5]
