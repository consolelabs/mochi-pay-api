package nftsaleentity

import "github.com/consolelabs/mochi-pay-api/internal/model"

type ITransfer interface {
	TransferToken(req *model.TransferRequest) error
}
