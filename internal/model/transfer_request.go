package model

import "time"

type TransferRequest struct {
	From   *Wallet   `json:"from"`
	Tos    []*Wallet `json:"tos"`
	Amount []string  `json:"amount"`
	Token  *Token    `json:"token"`
	Note   string    `json:"note"`
	// Status string

	CreatedAt *time.Time `json:"created_at"`
}
