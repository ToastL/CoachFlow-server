package models

import "time"

type ClientPlan struct {
	ID         int64     `json:"id"`
	TrainerID  int64     `json:"trainer_id"`
	ClientID   int64     `json:"client_id"`
	PlanID     int64     `json:"plan_id"`
	Status     string    `json:"status"`
	AssignedAt time.Time `json:"assigned_at"`
}