package transfer

import "github.com/gin-gonic/gin"

type ITransfer interface {
	Transfer(c *gin.Context)
}
