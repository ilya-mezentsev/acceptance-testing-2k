package services

import (
	"api_meta/models"
	"services/plugins/hash"
)

var (
	SomeHash              = hash.GetHashWithTimeAsKey("some-hash")
	BadAccountHash        = hash.GetHashWithTimeAsKey("bad-hash")
	PredefinedAccountHash = hash.GetHashWithTimeAsKey("account-hash-1")
	PredefinedTestObject1 = models.TestObject{
		Name: "USER",
		Hash: hash.GetHashWithTimeAsKey("object-hash-1"),
	}
	PredefinedTestObject2 = models.TestObject{
		Name: "ACTION",
		Hash: hash.GetHashWithTimeAsKey("object-hash-2"),
	}
	PredefinedTestCommand1 = models.TestCommandRecord{
		CommandSettings: models.CommandSettings{
			Name:               "GET",
			Hash:               hash.GetHashWithTimeAsKey("command-hash-1"),
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
			Hash:               hash.GetHashWithTimeAsKey("command-hash-1"),
			ObjectName:         "SETTINGS",
			Method:             "POST",
			BaseURL:            "https://link.com/api/v2",
			Endpoint:           "user/settings",
			PassArgumentsInURL: false,
		},
		Headers: "X-Header-1=x_value1;X-Header-2=x_value2",
		Cookies: "Cookie-1=some-data;Cookie-2=value",
	}
)
