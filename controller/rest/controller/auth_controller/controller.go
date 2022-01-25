package auth_controller

import (
	"encoding/base64"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckSetupMode(m *auth.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		authValid := false
		if len(authorization) > 0 {
			authValid = auth.ExistsToken(authorization)
		}

		ctx.JSON(200, gin.H{"setupMode": m.RequiresSetup(), "authValid": authValid})
	}
}

type loginRequest struct {
	Username string `json:"a"`
	Password string `json:"b"`
	Otp      string `json:"c"`
}

func (r loginRequest) Decoded(input string) string {
	bytes, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		return ""
	}

	return string(bytes)
}

func Login(m *auth.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request loginRequest

		if !libs.BindBody(ctx, &request) {
			return
		}

		username := request.Decoded(request.Username)
		password := request.Decoded(request.Password)
		otp := request.Otp

		sessionToken := m.Login(
			username,
			password,
			otp,
		)

		if sessionToken == "" {
			libs.RespondWithMessage(http.StatusUnauthorized, "Invalid credentials", ctx)
			return
		}

		ctx.JSON(
			http.StatusOK, gin.H{"token": sessionToken},
		)
	}
}

func Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	auth.Invalidate(authHeader)
}
