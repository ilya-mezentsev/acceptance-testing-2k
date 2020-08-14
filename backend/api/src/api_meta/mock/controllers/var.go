package controllers

import "services/plugins/hash"

var (
	BadAccountHash = hash.GetHashWithTimeAsKey("bad-account-hash")
	BadEntityHash  = hash.GetHashWithTimeAsKey("bad-entity-hash")
)
