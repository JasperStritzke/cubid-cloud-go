package auth

import (
	"errors"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"os"
	"path/filepath"
)

const authFile = "auth"

type Manager struct {
	Accounts []Account
}

func (m *Manager) RequiresSetup() bool {
	return len(m.Accounts) == 0
}

func (m *Manager) CreateAccount(account Account) error {
	if index, _ := m.GetAccount(account.Username); index != -1 {
		return errors.New("account already exists")
	}

	logger.Infof("Account %s was created", account.Username)

	m.Accounts = append(m.Accounts, account)
	return m.save()
}

func (m *Manager) DeleteAccount(username string) error {
	foundIndex, _ := m.GetAccount(username)

	if foundIndex == -1 {
		return errors.New("account not found")
	}

	m.Accounts[foundIndex] = m.Accounts[len(m.Accounts)-1] //Copy last element to element where account usually was
	m.Accounts[len(m.Accounts)-1] = Account{}              //Set last value to Account zero value
	m.Accounts = m.Accounts[:len(m.Accounts)-1]            //Slice the slice :0 (remove last value)
	return nil
}

func (m *Manager) GetAccount(username string) (int, *Account) {
	if m.Accounts != nil {
		for index, account := range m.Accounts {
			if account.Username == username {
				return index, &account
			}
		}
	}

	return -1, nil
}

func (m *Manager) save() error {
	return config.WriteConfig(authFile, fileutil.CodingGOB, m)
}

func (m *Manager) Login(username, password, otp string) string {
	index, account := m.GetAccount(username)

	if index == -1 {
		return ""
	}

	if account.ValidatePassword(password) && account.ValidateOTP(otp) {
		return GenerateSession(account)
	}

	return ""
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
