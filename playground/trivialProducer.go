package playground

import (
	"reactive-go/event"
	"reactive-go/pubsub"
	"time"
)

type TrivialProducer struct {
	trivialEventProducer pubsub.Producer
}

func NewTrivialProducer() *TrivialProducer{

	return &TrivialProducer{
		trivialEventProducer: pubsub.NewProducer(),
	}
}

func (t *TrivialProducer) StartProducing(){
	go func(){
		ch := time.Tick(1 * time.Second)
		i := 0
		for {
			<- ch
			i++
			t.trivialEventProducer.ValueEvent(event.Event("Trivial"), i)
		}
	}()
}