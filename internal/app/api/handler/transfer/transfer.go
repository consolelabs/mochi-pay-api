package transfer

import "github.com/gin-gonic/gin"

type handler struct {
}

func New() ITransfer {
	return &handler{}
}

func (h *handler) Transfer(c *gin.Context) {
}
