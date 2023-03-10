package main

import (
	"github.com/consolelabs/mochi-pay-api/internal/app/api"
	"github.com/consolelabs/mochi-pay-api/internal/appmain"
)

func main() {
	appmain.RunApplication("api", api.BindService)
}
