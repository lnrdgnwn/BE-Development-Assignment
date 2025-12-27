package routes

import (
	"be-soal-03/controllers"
	"be-soal-03/middlewares"

	"github.com/gofiber/fiber/v2"
)


func eventRoutes(api fiber.Router) {
	events := api.Group("/events")

	// public
	events.Get("/", controllers.GetEvents)
	events.Get("/:id", controllers.GetEventByID)

	// Admin only
	events.Post("/",
		middlewares.AuthMiddleware,
		middlewares.AdminMiddleware,
		controllers.CreateEvent,
	)

	events.Put("/:id",
		middlewares.AuthMiddleware,
		middlewares.AdminMiddleware,
		controllers.UpdateEvent,
	)
}
