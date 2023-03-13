package handler

import (
	"github.com/consolelabs/mochi-pay-api/internal/app/api/handler/transfer"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/controller"
)

type Handler struct {
	Transfer transfer.ITransfer
}

func New(p *appmain.Params) *Handler {
	ctrl := controller.New(p)
	return &Handler{
		Transfer: transfer.New(ctrl),
	}
}
