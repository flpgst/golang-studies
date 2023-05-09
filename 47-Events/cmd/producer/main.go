package main

import "github.com/flpgst/golang-studies/47-Events/pkg/events/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "hello world", "amq.direct")
}
