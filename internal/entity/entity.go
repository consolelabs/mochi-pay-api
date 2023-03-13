package entity

import (
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/db"
	transfer "github.com/consolelabs/mochi-pay-api/internal/entity/transfer"
)

type Entity struct {
	Transfer transfer.ITransfer
}

func New(p *appmain.Params, db *db.Store) *Entity {
	return &Entity{
		Transfer: transfer.New(p, db),
	}
}
