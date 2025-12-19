package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rmqConnStr = "amqp://guest:guest@localhost:5672/"
)

func main() {
	conn, err := amqp.Dial(rmqConnStr)
	if err != nil {
		log.Fatalf("cannot open RMQ connection: %v", err)
	}
	defer conn.Close()
	log.Print("connected to RMQ")
	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Printf("error welcoming: %v", err)
	}
	log.Print(username)
	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, username)
	_, queue, err := pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, queueName, routing.PauseKey, pubsub.QueueTransient)
	if err != nil {
		log.Fatalf("could not subscribe to pause: %v", err)
	}
	log.Printf("Queue %v declared and bound!\n", queue.Name)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("RMQ connection closed")
}
