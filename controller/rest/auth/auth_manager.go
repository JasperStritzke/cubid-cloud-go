package auth

import (
	"github.com/JasperStritzke/cubid-cloud/util/crypto"
)

const authFile = ".auth"

type Manager struct {
	Admin        *Account
	setupKeyHash string
}

func NewManager() *Manager {
	manager := &Manager{
		Admin:        nil,
		setupKeyHash: "",
	}

	return manager
}

func (m *Manager) CheckSetupKey(otherKey string) bool {
	return crypto.CheckPasswordHash(otherKey, "", m.setupKeyHash)
}

func (m *Manager) SetSetupKey(newKey string) {
	m.setupKeyHash = crypto.HashPassword(newKey, "")
}

type Account struct {
	Username     string
	passwordHash string
}

func (a *Account) CheckPassword(otherPassword string) bool {
	return crypto.CheckPasswordHash(otherPassword, "", a.passwordHash)
}

func (a *Account) SetPassword(newPassword string) {
	a.passwordHash = crypto.HashPassword(newPassword, "")
}
