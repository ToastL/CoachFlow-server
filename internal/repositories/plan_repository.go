package repositories

import (
	"context"
	"coachflow/internal/db"
	"coachflow/internal/models"
)

func CreatePlan(p models.Plan) error {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO plans (user_id, title, description) VALUES ($1, $2, $3)",
		p.UserID, p.Title, p.Description)
	return err
}

func GetAllPlans() ([]models.Plan, error) {
	rows, err := db.DB.Query(context.Background(),
		"SELECT id, user_id, title, description, created_at FROM plans ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plans := []models.Plan{}
	for rows.Next() {
		var p models.Plan
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.CreatedAt); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, nil
}