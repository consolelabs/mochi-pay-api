package api

import (
	"strings"

	"github.com/consolelabs/mochi-pay-api/internal/app/api/handler"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter(p *appmain.Params, h *handler.Handler) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.LoggerWithWriter(p.Logger().Writer(), "/healthz"),
		gin.Recovery(),
	)

	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowOrigins := []string{"*"}

		// allow all localhosts and all GET method
		if origin != "" && (strings.Contains(origin, "http://localhost") || c.Request.Method == "GET") {
			allowOrigins = []string{origin}
		} else {
			// suport wildcard cors: https://*.domain.com
			for _, url := range allowOrigins {
				if strings.Contains(origin, strings.Replace(url, "https://*", "", 1)) {
					allowOrigins = []string{origin}
					break
				}
			}
		}

		cors.New(
			cors.Config{
				AllowOrigins: allowOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{"Origin", "Host",
					"Content-Type", "Content-Length",
					"Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
				ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
				AllowCredentials: true,
			},
		)(c)
	})

	// handlers
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// payment group
	r.POST("/api/v1/transfer", h.Transfer.Transfer)
	r.GET("/api/v1/mochi-wallet", nil)
	r.POST("/api/v1/mochi-wallet/deposit", nil)
	r.POST("/api/v1/mochi-wallet/withdraw", nil)

	//
	// r.GET("/api/v1/mochi-wallet/transactions", nil)

	return r
}
