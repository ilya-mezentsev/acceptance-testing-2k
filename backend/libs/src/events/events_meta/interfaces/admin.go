package interfaces

// Admin event/channels
type (
	AdminEvents interface {
		DeleteAccount(accountHash string)
	}

	AdminChannels interface {
		DeleteAccount(func(accountHash string))
	}
)
