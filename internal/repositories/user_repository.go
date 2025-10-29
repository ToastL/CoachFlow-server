package repositories

import (
	"context"
	"coachflow/internal/db"
	"coachflow/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow(context.Background(), "SELECT id, username, email, role FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}