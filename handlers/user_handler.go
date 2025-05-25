package handlers

import (
	"github.com/gofiber/fiber/v2"
	"BelajarGolang/models"
)

func GetUser(c *fiber.Ctx) error {
	user := models.User{ID: 1, Name: "Miqdam", Age: 21}
	return c.JSON(user)
}
