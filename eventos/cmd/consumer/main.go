package main

import (
	"eventutils/rabbitmq"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "events")

	for msg := range msgs {
		fmt.Println("Received message:", string(msg.Body))
		msg.Ack(false)
	}
}
