package middleware

import (
	"github.com/bigchange/go-practice/testcase/hub"
	"github.com/gin-gonic/gin"
)

var (
	hubCtxKey = "hub"
)

func SetHub(h *hub.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(hubCtxKey, h)
		c.Next()
	}
}

func GetHubFromCtx(c *gin.Context) *hub.Hub {
	return c.MustGet(hubCtxKey).(*hub.Hub)
}
