package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rimo02/Crypto-Tracker/db"
	"github.com/rimo02/Crypto-Tracker/src/config"
	"github.com/rimo02/Crypto-Tracker/src/controllers"
	"github.com/rimo02/Crypto-Tracker/src/routes"
	"github.com/valyala/fasthttp"
	"time"
)

func init() {
	db.InitMongoDB()
	config.InitConfig()
}

func main() {
	// start a goroutine to fetch videos from Coingecko periodically every 2 houurs and update them on the database
	app := fiber.New()
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
				err := controllers.FetchCryptoData(ctx)
				if err != nil {
					log.Errorf("Error fetching videos and updating database %v", err)
					return
				}
			}
		}
	}()
	routes.SetSearchRoutes(app)
	app.Listen(":3000")
}
