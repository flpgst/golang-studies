package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: Task %d is running\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s: Task %d is running\n", "anonymous", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
