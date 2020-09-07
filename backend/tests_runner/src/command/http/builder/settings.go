package builder

import "net/http"

type (
	settings struct {
		Method             string `db:"method"`
		BaseURL            string `db:"base_url"`
		Endpoint           string `db:"endpoint"`
		Timeout            int    `db:"timeout"`
		PassArgumentsInURL bool   `db:"pass_arguments_in_url"`
		Headers            []keyValueMapping
		Cookies            []keyValueMapping
	}

	keyValueMapping struct {
		Key   string `db:"key"`
		Value string `db:"value"`
	}
)

func (s settings) GetMethod() string {
	return s.Method
}

func (s settings) GetBaseURL() string {
	return s.BaseURL
}

func (s settings) GetEndpoint() string {
	return s.Endpoint
}

func (s settings) GetTimeout() int {
	return s.Timeout
}

func (s settings) ShouldPassArgumentsInURL() bool {
	return s.PassArgumentsInURL
}

func (s settings) GetHeaders() map[string]string {
	headers := map[string]string{}
	for _, header := range s.Headers {
		headers[header.Key] = header.Value
	}

	return headers
}

func (s settings) GetCookies() []*http.Cookie {
	var cookies []*http.Cookie
	for _, cookie := range s.Cookies {
		cookies = append(cookies, &http.Cookie{
			Name:  cookie.Key,
			Value: cookie.Value,
		})
	}

	return cookies
}
