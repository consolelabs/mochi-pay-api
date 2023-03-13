package token

import (
	"github.com/google/uuid"
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

func (p *pg) GetById(id string) (token *model.Token, err error) {
	uuidTokenId, _ := uuid.Parse(id)
	return token, p.db.Where("id = ?", uuidTokenId).First(&token).Error
}
