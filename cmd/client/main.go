package main

import (
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
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
	wlcmMsg, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Printf("error welcoming: %v", err)
	}
	log.Print(wlcmMsg)
}
