package models

type FlightPriceEvent struct {
	ID            string `json:"id"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	Price         float64	`json:"price"`
	Currency      string `json:"currency"`
	Airline       string `json:"airline"`
	DepartureDate string `json:"departure_date"`
	FetchedAt     string `json:"fetched_at"`
}
