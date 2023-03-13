package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/entity"
	"github.com/consolelabs/mochi-pay-api/internal/model"
)

type handler struct {
	Params *appmain.Params
	Entity *entity.Entity
}

func New(p *appmain.Params, e *entity.Entity) ITransfer {
	return &handler{
		Params: p,
		Entity: e,
	}
}

func (h *handler) Transfer(c *gin.Context) {
	req := TransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Params.Logger().WithFields(logrus.Fields{"req": req}).Error(err, "[handler.TransferToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Entity.Transfer.TransferToken(&model.TransferRequest{
		From:   req.From,
		Tos:    req.Tos,
		Amount: req.Amount,
		Token:  req.Token,
		Note:   req.Note,
	})
	if err != nil {
		switch err.Error() {
		case "token not supported":
			c.JSON(http.StatusBadRequest, gin.H{"error": "token not support"})
			return
		case "insufficient balance":
			c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
			return
		default:
			h.Params.Logger().Error(err, "[handler.Transfer] - failed to transfer token on entities level")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
