package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SinglePageApp(urlPrefix, spaDirectory, apiPath string) gin.HandlerFunc {
	dir := static.LocalFile(spaDirectory, true)

	fileServer := http.FileServer(dir)

	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}

	return func(ctx *gin.Context) {
		if strings.HasPrefix(strings.ToLower(ctx.Request.URL.Path), apiPath) {
			ctx.Next()
			return
		}

		if dir.Exists(urlPrefix, ctx.Request.URL.Path) {
			fileServer.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		} else {
			ctx.Request.URL.Path = "/"
			fileServer.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
		}
	}
}
