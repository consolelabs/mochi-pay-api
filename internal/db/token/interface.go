package token

import "github.com/consolelabs/mochi-pay-api/internal/model"

type IToken interface {
	GetBySymbol(symbol string) (token *model.Token, err error)
}
