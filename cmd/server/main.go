package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
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
	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("error oppening RMQ channel: %v", err)
	}
	err = pubsub.PublishJSON(publishCh, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{
		IsPaused: true,
	})
	if err != nil {
		log.Printf("could not publish: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("the program is shutting down")
}
