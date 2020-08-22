package csrf

import (
	"fmt"
	"logger"
	"middlewares/plugins/base64"
	"net/http"
	"services/plugins/hash"
	"strings"
	"time"
)

type Middleware struct {
	PrivateKey string
}

func (m Middleware) CheckCSRFToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieToken, err := m.getTokenFromCookie(r)
		if err != nil {
			logger.WarningF("Cannot get csrf token from cookie or decode it")
			http.Error(w, "CSRF token cookies error", http.StatusForbidden)
			return
		}

		headerToken, err := m.getTokenFromHeader(r)
		if err != nil {
			logger.WarningF("Cannot get csrf token from headers or decode it")
			http.Error(w, "CSRF token headers error", http.StatusForbidden)
			return
		}

		if cookieToken == headerToken {
			next.ServeHTTP(w, r)
			m.setTokens(w)
		} else {
			logger.Warning("CSRF tokens didn't match")
			http.Error(w, "CSRF tokens didn't match", http.StatusForbidden)
		}
	})
}

func (m Middleware) getTokenFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(csrfKey)
	if err != nil {
		return "", err
	}

	token, err := base64.Decode(cookie.Value)
	if err != nil {
		return "", err
	}

	return strings.Split(token, keysSeparator)[0], nil
}

func (m Middleware) getTokenFromHeader(r *http.Request) (string, error) {
	token, err := base64.Decode(r.Header.Get(csrfKey))
	if err != nil {
		return "", err
	}

	return strings.Split(token, keysSeparator)[0], nil
}

func (m Middleware) setTokens(w http.ResponseWriter) {
	publicKey := base64.Encode(fmt.Sprintf("%d", time.Now().Unix()))
	token := base64.Encode(strings.Join([]string{
		hash.Md5(m.PrivateKey),
		keysSeparator,
		hash.Md5(publicKey),
	}, keysSeparator))

	w.Header().Add(csrfPublicTokenKey, publicKey)
	http.SetCookie(w, &http.Cookie{
		Name:     csrfKey,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}
