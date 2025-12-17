package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rmqURL = "amqp://guest:guest@localhost:5672/"
)

func main() {
	conn, err := amqp.Dial(rmqURL)
	if err != nil {
		log.Fatalf("couldn't start RMQ connection: %v", err)
	}
	defer conn.Close()
	log.Println("connection successfull")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("the program is shutting down")
}
