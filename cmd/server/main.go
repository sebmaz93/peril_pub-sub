package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rmqURL = "amqp://guest:guest@localhost:5672/"
)

func main() {
	conn, err := amqp.Dial(rmqURL)
	if err != nil {

	}
	defer conn.Close()
}
