package handler

import (
	"github.com/consolelabs/mochi-pay-api/internal/app/api/handler/transfer"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
)

type Handler struct {
	Transfer transfer.ITransfer
}

func New(p *appmain.Params) *Handler {
	return &Handler{
		Transfer: transfer.New(),
	}
}
