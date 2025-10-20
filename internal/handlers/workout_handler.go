package handlers

import (
	"coachflow/internal/models"
	"coachflow/internal/repositories"
	"coachflow/pkg/response"
	"github.com/gofiber/fiber/v3"
)

func CreateWorkout(c fiber.Ctx) error {
	var body models.Workout
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	if body.Title == "" {
		return response.Error(c, fiber.StatusBadRequest, "Title is required")
	}

	err := repositories.CreateWorkout(body)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not create workout")
	}

	return response.JSON(c, fiber.StatusCreated, fiber.Map{"message": "Workout created"})
}

func GetAllWorkouts(c fiber.Ctx) error {
	workouts, err := repositories.GetAllWorkouts()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch workouts")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{"workouts": workouts})
}