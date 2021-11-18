package playground

import (
	"fmt"
	"reactive-go/events"
)

type TrivialSubscriber struct {
	name string
	subscriber events.Subscriber
}


func NewTrivialSubscriber(name string) *TrivialSubscriber{
	return &TrivialSubscriber{
		name: name, subscriber: events.NewSubscriber(),
	}
}

func (t *TrivialSubscriber) StartSubscribing(){
	t.subscriber.OnEvent(events.Event("Trivial"), events.NewBehaviour().OnEvent(
		func(event interface{}){fmt.Printf("%s %s\n", t.name, event)}))
}