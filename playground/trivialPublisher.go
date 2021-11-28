package playground

import (
	"fmt"
	"reactive-go/event"
	"reactive-go/reactive"
	"time"
)

type TrivialPublisher struct {
	trivialEventPublisher reactive.Publisher
}

func NewTrivialPublisher() *TrivialPublisher{

	return &TrivialPublisher{
		trivialEventPublisher: reactive.NewPublisher(TrivialEventSet),
	}
}

func (t *TrivialPublisher) StartPublishing(){
	go func(){
		ch := time.Tick(1 * time.Second)
		i := 0
		for {
			<- ch
			i++
			if err := t.trivialEventPublisher.ValueEvent(event.Event("Hello"), i); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}