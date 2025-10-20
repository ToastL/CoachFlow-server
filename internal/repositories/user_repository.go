package repositories

import (
	"context"
	"coachflow/internal/db"
	"coachflow/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, username, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow(context.Background(), "SELECT id, username, email, password FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}