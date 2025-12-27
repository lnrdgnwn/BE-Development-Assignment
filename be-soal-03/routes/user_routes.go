package routes

import (
	"be-soal-03/controllers"
	"be-soal-03/middlewares"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(api fiber.Router) {
	users := api.Group("/users")

	// authenticated user
	users.Get("/me",
		middlewares.AuthMiddleware,
		controllers.GetMyProfile,
	)

	// admin only
	users.Get("/",
		middlewares.AuthMiddleware,
		middlewares.AdminMiddleware,
		controllers.GetUsers,
	)
}
