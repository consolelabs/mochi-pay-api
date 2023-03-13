package balance

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

func (p *pg) UpsertBatch(list []model.Balance) error {
	tx := p.db.Begin()
	for i, item := range list {
		err := tx.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "token_id"}, {Name: "profile_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"amount": gorm.Expr("balance.amount + ?", item.ChangedAmount)}),
			},
			clause.Returning{Columns: []clause.Column{{Name: "amount"}}},
		).Create(&item).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		list[i] = item
	}
	return tx.Commit().Error
}
