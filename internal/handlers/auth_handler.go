package handlers

import (
	"coachflow/internal/db"
	"coachflow/pkg/response"
	"coachflow/pkg/utils"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func Login(c fiber.Ctx) error {
	type LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var body LoginRequest
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	var id int64
	var hashed, email string
	err := db.DB.QueryRow(c.Context(), `
		SELECT id, email, password FROM users WHERE email=$1 OR username=$1
	`, body.Login).Scan(&id, &email, &hashed)
	if err != nil {
		return response.JSON(c, fiber.StatusUnauthorized, fiber.Map{
			"error": fiber.Map{"login": "Invalid username or email"},
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(body.Password)) != nil {
		return response.JSON(c, fiber.StatusUnauthorized, fiber.Map{
			"error": fiber.Map{"password": "Invalid password"},
		})
	}

	token, err := utils.GenerateJWT(1, body.Login)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not generate token")
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{
		"token": token,
	})
}

func Register(c fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body RegisterRequest
	if err := c.Bind().Body(&body); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	var emailExists, usernameExists bool

	err := db.DB.QueryRow(c.Context(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", body.Email).Scan(&emailExists)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Database error checking email")
	}

	err = db.DB.QueryRow(c.Context(), "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", body.Username).Scan(&usernameExists)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Database error checking username")
	}

	errors := fiber.Map{}
	if emailExists {
		errors["email"] = "Email is already in use"
	}
	if usernameExists {
		errors["username"] = "Username is already in use"
	}

	if len(errors) > 0 {
		return response.JSON(c, fiber.StatusConflict, fiber.Map{
			"error": errors,
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not hash password")
	}

	_, err = db.DB.Exec(c.Context(),
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		body.Username, body.Email, string(hashed))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not create user")
	}

	token, err := utils.GenerateJWT(1, body.Email)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not generate token")
	}

	return response.JSON(c, fiber.StatusCreated, fiber.Map{
		"token": token,
	})
}