package handlers

import (
	"coachflow/internal/models"
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
	"github.com/gofiber/fiber/v3"
)

func CreatePlan(c fiber.Ctx) error {
	var body models.Plan
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	if body.Title == "" {
		return response.Error(c, fiber.StatusBadRequest, "Title is required")
	}

	err := repositories.CreatePlan(body)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not create plan")
	}

	return response.JSON(c, fiber.StatusCreated, fiber.Map{"message": "Plan created"})
}

func GetPlans(c fiber.Ctx) error {
	plans, err := repositories.GetAllPlans()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch plans")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"plans": plans})
}