package main

import "fmt"

func main() {

	var minhaVar interface{} = "Filipe"
	println(minhaVar.(string))
	res, ok := minhaVar.(string)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
	res2, _ := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v e ok", res2)
}
