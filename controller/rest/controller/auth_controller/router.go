package auth_controller

import (
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func Mount(r *gin.RouterGroup, manager *auth.Manager) {
	r.GET("/checkSetupMode", CheckSetupMode(manager))
	r.POST("/pre-setup", middlewares.VerifySetupToken, PreFinishSetup)
	r.POST("/setup", middlewares.VerifySetupToken, FinishSetup(manager))
	r.POST("/login", Login(manager))
	r.POST("/logout", middlewares.RequireLoggedIn, Logout)
}
