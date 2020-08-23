package account_credentials

import "services/plugins/hash"

func GenerateAccountHash(login string) string {
	return hash.Md5(login)
}

func GenerateAccountPassword(login, password string) string {
	// trickster
	return hash.Md5(login + GenerateAccountHash(login) + password)
}
