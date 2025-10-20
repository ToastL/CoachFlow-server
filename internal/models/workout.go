package models

import "time"

type Workout struct {
	ID			int64		`json:"id"`
	UserID		int64		`json:"user_id"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	CreatedAt	time.Time	`json:"created_at"`
}