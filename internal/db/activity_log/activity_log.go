package activitylog

import (
	"gorm.io/gorm"

	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IActivityLog {
	return &store{
		db: db,
	}
}

func (s *store) CreateActivityLog(al *model.ActivityLog) error {
	return s.db.Create(al).Error
}
