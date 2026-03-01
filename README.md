# HermesLens

A real-time flight price tracker built with Go and Kafka. The idea is simple — flight prices change constantly and most people miss deals because they're not watching at the right time. HermesLens watches for you.

I built this to get hands-on with event streaming and Go. It's an ongoing project and I'm adding to it regularly.

---

## How it works

A producer polls flight prices on an interval and pushes each event into a Kafka topic. A consumer reads that stream, tracks price history per route, and fires an alert when it detects a meaningful drop.

```
Producer → Kafka (flight-prices) → Consumer → Alerts
```

---

## Stack

- **Go** — learning it through this project, great fit for concurrent producers/consumers
- **Apache Kafka** — handles the message streaming between services
- **franz-go** — Go Kafka client, pure Go with no C dependencies
- **Docker** — Kafka runs locally via Docker Compose

---

## Project layout

```
hermeslens/
├── cmd/
│   ├── producer/main.go      # fetches prices, pushes to Kafka
│   └── consumer/main.go      # reads stream, detects price drops
├── internal/
│   └── models/flight.go      # FlightPriceEvent struct
├── config/config.go           # broker address, topic names
└── docker-compose.yml         # spins up Kafka + Zookeeper
```

---

## Running it locally

You'll need Go 1.22+ and Docker.

```bash
# start Kafka
docker compose up -d

# install dependencies
go mod tidy

# run the producer
go run cmd/producer/main.go
```

To confirm messages are flowing:

```bash
docker exec -it <kafka-container> kafka-console-consumer \
  --bootstrap-server localhost:9092 \
  --topic flight-prices \
  --from-beginning
```

---

## What a price event looks like

```json
{
  "id": "20240301120000.123",
  "origin": "JFK",
  "destination": "LAX",
  "price": 487.32,
  "currency": "EUR",
  "airline": "Delta",
  "departure_date": "2024-04-01",
  "fetched_at": "2024-03-01T12:00:00Z"
}
```

---

## What's next

- [x] Kafka producer with mock price events
- [ ] Consumer with price trend detection
- [ ] Price drop alerts (email / webhook)
- [ ] Hook up a real flight price API
- [ ] Store price history
- [ ] Simple dashboard

---

## Note on config

API keys and production broker URLs go in a `.env` file which is gitignored. Nothing sensitive is committed to this repo.
