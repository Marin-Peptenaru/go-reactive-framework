package playground

import (
	"fmt"
	"reactive-go/event"
	"reactive-go/reactive"
	"time"
)

type TrivialProducer struct {
	trivialEventProducer reactive.Producer
}

func NewTrivialProducer() *TrivialProducer{

	return &TrivialProducer{
		trivialEventProducer: reactive.NewProducer(TrivialEventSet),
	}
}

func (t *TrivialProducer) StartProducing(){
	go func(){
		ch := time.Tick(1 * time.Second)
		i := 0
		for {
			<- ch
			i++
			err := t.trivialEventProducer.ValueEvent(event.Event("hello"), i)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}