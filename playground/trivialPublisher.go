package playground

import (
	"reactive-go/events"
	"time"
)

type TrivialPublisher struct {
	trivialEventPublisher events.Publisher
}

func NewTrivialPublisher() *TrivialPublisher{

	return &TrivialPublisher{
		trivialEventPublisher: events.NewPublisher(),
	}
}

func (t *TrivialPublisher) StartPublishing(){
	go func(){
		ch := time.Tick(1 * time.Second)
		for {
			<- ch
			t.trivialEventPublisher.ValueEvent(events.Event("Trivial"), "tick")
		}
	}()
}