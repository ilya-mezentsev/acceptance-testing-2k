package emitter

import (
	"events/events_meta/env"
	"events/events_meta/types"
	"time"
)

type System struct {
	messages chan types.Message
}

func (s System) CleanExpiredDBConnections(d time.Duration) {
	s.messages <- types.Message{
		EventName: env.SystemCleanExpiredDBConnections,
		Data:      d.Seconds(),
	}
}

func (s System) CleanExpiredDeletedAccountHashes(d time.Duration) {
	s.messages <- types.Message{
		EventName: env.SystemCleanExpiredAccountHashes,
		Data:      d.Seconds(),
	}
}
