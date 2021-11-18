package events

import (
	"github.com/reactivex/rxgo/v2"
)

type (
	Event string
	Behaviour struct {
		onEvent func(interface{})
		onError func(error)
		onDisposed func()
	}
)

type EventDrivenApplicationCore interface{
	publishEventOccurence(event Event, payload interface{})
	publishEventError(event Event, err error)
	NewEventBehaviour(event Event, behaviour *Behaviour)
	EventSource(event Event) rxgo.Observable
	newEventSource(event Event)

}

func NewBehaviour() *Behaviour {
	return &Behaviour{}
}

func (b *Behaviour) OnEvent(eventHandler func(interface{})) *Behaviour{
	b.onEvent = eventHandler
	return b
}

func (b *Behaviour) OnError(errorHanlder func(error)) *Behaviour{
	b.onError = errorHanlder
	return b
}

func (b *Behaviour) OnDisposed(disposalHandler func()) *Behaviour{
	b.onDisposed = disposalHandler
	return b
}

