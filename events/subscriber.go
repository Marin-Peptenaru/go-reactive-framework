package events

import "github.com/reactivex/rxgo/v2"

type Subscriber interface {
	GetEventSource(event Event) rxgo.Observable
	OnEvent(event Event, behaviour *Behaviour)
}

type subscriberImpl struct {

}


func (s *subscriberImpl) GetEventSource(event Event) rxgo.Observable{
	return eventDrivenApplication().EventSource(event)
}

func (s *subscriberImpl) OnEvent(event Event, behaviour *Behaviour){
	eventDrivenApplication().NewEventBehaviour(event, behaviour);
}

func NewSubscriber() Subscriber{
	return &subscriberImpl{}
}