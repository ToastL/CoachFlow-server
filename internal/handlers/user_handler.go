package handlers

import (
	"github.com/gofiber/fiber/v3"
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
)

func GetUsers(c fiber.Ctx) error {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch users")
	}
	return response.JSON(c, fiber.StatusOK, users)
}

func GetUserByID(c fiber.Ctx) error {
	id := c.Params("id")
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch user")
	}
	
	if user == nil {
		return response.Error(c, fiber.StatusNotFound, "User not found")
	}
	return response.JSON(c, fiber.StatusOK, user)
}