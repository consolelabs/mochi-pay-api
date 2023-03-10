package api

import (
	"net/http"

	"github.com/consolelabs/social-payment-api/internal/appmain"
)

// BindService creates the backend service and binds it to the serving harness
func BindService(p *appmain.Params, b *appmain.Bindings) appmain.IServer {
	router := setupRouter()

	if !p.Config().IsSet("PORT") {
		p.Logger().Fatal("PORT environment variable not set")
	}

	srv := &http.Server{
		Addr:    ":" + p.Config().GetString("PORT"),
		Handler: router,
	}
	return srv
}
