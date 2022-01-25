package middlewares

import (
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifySetupToken(ctx *gin.Context) {
	setupToken := ctx.GetHeader("X-Setup-Token")

	if !auth.VerifySetupToken(setupToken) {
		libs.RespondWithMessage(http.StatusUnauthorized, "Request must contain a valid setup-token", ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
