package pubsub

import "reactive-go/event"

type EventEmitter interface {
	ValueEvent(event.Event, interface{})
	ErrorEvent(event.Event, error)
}

type Publisher interface {
	EventEmitter
}

type Producer interface {
	EventEmitter
}

type emitter struct{
	strategy PropagationStrategy
}

func NewPublisher(publishedEvents ...event.Event) Publisher {
	return &emitter{strategy: PUBLISH}
}

func NewProducer() Producer{
	return &emitter{strategy: PRODUCE}
}


func (em *emitter) ValueEvent(e event.Event, payload interface{}){
	publishEventOccurence(e, payload, em.strategy )
}

func (em * emitter)	ErrorEvent(e event.Event, err error){
	publishEventError(e, err, em.strategy)
}