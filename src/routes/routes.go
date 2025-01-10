package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rimo02/Crypto-Tracker/src/controllers"
)

var SetSearchRoutes = func(app *fiber.App) {
	app.Get("/stats", controllers.LatestCryptoData)       
	app.Get("/deviation", controllers.CalculateDeviation)
}
