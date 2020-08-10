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
)
