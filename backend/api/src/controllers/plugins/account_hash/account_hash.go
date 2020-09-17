package account_hash

import "net/http"

func ExtractFromRequest(r *http.Request) string {
	return r.Header.Get("AAT-Account-Hash")
}
