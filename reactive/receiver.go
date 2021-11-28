package reactive

import (
	"reactive-go/behavior"
	"reactive-go/event"

	"github.com/reactivex/rxgo/v2"
)

type EventReceiver interface {
	GetEventSource(event event.Event) (rxgo.Observable, error)
	OnEvent(event event.Event, behavior *behavior.Behavior) error
}

type Subscriber interface {
	EventReceiver
}

type Consumer interface {
	EventReceiver
}

type receiver struct {
	reactive
}

func (r *receiver) GetEventSource(e event.Event) (rxgo.Observable, error) {
	if !r.events.Declared(e) {
		return nil, event.UndeclaredEvent(e)
	}

	return eventSource(e, r.strategy), nil
}

func (r *receiver) OnEvent(e event.Event, behavior *behavior.Behavior) error {

	if !r.events.Declared(e) {
		return event.UndeclaredEvent(e)
	}

	newEventBehaviour(e, behavior, r.strategy)
	return nil
}

func NewSubscriber(subsribedEvents event.EventSet) Subscriber {
	return &receiver{
		reactive: reactive{
			strategy: PUBLISH,
			events:   subsribedEvents,
		},
	}
}

func NewConsumer(consumedEvents event.EventSet) Consumer {
	return &receiver{
		reactive: reactive{
			strategy: PRODUCE,
			events:   consumedEvents,
		},
	}
}
