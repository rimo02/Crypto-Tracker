package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"time"
)

type CryptoPrice struct {
	Coin      string    `bson:"coin"`
	Price     float64   `bson:"currentprice"`
	Timestamp time.Time `bson:"timestamp"`
}

func CalculateDeviation(c *fiber.Ctx) error {
	coin := c.Query("coin")
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

	// fetch the last 100 records
	collection := db.Collection(coin)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalRecords, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error counting documents:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count records in the database",
		})
	}
	limit := int64(100)
	if limit < 100 {
		limit = totalRecords
	}
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}).SetLimit(limit)
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		fmt.Println("Error fetching records:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch records from the database",
		})
	}
	defer cursor.Close(ctx)

	//calculating sdv
	var prices []CryptoPrice
	if err := cursor.All(ctx, &prices); err != nil {
		fmt.Println("Error decoding records:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode records",
		})
	}

	if len(prices) <= 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No data found for the specified coin",
		})
	}

	// calculate mean
	var mean, stdv float64
	for _, p := range prices {
		mean += p.Price
	}
	mean = mean / float64(len(prices))
	for _, p := range prices {
		stdv += math.Pow(p.Price-mean, 2)
	}
	stdv = math.Sqrt(stdv / float64(len(prices)))

	return c.JSON(fiber.Map{
		"coin":          coin,
		"total_records": totalRecords,
		"mean_price":    mean,
		"deviation":     stdv,
	})
}
