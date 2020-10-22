package types

import "events/events_meta/interfaces"

type Channels struct {
	System interfaces.SystemChannels
	Admin  interfaces.AdminChannels
}
