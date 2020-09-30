package account_hash

import (
	"logger"
	"net/http"
)

func ExtractFromRequest(r *http.Request) string {
	accountHash, err := r.Cookie("AT2K-Session")
	if err != nil {
		logger.WarningF("AT2K-Session cookie is not provided")
		return ""
	}

	return accountHash.Value
}
