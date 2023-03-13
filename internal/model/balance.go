package model

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	Id            uuid.UUID `json:"id"`
	ProfileId     string    `json:"profile_id"`
	TokenId       string    `json:"token_id"`
	Amount        float64   `json:"amount"`
	ChangedAmount float64   `json:"-" gorm:"-"` // This is not a column in the database
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
