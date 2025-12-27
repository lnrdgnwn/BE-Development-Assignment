package routes

import (
	"be-soal-03/controllers"

	"github.com/gofiber/fiber/v2"
)

func authRoutes(api fiber.Router) {
	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}