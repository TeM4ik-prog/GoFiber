package routes

import (
	"fiber-app/app/controllers"
	"fiber-app/app/services"
	"github.com/gofiber/fiber/v2"
)
	

func RegisterUserRoutes(router fiber.Router, routePath string) {
	userService := services.NewUserService()
	userController := controllers.NewUsersController(userService)

	user := router.Group(routePath)

	user.Get("/", userController.GetUsers)
	user.Post("/create", userController.CreateUser)
	// user.Post("/reset", userController.ResetCounter)
}
