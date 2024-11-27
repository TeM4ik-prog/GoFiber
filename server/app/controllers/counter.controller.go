package controllers

import (
	"fiber-app/app/services"

	"github.com/gofiber/fiber/v2"
)

type CounterController struct {
	CounterService *services.CounterService
}

func NewCounterController(service *services.CounterService) *CounterController {
	return &CounterController{CounterService: service}
}

func (cc *CounterController) GetCounter(c *fiber.Ctx) error {
	count := cc.CounterService.GetCount()
	return c.JSON(fiber.Map{"count": count})
}

func (cc *CounterController) IncrementCounter(c *fiber.Ctx) error {
	count := cc.CounterService.Increment()
	return c.JSON(fiber.Map{"count": count})
}

func (cc *CounterController) ResetCounter(c *fiber.Ctx) error {
	count := cc.CounterService.Reset()
	return c.JSON(fiber.Map{"count": count})
}
