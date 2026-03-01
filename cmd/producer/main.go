package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/nicklas-17/hermeslens/config"
	"github.com/nicklas-17/hermeslens/internal/models"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {

	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.KafkaBroker),
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
	defer client.Close()

	admin := kadm.NewClient(client)
	_, err = admin.CreateTopics(context.Background(), 1, 1, nil, config.FlightPricesTopic)
	if err != nil {
		log.Printf("topic may already exist, continuing: %v", err)
	} else {
		log.Printf("topic '%s' ready", config.FlightPricesTopic)
	}

	routes := []struct {
		origin      string
		destination string
		airline     string
	}{
		{"JFK", "LAX", "Delta"},
		{"CDG", "LHR", "Air France"},
		{"DXB", "SIN", "Emirates"},
	}

	ticker := time.NewTicker(10 * time.Second)
	log.Println("Producer started. Sending price events every 10 seconds...")

	for range ticker.C {

		route := routes[rand.Intn(len(routes))]


		event := models.FlightPriceEvent{
			ID:            generateID(),
			Origin:        route.origin,
			Destination:   route.destination,
			Price:         randomPrice(200, 1200),
			Currency:      "EUR",
			Airline:       route.airline,
			DepartureDate: time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
			FetchedAt:     time.Now().UTC().Format(time.RFC3339),
		}


		payload, err := json.Marshal(event)
		if err != nil {
			log.Printf("failed to marshal event: %v", err)
			continue 
		}

		record := &kgo.Record{
			Topic: config.FlightPricesTopic,
			Key:   []byte(event.Origin + "-" + event.Destination),
			Value: payload,
		}

		if err := client.ProduceSync(context.Background(), record).FirstErr(); err != nil {
			log.Printf("failed to produce record: %v", err)
			continue
		}

		log.Printf("produced → %s-%s | %s | $%.2f",
			event.Origin, event.Destination, event.Airline, event.Price)
	}
}


func generateID() string {
	return time.Now().Format("20060102150405.999999999")
}

func randomPrice(min, max float64) float64 {
	price := min + rand.Float64()*(max-min)
	return float64(int(price*100)) / 100
}
