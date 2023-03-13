package activitylog

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

func (p *pg) CreateActivityLog(al *model.ActivityLog) error {
	return p.db.Create(al).Error
}
