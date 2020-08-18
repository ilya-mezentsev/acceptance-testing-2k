package session

import (
	"net/http"
	"time"
)

func setSessionCookie(w http.ResponseWriter, accountHash string) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    accountHash,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
}

func getSessionCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(CookieName)
}

func unsetSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
