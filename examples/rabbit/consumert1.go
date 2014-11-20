package main

import (
	"github.com/stormasm/plum/rabbit"
	"log"
	"time"
)

var (
	uri          string        = "amqp://guest:guest@localhost:5672/"
	exchange     string        = "test-exchange"
	exchangeType string        = "direct"
	queue        string        = "test-queue"
	bindingKey   string        = "test-key"
	consumerTag  string        = "simple-consumer"
	lifetime     time.Duration = 60 * time.Second
)

func main() {
	c, err := rabbit.NewConsumer(uri, exchange, exchangeType, queue, bindingKey, consumerTag)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if lifetime > 0 {
		log.Printf("running for %s", lifetime)
		time.Sleep(lifetime)
	} else {
		log.Printf("running forever")
		select {}
	}

	log.Printf("shutting down")

	if err := c.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}
