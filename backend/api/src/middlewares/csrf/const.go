package csrf

const (
	websocketRequestKey = "X-WS-Request"
	csrfKey             = "X-CSRF-Token"
	csrfPublicTokenKey  = "X-CSRF-Public-Token"
	keysSeparator       = "|"
)
