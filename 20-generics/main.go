package main

type NumberInt int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{"Filipe": 1000, "Maria": 1500, "Joao": 450}
	m2 := map[string]float64{"Filipe": 1000.00, "Maria": 1500.30, "Joao": 450.99}
	m3 := map[string]NumberInt{"Filipe": 1000, "Maria": 1500, "Joao": 450}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	Compara(10, 100)

}
