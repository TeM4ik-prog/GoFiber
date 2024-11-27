package routes

import (
	"fiber-app/app/controllers"
	"fiber-app/app/services"
	"github.com/gofiber/fiber/v2"
)
	

func RegisterCounterRoutes(router fiber.Router, routePath string) {
	counterService := services.NewCounterService()
	counterController := controllers.NewCounterController(counterService)

	counter := router.Group(routePath)

	counter.Get("/", counterController.GetCounter)
	counter.Post("/increment", counterController.IncrementCounter)
	counter.Post("/reset", counterController.ResetCounter)
}
