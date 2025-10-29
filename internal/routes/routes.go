package routes

import (
	"github.com/gofiber/fiber/v3"
	"coachflow/internal/handlers"
	"coachflow/internal/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

	// Users
	users := api.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Get("/me", middlewares.AuthMiddleware, handlers.GetCurrentUser)
	users.Get("/:id", handlers.GetUserByID)

	requests := api.Group("/requests", middlewares.AuthMiddleware)
	requests.Post("/", handlers.CreateRequest)
	requests.Get("/", handlers.GetRequests)
	requests.Get("/sent", handlers.GetSentRequests)
	requests.Put("/:id/accept", handlers.AcceptRequest)
	requests.Put("/:id/reject", handlers.RejectRequest)

	clientPlan := api.Group("/clients", middlewares.AuthMiddleware)
	clientPlan.Get("/", handlers.GetClients)
	clientPlan.Get("/:clientId/plans", handlers.GetClientPlans)
	clientPlan.Post("/:clientId/plans", handlers.AssignPlan)

	// Workouts
	workouts := api.Group("/workouts")
	workouts.Get("/", handlers.GetAllWorkouts)
	workouts.Post("/", handlers.CreateWorkout)

	// Plans
	plans := api.Group("/plans")
	plans.Get("/", handlers.GetPlans)
	plans.Post("/", handlers.CreatePlan)
}