package validation

import (
	"fmt"
	"logger"
	"net/url"
	"regexp"
)

var validationRuleToMethod = map[string]func(string) bool{
	"md5-hash":            IsMd5Hash,
	"regular-name":        IsRegularName,
	"model-field-name":    IsFieldName,
	"meaning-http-method": IsMeaningHttpMethod,
	"base-url":            IsValidBaseUrl,
	"endpoint":            IsValidEndpoint,
}

var (
	regularNamePattern       = regexp.MustCompile(`^[-a-zA-Z0-9_]{1,64}$`)
	fieldNamePattern         = regexp.MustCompile(`^(?P<updateTarget>[a-zA-Z_]+?:)?[a-zA-Z0-9_]{1,64}$`)
	md5Pattern               = regexp.MustCompile(`^[a-f0-9]{32}$`)
	meaningHttpMethodPattern = regexp.MustCompile(`POST|GET|PATCH|DELETE`)
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

func IsMeaningHttpMethod(s string) bool {
	validMethod := meaningHttpMethodPattern.MatchString(s)
	if !validMethod {
		logger.WarningF("Invalid http method: %s", s)
	}

	return validMethod
}

func IsValidBaseUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		logger.WarningF("Invalid base url: %s, err: %v", s, err)
	}

	return err == nil
}

func IsValidEndpoint(s string) bool {
	_, err := url.ParseRequestURI(fmt.Sprintf("https://link.com/%s", s))
	if err != nil {
		logger.WarningF("Invalid endpoint: %s, err: %v", s, err)
	}

	return err == nil
}
