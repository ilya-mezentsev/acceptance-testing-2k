package types

type (
	Arguments interface {
		Value() string
		AmpersandSeparated() (string, error)
	}
)
