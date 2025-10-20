package response

import "github.com/gofiber/fiber/v3"

func JSON(c fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}

func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": true,
		"message": message,
	})
}