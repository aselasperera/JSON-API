// main.go
package main

import (
	"json-api/database"
	"json-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := database.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
