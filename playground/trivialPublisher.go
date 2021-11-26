package playground

import (
	"reactive-go/event"
	"reactive-go/pubsub"
	"time"
)

type TrivialPublisher struct {
	trivialEventPublisher pubsub.Publisher
}

func NewTrivialPublisher() *TrivialPublisher{

	return &TrivialPublisher{
		trivialEventPublisher: pubsub.NewPublisher(),
	}
}

func (t *TrivialPublisher) StartPublishing(){
	go func(){
		ch := time.Tick(1 * time.Second)
		i := 0
		for {
			<- ch
			i++
			t.trivialEventPublisher.ValueEvent(event.Event("Trivial"), i)
		}
	}()
}