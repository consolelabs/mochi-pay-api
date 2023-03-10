package transfer

import (
	"github.com/consolelabs/mochi-pay-api/internal/apperror/apierror"
	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type TransferRequest struct {
	From   *model.Wallet   `json:"from"`
	Tos    []*model.Wallet `json:"tos"`
	Amount []string        `json:"amount"`
	Token  *model.Token    `json:"token"`
	Note   string          `json:"note"`
}

func (t *TransferRequest) Validate() error {
	if t.From == nil {
		return apierror.ErrFromWalletRequired
	}
	if len(t.Tos) == 0 {
		return apierror.ErrAmountMismatch
	}
	if len(t.Amount) == 0 {
		return apierror.ErrAmountMismatch
	}

	return nil
}
