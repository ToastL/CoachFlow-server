package handlers

import (
	"coachflow/internal/models"
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func CreateRequest(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var body models.Request
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}
	if body.ToID == 0 || body.Type == "" {
		return response.Error(c, fiber.StatusBadRequest, "Missing fields")
	}

	body.FromID = int64(userID)

	err := repositories.CreateRequest(body)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create request")
	}

	return response.JSON(c, fiber.StatusCreated, fiber.Map{"message": "Request sent"})
}

func GetRequests(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	requests, err := repositories.GetRequests(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch requests")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"requests": requests})
}

func GetSentRequests(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	requests, err := repositories.GetSentRequests(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch requests")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"requests": requests})
}

func AcceptRequest(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request ID")
	}

	err = repositories.UpdateRequestStatus(id, true)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to accept request")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"message": "Request accepted"})
}

func RejectRequest(c fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request ID")
	}

	err = repositories.UpdateRequestStatus(id, false)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to reject request")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"message": "Request rejected"})
}