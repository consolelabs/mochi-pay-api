package balance

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

func (p *pg) GetBalanceByTokenID(profileId string, tokenId string) (balance *model.Balance, err error) {
	return balance, p.db.First(&balance, "profile_id = ? AND token_id = ?", profileId, tokenId).Error
}
