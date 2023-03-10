package main

import (
	"github.com/consolelabs/social-payment-api/internal/app/api"
	"github.com/consolelabs/social-payment-api/internal/appmain"
)

func main() {
	appmain.RunApplication("api", api.BindService)
}
