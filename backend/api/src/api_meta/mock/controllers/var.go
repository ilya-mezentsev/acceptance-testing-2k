package controllers

import "services/plugins/hash"

var (
	BadAccountHash = hash.Md5WithTimeAsKey("bad-account-hash")
	BadEntityHash  = hash.Md5WithTimeAsKey("bad-entity-hash")
)
