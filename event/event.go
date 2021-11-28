package event

type Event string

type EventSet interface {
	Declared(Event) bool
}

type eventSet struct {
	events map[Event]bool
}

func NewEventSet(events ...Event) EventSet {
	set := &eventSet{events: make(map[Event]bool)}

	for _, e := range events {
		set.events[e] = true
	}

	return set
}

func (set *eventSet) Declared(e Event) bool {
	_, declared := set.events[e]
	return declared
}
