package main

import "eventutils/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	body := []byte("Hello, RabbitMQ!")
	if err := rabbitmq.Publish(ch, body, "amq.direct"); err != nil {
		panic(err)
	}
}
