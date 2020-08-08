package command

import (
	"net/http"
)

type MockSettings struct {
	Method, BaseURL, Endpoint string
	PassArgumentsInURL        bool
	Headers                   map[string]string
	Cookies                   []*http.Cookie
}

func (s MockSettings) GetMethod() string {
	return s.Method
}

func (s MockSettings) GetBaseURL() string {
	return s.BaseURL
}

func (s MockSettings) GetEndpoint() string {
	return s.Endpoint
}

func (s MockSettings) GetHeaders() map[string]string {
	return s.Headers
}

func (s MockSettings) GetCookies() []*http.Cookie {
	return s.Cookies
}

func (s MockSettings) ShouldPassArgumentsInURL() bool {
	return s.PassArgumentsInURL
}
