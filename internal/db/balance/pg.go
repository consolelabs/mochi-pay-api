package balance

import "github.com/consolelabs/mochi-pay-api/internal/model"

type Store interface {
	GetBalanceByTokenID(profileId string, tokenId string) (balance *model.Balance, err error)
	UpsertBatch(list []model.Balance) error
}
