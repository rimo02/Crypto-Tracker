package model

import "time"

type CoinsApi struct {
	ID                       string    `json:"id"`
	Symbol                   string    `json:"symbol"`
	Name                     string    `json:"name"`
	CurrentPrice             float64   `json:"current_price"`
	MarketCap                float64   `json:"market_cap"`
	PriceChangePercentage24h float64   `json:"price_change_percentage_24h"`
	Timestamp                time.Time `json:"timestamp"`
}
