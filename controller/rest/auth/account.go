package auth

import (
	"github.com/JasperStritzke/cubid-cloud/util/crypto"
	"github.com/pquerna/otp/totp"
)

type Account struct {
	Username     string
	PasswordHash []byte
	TOTP         string
	Permissions  AccountPermissions
}

func (a *Account) ValidatePassword(providedPassword string) bool {
	return crypto.CheckPasswordHash(providedPassword, a.PasswordHash)
}

func (a *Account) ValidateOTP(otp string) bool {
	return totp.Validate(otp, a.TOTP)
}

type AccountPermissions struct {
	AdminAnyThing bool
	Nodes         []string
}

func (ap *AccountPermissions) CheckNode(node string) bool {
	if ap.AdminAnyThing {
		return true
	}

	for _, current := range ap.Nodes {
		if current == node {
			return true
		}
	}

	return false
}
