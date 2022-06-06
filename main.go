package main

import (
	"github.com/bruceherve/infura-task/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api", handlers.Greeting)
	app.Get("/api/latest-block", handlers.GetBlockByNumber)
	app.Get("/api/get-transaction", handlers.GetTransactionByBlockHashAndIndex)
	app.Get("/api/get-balance", handlers.GetBalance)
	app.Get("/api/get-transaction-count", handlers.GetBlockTransactionCountByNumber)
	app.Get("/api/get-gas-price", handlers.GetGasPrice)

}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":4000") 

}
