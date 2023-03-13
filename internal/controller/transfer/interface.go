package transfer

import (
	"github.com/consolelabs/mochi-pay-api/internal/apperror/apierror"
	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type ITransfer interface {
	TransferToken(req *model.TransferRequest) *apierror.ApiError
}
