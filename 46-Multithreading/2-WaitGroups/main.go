package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: Task %d is running\n", name, i)
		time.Sleep(1 * time.Second)
		wg.Done()

	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(30)

	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s: Task %d is running\n", "anonymous", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
