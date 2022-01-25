package middlewares

import (
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireLoggedIn(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) == 0 || !auth.ExistsToken(authHeader) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}

func RequirePermission(nodes ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) == 0 || !auth.HasPermissions(authHeader, nodes...) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
