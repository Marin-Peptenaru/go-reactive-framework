package pubsub

import (
	"reactive-go/behavior"
	"reactive-go/event"

	"github.com/reactivex/rxgo/v2"
)

type Subscriber interface {
	GetEventSource(event event.Event) rxgo.Observable
	OnEvent(event event.Event, behavior *behavior.Behavior)
}

type subscriberImpl struct {

}


func (s *subscriberImpl) GetEventSource(event event.Event) rxgo.Observable{
	return eventSource(event)
}

func (s *subscriberImpl) OnEvent(event event.Event, behaviour *behavior.Behavior){
	newEventBehaviour(event, behaviour);
}

func NewSubscriber() Subscriber{
	return &subscriberImpl{}
}