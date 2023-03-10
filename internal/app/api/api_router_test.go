package api

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"

	"github.com/consolelabs/mochi-pay-api/internal/appmain"
)

// Test_loadV1Routes simply test we load route and handler correctly
func Test_setupRouter(t *testing.T) {
	p := &appmain.Params{}

	expectedRoutes := map[string]map[string]gin.RouteInfo{
		"/healthz": {
			"GET": {
				Method:  "GET",
				Handler: "github.com/consolelabs/mochi-pay-api/internal/app/api.setupRouter.func2",
			},
		},
	}

	router := setupRouter(p, nil)

	routeInfo := router.Routes()

	for _, r := range routeInfo {
		require.NotNil(t, r.HandlerFunc)
		expected, ok := expectedRoutes[r.Path][r.Method]
		require.True(t, ok, fmt.Sprintf("unexpected path: %s", r.Path))
		ignoreFields := cmpopts.IgnoreFields(gin.RouteInfo{}, "HandlerFunc", "Path")
		if !cmp.Equal(expected, r, ignoreFields) {
			t.Errorf("route mismatched. \n route path: %v \n diff: %+v", r.Path,
				cmp.Diff(expected, r, ignoreFields))
			t.FailNow()
		}
	}
}
