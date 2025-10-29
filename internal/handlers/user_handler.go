package handlers

import (
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch users")
	}
	return response.JSON(c, fiber.StatusOK, users)
}

func GetUserByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}
	id := uint(idUint64)

	user, err := repositories.GetUserByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch user")
	}
	
	if user == nil {
		return response.Error(c, fiber.StatusNotFound, "User not found")
	}
	return response.JSON(c, fiber.StatusOK, user)
}

func GetCurrentUser(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch user")
	}

	if user == nil {
		return response.Error(c, fiber.StatusNotFound, "User not found")
	}
	return response.JSON(c, fiber.StatusOK, user)
}