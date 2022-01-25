package auth

import (
	"github.com/AlexanderGrom/go-event"
	"github.com/JasperStritzke/cubid-cloud/controller/event_names"
	"github.com/JasperStritzke/cubid-cloud/util/cache"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/random"
	"time"
)

const timeout = 30 * time.Minute

var sessions = cache.NewCache(timeout)

func init() {
	_ = event.On(event_names.Shutdown, sessions.Shutdown)
}

type Session struct {
	Account *Account
	Token   string
}

func GenerateSession(account *Account) string {
	token, err := random.GenerateRandomString(64)

	if err != nil {
		return ""
	}

	session := &Session{
		Account: account,
		Token:   token,
	}

	Invalidate(account.Username)

	sessions.Put(token, session)
	sessions.Put(account.Username, session)

	logger.Warnf("User %s logged in.", account.Username)

	return token
}

func ExistsToken(token string) bool {
	return sessions.Exists(token)
}

func HasPermission(token, node string) bool {
	return HasPermissions(token, node)
}

func HasPermissions(token string, nodes ...string) bool {
	var session *Session
	found := sessions.Get(token, &session)

	if found {
		if session.Account.Permissions.AdminAnyThing {
			return true
		}

		for _, node := range nodes {
			if !session.Account.Permissions.CheckNode(node) {
				return false
			}
		}

		return true
	}

	return false
}

func Invalidate(usernameOrToken string) {
	var session *Session

	if sessions.Get(usernameOrToken, &session) {
		sessions.Invalidate(session.Token)
		sessions.Invalidate(session.Account.Username)
	}
}
