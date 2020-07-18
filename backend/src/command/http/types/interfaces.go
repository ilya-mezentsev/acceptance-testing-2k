package types

import "net/http"

type (
	Settings interface {
		GetMethod() string
		GetBaseURL() string
		GetEndpoint() string
		GetHeaders() map[string]string
		GetCookies() []http.Cookie
	}

	Arguments interface {
		Value() string
		AmpersandSeparated() (string, error)
	}
)
