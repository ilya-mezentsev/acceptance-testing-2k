package services

import (
	"api_meta/models"
	"services/plugins/account_credentials"
	"services/plugins/hash"
	"test_case_runner"
)

var (
	ExistsLogin           = "exists-login"
	ExistsPassword        = "exists-password"
	BadLogin              = "bad-login"
	BadPassword           = "bad-password"
	SomeHash              = hash.Md5WithTimeAsKey("some-hash")
	BadAccountHash        = account_credentials.GenerateAccountHash(BadLogin)
	ExistsAccountHash     = account_credentials.GenerateAccountHash(ExistsLogin)
	PredefinedAccountHash = hash.Md5WithTimeAsKey("account-hash-1")
	PredefinedTestObject1 = models.TestObject{
		Name: "USER",
		Hash: hash.Md5WithTimeAsKey("object-hash-1"),
	}
	PredefinedTestObject2 = models.TestObject{
		Name: "ACTION",
		Hash: hash.Md5WithTimeAsKey("object-hash-2"),
	}
	PredefinedTestCommand1 = models.TestCommandRecord{
		CommandSettings: models.CommandSettings{
			Name:               "GET",
			Hash:               hash.Md5WithTimeAsKey("command-hash-1"),
			ObjectName:         "SETTINGS",
			Method:             "GET",
			BaseURL:            "https://link.com/api/v2",
			Endpoint:           "user/settings",
			PassArgumentsInURL: true,
		},
		Headers: "X-Header-1=x_value1;X-Header-2=x_value2",
		Cookies: "Cookie-1=some-data;Cookie-2=value",
	}
	PredefinedTestCommand2 = models.TestCommandRecord{
		CommandSettings: models.CommandSettings{
			Name:               "CREATE",
			Hash:               hash.Md5WithTimeAsKey("command-hash-1"),
			ObjectName:         "SETTINGS",
			Method:             "POST",
			BaseURL:            "https://link.com/api/v2",
			Endpoint:           "user/settings",
			PassArgumentsInURL: false,
		},
		Headers: "X-Header-1=x_value1;X-Header-2=x_value2",
		Cookies: "Cookie-1=some-data;Cookie-2=value",
	}
	MockTestCasesReport = &test_case_runner.TestsReport{
		Report: &test_case_runner.TestCaseRunReport{
			PassedCount: 3,
			FailedCount: 4,
		},
	}
)
