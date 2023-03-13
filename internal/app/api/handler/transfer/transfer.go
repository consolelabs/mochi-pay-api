package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/consolelabs/mochi-pay-api/internal/apperror/apierror"
	"github.com/consolelabs/mochi-pay-api/internal/controller"
	"github.com/consolelabs/mochi-pay-api/internal/model"
	"github.com/consolelabs/mochi-pay-api/internal/view"
)

type handler struct {
	controller *controller.Controller
}

func New(controller *controller.Controller) ITransfer {
	return &handler{
		controller: controller,
	}
}

var (
	logger = logrus.WithFields(logrus.Fields{
		"component": "handler.transfer",
	})
)

func (h *handler) Transfer(c *gin.Context) {
	logger.Debug("api call ", c.Request.RequestURI)
	defer logger.Debug("api finish ", c.Request.RequestURI)

	req := TransferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err, "[handler.TransferToken] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.controller.Transfer.TransferToken(&model.TransferRequest{
		From:    req.From,
		Tos:     req.Tos,
		Amount:  req.Amount,
		TokenId: req.TokenId,
		Note:    req.Note,
	})
	if err != nil {
		logger.Error(err, "[handler.Transfer] - failed to transfer token on controller level")
		c.JSON(err.StatusCode(), apierror.New(err.Message(), err.StatusCode(), apierror.APICode(int64(err.StatusCode()))))
		return
	}

	c.JSON(http.StatusOK, view.ToRespSuccess())
}
