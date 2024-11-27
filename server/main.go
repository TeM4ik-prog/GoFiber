package main

import (
	"fiber-app/app/routes"
	"fiber-app/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	api := app.Group("/api")

	routes.RegisterCounterRoutes(api, "/counter")
	routes.RegisterUserRoutes(api, "/users")

	app.Static("/", "../client/dist")
	app.Listen(":3000")
}
