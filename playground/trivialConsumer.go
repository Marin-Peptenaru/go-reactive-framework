package playground

import (
	"fmt"
	"reactive-go/behavior"
	"reactive-go/event"
	"reactive-go/pubsub"
)

type TrivialConsumer struct {
	name string
	consumer pubsub.Consumer
}


func NewTrivialConsumer(name string) *TrivialConsumer{
	return &TrivialConsumer{
		name: name, consumer: pubsub.NewConsumer(event.NewEventSet(event.Event("Trivial"))),
	}
}

func (t *TrivialConsumer) StartConsuming(){

	b := behavior.New()

	b.OnEvent = func(event interface{}){fmt.Printf("%s %s\n", t.name, event)}

	err := t.consumer.OnEvent(event.Event("abc"), b)
	if err != nil {
		fmt.Println(err.Error())
	}
}