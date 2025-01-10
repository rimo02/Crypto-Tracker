package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rimo02/Crypto-Tracker/src/config"
	"github.com/rimo02/Crypto-Tracker/src/model"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var cryptocoins = []string{"bitcoin", "matic-network", "ethereum"}
var collections = make(map[string]*mongo.Collection)
var db *mongo.Database
func SetCollection(client *mongo.Client) {
	for _, coin := range cryptocoins {
		collections[coin] = client.Database("crypto-db").Collection(coin)
	}
	db = client.Database("crypto-db")
}

func fetchdata(key string) ([]model.CoinsApi, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s",
		"bitcoin,ethereum,matic-network")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-cg-demo-api-key", key)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []model.CoinsApi
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for i := range data {
		data[i].Timestamp = time.Now()
	}
	return data, nil
}

func storedata(ctx context.Context, data []model.CoinsApi) error {
	for _, coin := range data {
		collection, exist := collections[coin.ID]
		if !exist {
			fmt.Printf("Collection for coin %s does not exist", coin.ID)
			continue
		}

		_, err := collection.InsertOne(ctx, coin)
		if err != nil {
			return fmt.Errorf("failed to store data for coin %s: %v", coin.ID, err)
		}
	}
	return nil
}

func FetchCryptoData(c *fiber.Ctx) error {
	ctx := context.Background()
	key := config.GetApiKey()

	if key == "" {
		fmt.Println("W")
		return nil
	}

	data, err := fetchdata(key)
	if err != nil {
		fmt.Printf("Error fetching data: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch data")
	}

	err = storedata(ctx, data)
	if err != nil {
		fmt.Printf("Error storing data: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to store data")
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
