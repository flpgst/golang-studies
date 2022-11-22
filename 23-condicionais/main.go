package main

func main() {
	a, b, c := 0, 2, 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b || c > a")
	}

	switch b {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}

}
