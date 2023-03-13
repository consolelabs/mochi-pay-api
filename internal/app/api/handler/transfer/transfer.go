package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
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
		From:   req.From,
		Tos:    req.Tos,
		Amount: req.Amount,
		Token:  req.Token,
		Note:   req.Note,
	})
	if err != nil {
		switch err.Error() {
		case "token not supported":
			pp.Println("in this case 1")
			c.JSON(http.StatusBadRequest, apierror.New("token not supported", 400, apierror.Code400))
			return
		case "insufficient balance":
			c.JSON(http.StatusBadRequest, apierror.New("insufficient balance", 400, apierror.Code400))
			return
		default:
			logger.Error(err, "[handler.Transfer] - failed to transfer token on entities level")
			c.JSON(http.StatusInternalServerError, apierror.New("Something went wrong", 500, apierror.Code500))
			return
		}
	}

	c.JSON(http.StatusOK, view.ToRespSuccess())
}
