package main

import "fmt"

func main() {
	total := func() int {
		return sum(5, 10, 43, 43, 435, 5451, 1)
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
