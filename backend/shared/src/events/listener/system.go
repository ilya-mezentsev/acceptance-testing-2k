package listener

import (
	"time"
)

type System struct {
	cleanExpiredDBConnectionsListeners []func(d time.Duration)
	cleanExpiredAccountHashesListeners []func(d time.Duration)
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

func (s System) EmitCleanExpiredAccountHashes(d time.Duration) {
	for _, listener := range s.cleanExpiredAccountHashesListeners {
		listener(d)
	}
}

func (s *System) CleanExpiredAccountHashes(f func(d time.Duration)) {
	s.cleanExpiredAccountHashesListeners = append(s.cleanExpiredAccountHashesListeners, f)
}
