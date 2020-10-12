package interfaces

type (
	AccountsRepository interface {
		GetNonVerifiedAccountsCreatedDayAgo() ([]string, error)
		DeleteAccounts(hashes []string) error
	}
)
