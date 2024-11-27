package controllers

import (
	"fiber-app/app/services"
	"fiber-app/database/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUsersController(service *services.UserService) *UserController {
	return &UserController{UserService: service}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {

	return c.JSON(uc.UserService.GetUsers())

}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User

	println(c.BodyParser(user))

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request data"})
	}

	createdUser := uc.UserService.CreateUser(user)

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}
