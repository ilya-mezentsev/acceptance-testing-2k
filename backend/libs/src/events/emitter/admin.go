package emitter

import (
	"events/events_meta/env"
	"events/events_meta/types"
)

type Admin struct {
	messages chan types.Message
}

func (a Admin) DeleteAccount(accountHash string) {
	a.messages <- types.Message{
		EventName: env.AdminDeleteAccount,
		Data:      accountHash,
	}
}
