package main

import (
	"log"

	"github.com/nicklas.17/hermeslens/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main () {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.KafkaBroker),
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
	defer client.Close()
}
