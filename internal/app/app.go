package app

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"coachflow/internal/config"
	"coachflow/internal/db"
	"coachflow/internal/routes"
	"coachflow/internal/socket"
)

type App struct {
	router *fiber.App
	config *config.Config
}

func NewApp() *App {
	cfg := config.Load()

	db.Connect(cfg.DatabaseURL)

	f := fiber.New()

	f.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	sio := socket.InitSocket()
	f.Get("/socket/*", func(c *fiber.Ctx) error {
		sio.ServeHTTP(c.Context())
		return nil
	})

	routes.SetupRoutes(f)

	return &App{
		router: f,
		config: cfg,
	}
}

func Run(a *App) error {
	addr := fmt.Sprintf(":%s", a.config.Port)
	return a.router.Listen(addr)
}