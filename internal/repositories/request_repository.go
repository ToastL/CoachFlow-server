package repositories

import (
	"coachflow/internal/db"
	"coachflow/internal/models"
	"context"
)

func CreateRequest(r models.Request) error {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO requests (from_id, to_id, type, status) VALUES ($1, $2, $3, 'pending') ON CONFLICT DO NOTHING",
		r.FromID, r.ToID, r.Type)
	return err
}

func GetRequests(userID uint) ([]models.Request, error) {
	rows, err := db.DB.Query(context.Background(),
		`SELECT id, from_id, to_id, type, status, created_at
		 FROM requests
		 WHERE to_id=$1
		 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := []models.Request{}
	for rows.Next() {
		var r models.Request
		if err := rows.Scan(&r.ID, &r.FromID, &r.ToID, &r.Type, &r.Status, &r.CreatedAt); err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}
	return requests, nil
}

func GetSentRequests(userID uint) ([]models.Request, error) {
	rows, err := db.DB.Query(context.Background(),
		`SELECT id, from_id, to_id, type, status, created_at
		 FROM requests
		 WHERE from_id=$1
		 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := []models.Request{}
	for rows.Next() {
		var r models.Request
		if err := rows.Scan(&r.ID, &r.FromID, &r.ToID, &r.Type, &r.Status, &r.CreatedAt); err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}
	return requests, nil
}

func UpdateRequestStatus(id int64, accept bool) error {
	status := "rejected"
	if accept {
		status = "accepted"
	}

	_, err := db.DB.Exec(context.Background(),
		"UPDATE requests SET status=$1 WHERE id=$2", status, id)
	
	if err != nil {
		return err
	}

	if accept {
		_, err = db.DB.Exec(context.Background(),
			`INSERT INTO trainer_clients (trainer_id, client_id)
			 SELECT CASE WHEN type='trainer' THEN to_id ELSE from_id END AS trainer_id,
			 		CASE WHEN type='trainer' THEN from_id ELSE to_id END AS client_id
			 FROM requests WHERE id=$1`, id)
	}
	return err
}