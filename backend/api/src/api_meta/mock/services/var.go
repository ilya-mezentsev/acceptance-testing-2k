package services

import (
	"api_meta/models"
	"services/plugins/account_credentials"
	"services/plugins/hash"
	"test_case_runner"
)

var (
	ExistsLogin            = "exists-login"
	ExistsPassword         = "exists-password"
	BadLogin               = "bad-login"
	BadPassword            = "bad-password"
	SomeHash               = hash.Md5WithTimeAsKey("some-hash")
	BadAccountHash         = account_credentials.GenerateAccountHash(BadLogin)
	ExistsAccountHash      = account_credentials.GenerateAccountHash(ExistsLogin)
	PredefinedAccountHash  = hash.Md5WithTimeAsKey("account-hash-1")
	PredefinedCommandHash1 = hash.Md5WithTimeAsKey("command-hash-1")
	PredefinedCommandHash2 = hash.Md5WithTimeAsKey("command-hash-2")
	PredefinedHeaderHash1  = hash.Md5WithTimeAsKey("header-hash-1")
	PredefinedHeaderHash2  = hash.Md5WithTimeAsKey("header-hash-2")
	PredefinedCookieHash1  = hash.Md5WithTimeAsKey("cookie-hash-1")
	PredefinedCookieHash2  = hash.Md5WithTimeAsKey("cookie-hash-2")
	PredefinedTestObject1  = models.TestObject{
		Name: "USER",
		Hash: hash.Md5WithTimeAsKey("object-hash-1"),
	}
	PredefinedTestObject2 = models.TestObject{
		Name: "ACTION",
		Hash: hash.Md5WithTimeAsKey("object-hash-2"),
	}
	PredefinedTestCommand1 = models.CommandSettings{
		Name:               "GET",
		Hash:               PredefinedCommandHash1,
		ObjectHash:         "SETTINGS",
		Method:             "GET",
		BaseURL:            "https://link.com/api/v2",
		Endpoint:           "user/settings",
		Timeout:            3,
		PassArgumentsInURL: true,
	}
	PredefinedTestCommand2 = models.CommandSettings{
		Name:               "CREATE",
		Hash:               PredefinedCommandHash2,
		ObjectHash:         "SETTINGS",
		Method:             "POST",
		BaseURL:            "https://link.com/api/v2",
		Endpoint:           "user/settings",
		Timeout:            3,
		PassArgumentsInURL: false,
	}
	PredefinedHeader1 = models.KeyValueMapping{
		Hash:        PredefinedHeaderHash1,
		Key:         "Key-1",
		Value:       "Value-1",
		CommandHash: PredefinedCommandHash1,
	}
	PredefinedHeader2 = models.KeyValueMapping{
		Hash:        PredefinedHeaderHash2,
		Key:         "Key-2",
		Value:       "Value-2",
		CommandHash: PredefinedCommandHash2,
	}
	PredefinedCookie1 = models.KeyValueMapping{
		Hash:        PredefinedCookieHash1,
		Key:         "Key-1",
		Value:       "Value-1",
		CommandHash: PredefinedCommandHash1,
	}
	PredefinedCookie2 = models.KeyValueMapping{
		Hash:        PredefinedCookieHash2,
		Key:         "Key-2",
		Value:       "Value-2",
		CommandHash: PredefinedCommandHash2,
	}
	MockTestCasesReport = &test_case_runner.TestsReport{
		Report: &test_case_runner.TestCaseRunReport{
			PassedCount: 3,
			FailedCount: 4,
		},
	}
)
