package main

import "fmt"

func publish(name string, ch chan<- string) {
	ch <- name
}

func read(ch chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go publish("hello", ch)
	read(ch)
}
