package libs

import "github.com/gin-gonic/gin"

func RespondWithMessage(code int, message string, ctx *gin.Context) {
	ctx.JSON(code, gin.H{"message": message})
}

func ResponseStaticMessage(msg string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": msg})
	}
}
