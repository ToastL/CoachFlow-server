package models

import "time"

type Request struct {
	ID        int64     `json:"id"`
	FromID    int64     `json:"from_id"`
	ToID      int64     `json:"to_id"`
	Type      string    `json:"type"`     // "trainer" or "client"
	Status    string    `json:"status"`   // "pending", "accepted", "rejected"
	CreatedAt time.Time `json:"created_at"`
}