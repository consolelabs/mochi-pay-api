package token

import "github.com/consolelabs/mochi-pay-api/internal/model"

type Store interface {
	GetById(id string) (token *model.Token, err error)
}
