package entity

import "time"

type (
	Status string

	ModerationRequest struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		Text        string    `json:"text"`
		SubmittedAt time.Time `json:"submitted_at"`
		Status      Status    `json:"status"`
	}
)

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)
