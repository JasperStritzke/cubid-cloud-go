package auth

type Account struct {
	Email        string
	PasswordHash string
	PasswordSalt string
	TOTP         string
	Permissions  AccountPermissions
}

type AccountPermissions struct {
	AdminAnyThing bool
	Nodes         []string
}
