package handlers

import (
	"coachflow/internal/models"
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func AssignPlan(c fiber.Ctx) error {
	trainerID := c.Locals("userID").(int64)

	clientID, err := strconv.ParseInt(c.Params("clientId"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid client ID")
	}

	var body struct {
		PlanID int64 `json:"plan_id"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid body")
	}
	if body.PlanID == 0 {
		return response.Error(c, fiber.StatusBadRequest, "Missing plan_id")
	}

	cp := models.ClientPlan{
		TrainerID: trainerID,
		ClientID:  clientID,
		PlanID:    body.PlanID,
	}

	err = repositories.AssignPlan(cp)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to assign plan")
	}

	return response.JSON(c, fiber.StatusCreated, fiber.Map{"message": "Plan assigned successfully"})
}

func GetClientPlans(c fiber.Ctx) error {
	clientID, err := strconv.ParseInt(c.Params("clientId"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid client ID")
	}

	plans, err := repositories.GetClientPlans(clientID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch plans")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"plans": plans})
}

func GetTrainerPlans(c fiber.Ctx) error {
	trainerID := c.Locals("userID").(int64)

	plans, err := repositories.GetTrainerPlans(trainerID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch trainer plans")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"plans": plans})
}