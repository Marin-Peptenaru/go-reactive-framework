package pubsub

import (
	"reactive-go/behavior"
	"reactive-go/event"

	"github.com/reactivex/rxgo/v2"
)

type EventReceiver interface {
	GetEventSource(event event.Event) rxgo.Observable
	OnEvent(event event.Event, behavior *behavior.Behavior)
}

type Subscriber interface {
	EventReceiver
}

type Consumer interface {
	EventReceiver
}

type receiver struct {
	strategy PropagationStrategy
}


func (r *receiver) GetEventSource(event event.Event) rxgo.Observable{
	return eventSource(event, r.strategy)
}

func (r *receiver) OnEvent(event event.Event, behavior *behavior.Behavior){
	newEventBehaviour(event, behavior, r.strategy);
}

func NewSubscriber() Subscriber{
	return &receiver{strategy: PUBLISH}
}

func NewConsumer() Consumer {
	return &receiver{strategy: PRODUCE}
}
