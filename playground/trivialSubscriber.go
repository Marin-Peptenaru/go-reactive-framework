package playground

import (
	"fmt"
	"reactive-go/behavior"
	"reactive-go/event"
	"reactive-go/pubsub"
)

type TrivialSubscriber struct {
	name string
	subscriber pubsub.Subscriber
}


func NewTrivialSubscriber(name string) *TrivialSubscriber{
	return &TrivialSubscriber{
		name: name, subscriber: pubsub.NewSubscriber(TrivialEventSet),
	}
}

func (t *TrivialSubscriber) StartSubscribing(){

	b := behavior.New()

	b.OnEvent = func(event interface{}){fmt.Printf("%s %s\n", t.name, event)}

	if err := t.subscriber.OnEvent(event.Event("abcd"), b); err != nil {
		fmt.Println(err.Error())
	}
}