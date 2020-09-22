package account_hash

import (
	"logger"
	"net/http"
)

func ExtractFromRequest(r *http.Request) string {
	accountHash, err := r.Cookie("AAT-Session")
	if err != nil {
		logger.WarningF("AAT-Session cookie is not provided")
		return ""
	}

	return accountHash.Value
}
