package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-to-do/database"
	"go-to-do/pkg"
	"go-to-do/routes"
)

func main() {
	pkg.InitDatabase()
	database.RunMigration()

	app := fiber.New()
	routes.SetupRoutes(app)

	err := app.Listen(":3030")
	fmt.Println("Server started at port 3030")
	if err != nil {
		panic("Failed to start server")
	}
}
