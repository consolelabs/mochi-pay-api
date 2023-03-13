package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type TransferLog struct {
	Id                  uuid.UUID      `json:"id"`
	SenderProfileId     string         `json:"sender_profile_id"`
	RecipientsProfileId pq.StringArray `json:"recipients_profile_id" gorm:"type:varchar(256)[];"`
	NumberReceiver      int64          `json:"number_receiver"`
	TokenId             string         `json:"token_id"`
	Amount              float64        `json:"amount"`
	Status              string         `json:"status"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           time.Time      `json:"deleted_at"`
	Note                string         `json:"note"`
}
