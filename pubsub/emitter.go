package pubsub

import "reactive-go/event"

type EventEmitter interface {
	ValueEvent(e event.Event, payload interface{})
	ErrorEvent(e event.Event, err error)
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

func (publisher *emitter) ValueEvent(e event.Event, payload interface{}) {
	publishEventOccurence(e, payload)
}

func (publisher *emitter) ErrorEvent(e event.Event, err error) {
	publishEventError(e, err)
}