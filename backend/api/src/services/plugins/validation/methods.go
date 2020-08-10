package validation

import (
	"plugins/logger"
	"regexp"
)

var validationRuleToMethod = map[string]func(string) bool{
	"md5-hash":         IsMd5Hash,
	"regular-name":     IsRegularName,
	"model-field-name": IsFieldName,
}

var (
	regularNamePattern = regexp.MustCompile(`^[-a-zA-Z0-9_]{1,64}$`)
	fieldNamePattern   = regexp.MustCompile(`^(?P<updateTarget>[a-zA-Z_]+?:)?[a-zA-Z0-9_]{1,64}$`)
	md5Pattern         = regexp.MustCompile(`^[a-f0-9]{32}$`)
)

func IsMd5Hash(s string) bool {
	validMd5Hash := md5Pattern.MatchString(s)
	if !validMd5Hash {
		logger.WarningF("Invalid MD5 hash: %s", s)
	}

	return validMd5Hash
}

func IsRegularName(s string) bool {
	validRegularName := regularNamePattern.MatchString(s)
	if !validRegularName {
		logger.WarningF("Invalid regular name: %s", s)
	}

	return validRegularName
}

func IsFieldName(s string) bool {
	validFieldName := fieldNamePattern.MatchString(s)
	if !validFieldName {
		logger.WarningF("Invalid field name: %s", s)
	}

	return validFieldName
}
