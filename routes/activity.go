package routes

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"go-to-do/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Use(fiberLogger.New())

	app.Get("/activity-groups", handlers.GetActivities)
	app.Post("/activity-groups", handlers.CreateActivity)
	app.Get("/activity-groups/:id", handlers.GetActivity)
	app.Patch("/activity-groups/:id", handlers.UpdateActivity)
	app.Delete("/activity-groups/:id", handlers.DeleteActivity)

	app.Get("/todo-items", handlers.GetTodos)
	app.Post("/todo-items", handlers.CreateTodo)
	app.Get("/todo-items/:id", handlers.GetTodo)
	app.Patch("/todo-items/:id", handlers.UpdateTodo)
	app.Delete("/todo-items/:id", handlers.DeleteTodo)
}
