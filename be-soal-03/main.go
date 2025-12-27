package main

import (
	"log"

	"be-soal-03/config"
	"be-soal-03/database"
	"be-soal-03/routes"

	_ "be-soal-03/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title Ticket Online API
// @version 1.0
// @description Backend API for Online Event Ticket Platform
// @termsOfService http://example.com/terms/

// @contact.name Backend Developer - Leonardo Gunawan
// @contact.email leonardogunawan15@email.com

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	config.ENVLoad()
	database.Connect()
	database.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	
	app.Get("/swagger/*", swagger.HandlerDefault)
	
	routes.RoutesList(app)
	
	log.Fatal(app.Listen(":8080"))
}
