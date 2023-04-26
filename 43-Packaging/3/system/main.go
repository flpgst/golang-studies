package main

import (
	"fmt"

	"github.com/flpgst/estudosGo/43-Packaging/2/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
