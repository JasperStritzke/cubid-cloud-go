package auth_controller

import (
	"bytes"
	"encoding/base64"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/auth"
	"github.com/JasperStritzke/cubid-cloud/controller/rest/libs"
	"github.com/JasperStritzke/cubid-cloud/util/crypto"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"image/png"
	"net/http"
)

var preSetupFinishedAccount *auth.Account

func PreFinishSetup(ctx *gin.Context) {
	var account = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if !libs.BindBody(ctx, &account) {
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "cubid.cloud",
		AccountName: account.Username,
	})

	if err != nil {
		libs.RespondWithMessage(http.StatusBadRequest, "Invalid request parameters", ctx)
		return
	}

	password, err := base64.StdEncoding.DecodeString(account.Password)

	if err != nil {
		libs.RespondWithMessage(http.StatusBadRequest, "Invalid request parameters", ctx)
		return
	}

	img, err := key.Image(256, 256)

	if err != nil {
		libs.RespondWithMessage(http.StatusBadRequest, "Invalid request parameters", ctx)
		return
	}

	preSetupFinishedAccount = &auth.Account{
		Username:     account.Username,
		PasswordHash: crypto.HashPassword(string(password)),
		TOTP:         key.Secret(),
		Permissions: auth.AccountPermissions{
			AdminAnyThing: true,
			Nodes:         nil,
		},
	}

	buff := new(bytes.Buffer)

	err = png.Encode(buff, img)

	if err != nil {
		libs.RespondWithMessage(http.StatusBadRequest, "Invalid request parameters", ctx)
		return
	}

	ctx.JSON(200, gin.H{"totp": base64.StdEncoding.EncodeToString(buff.Bytes())})
}

func FinishSetup(m *auth.Manager) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var totpCode = struct {
			Totp string `json:"totp"`
		}{}

		if !libs.BindBody(ctx, &totpCode) {
			return
		}

		if preSetupFinishedAccount == nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !preSetupFinishedAccount.ValidateOTP(totpCode.Totp) {
			libs.RespondWithMessage(http.StatusUnauthorized, "TOTP was invalid", ctx)
			return
		}

		err := m.CreateAccount(*preSetupFinishedAccount)

		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(200, gin.H{
			"username": preSetupFinishedAccount.Username,
		})

		auth.DeleteSetupToken()
	}
}
