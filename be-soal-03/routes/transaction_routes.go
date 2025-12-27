package routes

import (
	"be-soal-03/controllers"
	"be-soal-03/middlewares"

	"github.com/gofiber/fiber/v2"
)

func transactionRoutes(api fiber.Router) {
	transactions := api.Group("/transactions")

	transactions.Post("/",
		middlewares.AuthMiddleware,
		controllers.CreateTransaction,
	)

	transactions.Get("/",
		middlewares.AuthMiddleware,
		controllers.GetMyTransactions,
	)
}
