package token

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

func (p *pg) GetBySymbol(symbol string) (token *model.Token, err error) {
	return token, p.db.Where("symbol ILIKE ?", symbol).First(&token).Error
}
