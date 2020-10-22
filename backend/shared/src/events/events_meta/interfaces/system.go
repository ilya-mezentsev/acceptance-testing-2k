package interfaces

import "time"

// System event/channels
type (
	SystemEvents interface {
		CleanExpiredDBConnections(d time.Duration)
		CleanExpiredDeletedAccountHashes(d time.Duration)
	}

	SystemChannels interface {
		CleanExpiredDBConnections(func(d time.Duration))
		CleanExpiredAccountHashes(func(d time.Duration))
	}
)
