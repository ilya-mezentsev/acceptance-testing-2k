package interfaces

import "time"

// System event/channels
type (
	SystemEvents interface {
		CleanExpiredDBConnections(d time.Duration)
	}

	SystemChannels interface {
		CleanExpiredDBConnections(func(d time.Duration))
	}
)
