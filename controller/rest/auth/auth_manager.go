package auth

import (
	"errors"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"os"
	"path/filepath"
)

const authFile = "acp-auth.gob"

type Manager struct {
	Accounts []Account
}

func (m *Manager) RequiresSetup() bool {
	return len(m.Accounts) == 0
}

func (m *Manager) CreateAccount(account Account) error {
	m.Accounts = append(m.Accounts, account)
	return m.save()
}

func (m *Manager) DeleteAccount(email string) error {
	foundIndex, _ := m.GetAccount(email)

	if foundIndex == -1 {
		return errors.New("account not found")
	}

	m.Accounts[foundIndex] = m.Accounts[len(m.Accounts)-1] //Copy last element to element where account usually was
	m.Accounts[len(m.Accounts)-1] = Account{}
	m.Accounts = m.Accounts[:len(m.Accounts)-1]
	return nil
}

func (m *Manager) GetAccount(email string) (int, *Account) {
	for index, account := range m.Accounts {
		if account.Email == email {
			return index, &account
		}
	}

	return -1, nil
}

func (m *Manager) save() error {
	return config.WriteConfig(authFile, fileutil.CodingGOB, m)
}

func NewManager() *Manager {
	_ = config.InitConfigIfNotExists(authFile, fileutil.CodingGOB, Manager{Accounts: nil})

	var manager Manager
	_ = config.LoadConfig(authFile, fileutil.CodingGOB, &manager)

	if manager.RequiresSetup() {
		logger.Warn("No control-panel Account exists yet. You can create one by using the setup token from the \"setupkey\" file.")
		GenerateSetupToken()

		wd, err := os.Getwd()

		if err != nil {
			logger.Infof("Path to setupkey file: %s", filepath.Join(wd, "setupkey"))
		}
	}

	return &manager
}
