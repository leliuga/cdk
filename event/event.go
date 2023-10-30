package event

// NewEvent creates a new event.
func NewEvent(options *Options) *Event {
	return &Event{Options: options}
}
