package listener

type Admin struct {
	deleteAccountListeners []func(accountHash string)
}

func NewAdminChannel() *Admin {
	return &Admin{deleteAccountListeners: []func(accountHash string){}}
}

func (a Admin) EmitDeleteAccount(accountHash string) {
	for _, listener := range a.deleteAccountListeners {
		listener(accountHash)
	}
}

func (a *Admin) DeleteAccount(f func(accountHash string)) {
	a.deleteAccountListeners = append(a.deleteAccountListeners, f)
}
