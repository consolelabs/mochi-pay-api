package controller

import (
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	transfer "github.com/consolelabs/mochi-pay-api/internal/controller/transfer"
)

type Controller struct {
	Transfer transfer.ITransfer
}

func New(p *appmain.Params) *Controller {
	return &Controller{
		Transfer: transfer.New(p),
	}
}
