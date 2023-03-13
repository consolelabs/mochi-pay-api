package handler

import (
	"github.com/consolelabs/mochi-pay-api/internal/app/api/handler/transfer"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/entity"
)

type Handler struct {
	Transfer transfer.ITransfer
}

func New(p *appmain.Params, entity *entity.Entity) *Handler {
	return &Handler{
		Transfer: transfer.New(p, entity),
	}
}
