package main

import "fmt"

func main() {
	fmt.Println(sum(5, 10, 43, 43, 435, 5451, 1))

}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
