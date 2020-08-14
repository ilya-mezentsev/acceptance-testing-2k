package account_credentials

import "services/plugins/hash"

func GenerateAccountHash(login, password string) string {
	return hash.Md5(hash.Md5(login) + hash.Md5(password))
}

func GenerateAccountPassword(login, password string) string {
	// trickster
	return hash.Md5(login + GenerateAccountHash(login, password) + password)
}
