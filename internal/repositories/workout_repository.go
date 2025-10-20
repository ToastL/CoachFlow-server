package repositories

import (
	"context"
	"coachflow/internal/db"
	"coachflow/internal/models"
)

func CreateWorkout(w models.Workout) error {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO workouts (user_id, title, description) VALUES ($1, $2, $3)",
		w.UserID, w.Title, w.Description)
	return err
}

func GetAllWorkouts() ([]models.Workout, error) {
	rows, err := db.DB.Query(context.Background(),
		"SELECT id, user_id, title, description, created_at FROM workouts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := []models.Workout{}
	for rows.Next() {
		var w models.Workout
		if err := rows.Scan(&w.ID, &w.UserID, &w.Title, &w.Description, &w.CreatedAt); err != nil {
			return nil, err
		}
		workouts = append(workouts, w)
	}
	return workouts, nil
}