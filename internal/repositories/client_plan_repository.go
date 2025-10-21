package repositories

import (
	"context"
	"coachflow/internal/db"
	"coachflow/internal/models"
)

func AssignPlan(cp models.ClientPlan) error {
	_, err := db.DB.Exec(context.Background(),
		`INSERT INTO client_plans (trainer_id, client_id, plan_id, status)
		 VALUES ($1, $2, $3, 'active')`,
		cp.TrainerID, cp.ClientID, cp.PlanID)
	return err
}

func GetClientPlans(clientID int64) ([]models.ClientPlan, error) {
	rows, err := db.DB.Query(context.Background(),
		`SELECT id, trainer_id, client_id, plan_id, status, assigned_at
		 FROM client_plans
		 WHERE client_id=$1
		 ORDER BY assigned_at DESC`, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.ClientPlan
	for rows.Next() {
		var p models.ClientPlan
		if err := rows.Scan(&p.ID, &p.TrainerID, &p.ClientID, &p.PlanID, &p.Status, &p.AssignedAt); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, nil
}

func GetTrainerPlans(trainerID int64) ([]models.ClientPlan, error) {
	rows, err := db.DB.Query(context.Background(),
		`SELECT id, trainer_id, client_id, plan_id, status, assigned_at
		 FROM client_plans
		 WHERE trainer_id=$1
		 ORDER BY assigned_at DESC`, trainerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.ClientPlan
	for rows.Next() {
		var p models.ClientPlan
		if err := rows.Scan(&p.ID, &p.TrainerID, &p.ClientID, &p.PlanID, &p.Status, &p.AssignedAt); err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, nil
}