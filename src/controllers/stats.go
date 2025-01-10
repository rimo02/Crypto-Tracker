package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rimo02/Crypto-Tracker/src/config"
	"github.com/rimo02/Crypto-Tracker/src/model"
	"net/http"
	"time"
)

type CryptoData struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	PriceUSD float64 `json:"price_usd"`
}

func LatestCryptoData(c *fiber.Ctx) error {
	coin := c.Query("coin")
	key := config.GetApiKey()
	if coin == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing query parameter: coin",
		})
	}

	// Validate the coin name
	validCoins := map[string]bool{"bitcoin": true, "ethereum": true, "matic-network": true}
	if !validCoins[coin] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid coin name: %s. Allowed coins: bitcoin, ethereum, matic-network", coin),
		})
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s",
		coin)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-cg-demo-api-key", key)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch cryptocurrency data",
		})
	}
	defer resp.Body.Close()

	var data []model.CoinsApi
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse cryptocurrency data",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Coin":         coin,
		"Symbol":       data[0].Symbol,
		"Price":        data[0].CurrentPrice,
		"MarketCap":    data[0].MarketCap,
		"24HourChange": data[0].PriceChangePercentage24h,
	})
}
