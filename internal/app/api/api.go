package api

import (
	"net/http"

	"github.com/consolelabs/mochi-pay-api/internal/app/api/handler"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
	"github.com/consolelabs/mochi-pay-api/internal/db"
	"github.com/consolelabs/mochi-pay-api/internal/entity"
)

// BindService creates the backend service and binds it to the serving harness
func BindService(p *appmain.Params, b *appmain.Bindings) appmain.IServer {
	db := db.New(p)
	entity := entity.New(p, db)
	handler := handler.New(p, entity)
	router := setupRouter(p, handler)

	if !p.Config().IsSet("PORT") {
		p.Logger().Fatal("PORT environment variable not set")
	}

	srv := &http.Server{
		Addr:    ":" + p.Config().GetString("PORT"),
		Handler: router,
	}
	return srv
}
