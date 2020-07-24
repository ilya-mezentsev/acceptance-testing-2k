package types

import "net/http"

type (
	Settings interface {
		GetMethod() string
		GetBaseURL() string
		GetEndpoint() string
		ShouldPassArgumentsInURL() bool
		GetHeaders() map[string]string
		GetCookies() []*http.Cookie
	}

	Arguments interface {
		Value() string
		IsSlashSeparated() bool
		AmpersandSeparated() (string, error)
	}
)
