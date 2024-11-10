package main

import (
	"log"
	"producer/config"
	"producer/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	kafkaProducer := config.NewKafkaProducer()
	defer kafkaProducer.Close() // Close gracefully

	api := app.Group("/api/v1") // Group API routes
	routes.RegisterCommentRoutes(api, kafkaProducer) // Register routes

	log.Fatal(app.Listen(":3000")) // Start server
}