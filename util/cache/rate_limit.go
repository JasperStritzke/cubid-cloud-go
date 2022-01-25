package cache

import (
	"github.com/gin-gonic/gin"
)

func RateLimit(ctx *gin.Context) {
	//TODO add rate limiting

	ctx.Next()
}
