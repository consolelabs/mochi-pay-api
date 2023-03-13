package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ActivityLog struct {
	Id             uuid.UUID      `json:"id"`
	ProfileId      string         `json:"profile_id"`
	Receiver       pq.StringArray `json:"receiver" gorm:"type:varchar(256)[];"`
	NumberReceiver int64          `json:"number_receiver"`
	TokenId        string         `json:"token_id"`
	Amount         float64        `json:"amount"`
	Status         string         `json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      time.Time      `json:"deleted_at"`
	Note           string         `json:"note"`
}
