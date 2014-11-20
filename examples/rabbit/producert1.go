package main

import (
	"github.com/stormasm/plum/rabbit"
	"log"
)

var	(
	uri string = "amqp://guest:guest@localhost:5672/"
	exchangeName string = "test-exchange"
	exchangeType string = "direct"
	routingKey string = "test-key"
	body  string = "ralph in socorro"
	reliable bool = true
)

func main() {
	if err := rabbit.Publish(uri, exchangeName, exchangeType, routingKey, body, reliable); err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("published %dB OK", len(body))
}
