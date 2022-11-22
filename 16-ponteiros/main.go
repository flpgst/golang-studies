package main

func soma(a, b *int) int {
	*a = 20
	*b = 20
	return *a + *b
}

func main() {
	var elem1 = 10
	var elem2 = 5

	soma(&elem1, &elem2)

	println(elem1)
	println(elem2)

}
