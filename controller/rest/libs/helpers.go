package libs

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// BindBody returns true if binding was successful
func BindBody(ctx *gin.Context, obj interface{}) bool {
	if err := ctx.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		RespondWithMessage(http.StatusBadRequest, "Invalid request parameters", ctx)
		return false
	}

	return true
}
