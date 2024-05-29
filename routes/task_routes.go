// routes/task_routes.go
package routes

import (
	"json-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/tasks", handlers.GetTasks)
	app.Get("/tasks/:id", handlers.GetTask)
	app.Post("/tasks", handlers.CreateTask)
	app.Put("/tasks/:id", handlers.UpdateTask)
	app.Delete("/tasks/:id", handlers.DeleteTask)
}
