package token

import (
	"gorm.io/gorm"

	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) IToken {
	return &store{
		db: db,
	}
}

func (s *store) GetBySymbol(symbol string) (token *model.Token, err error) {
	return token, s.db.Where("symbol ILIKE ?", symbol).First(&token).Error
}
