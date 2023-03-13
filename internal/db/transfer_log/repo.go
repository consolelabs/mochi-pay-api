package transferlog

import (
	"gorm.io/gorm"

	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type pg struct {
	db *gorm.DB
}

func New(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (p *pg) CreateTransferLog(al *model.TransferLog) error {
	return p.db.Create(al).Error
}
