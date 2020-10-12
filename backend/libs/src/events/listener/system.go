package listener

import (
	"time"
)

type System struct {
	cleanExpiredDBConnectionsListeners []func(d time.Duration)
}

func NewSystemChannel() *System {
	return &System{cleanExpiredDBConnectionsListeners: []func(d time.Duration){}}
}

func (s System) EmitCleanExpiredDBConnections(d time.Duration) {
	for _, listener := range s.cleanExpiredDBConnectionsListeners {
		listener(d)
	}
}

func (s *System) CleanExpiredDBConnections(f func(d time.Duration)) {
	s.cleanExpiredDBConnectionsListeners = append(s.cleanExpiredDBConnectionsListeners, f)
}
