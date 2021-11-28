package playground

import (
	"fmt"
	"reactive-go/event"
	"reactive-go/pubsub"
	"time"
)

type TrivialProducer struct {
	trivialEventProducer pubsub.Producer
}

func NewTrivialProducer() *TrivialProducer{

	return &TrivialProducer{
		trivialEventProducer: pubsub.NewProducer(TrivialEventSet),
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