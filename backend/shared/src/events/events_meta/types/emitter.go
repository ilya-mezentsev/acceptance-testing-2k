package types

import "events/events_meta/interfaces"

type Events struct {
	System interfaces.SystemEvents
	Admin  interfaces.AdminEvents
}
