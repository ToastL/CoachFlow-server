package middlewares

import (
	"strings"

	"coachflow/pkg/response"
	"coachflow/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return response.Error(c, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid token format. Use: Bearer <token>")
	}

	claims, err := utils.ValidateJWT(tokenStr)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid or expired token")
	}

	c.Locals("userID", claims.UserID)
	c.Locals("email", claims.Email)
	
	return c.Next()
}