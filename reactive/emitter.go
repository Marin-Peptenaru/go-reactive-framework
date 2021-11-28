package reactive

import "reactive-go/event"

type EventEmitter interface {
	ValueEvent(event.Event, interface{}) error
	ErrorEvent(event.Event, error) error
}

type Publisher interface {
	EventEmitter
}

type Producer interface {
	EventEmitter
}

type emitter struct {
	reactive
}

func NewPublisher(publishedEvents event.EventSet) Publisher {
	return &emitter{
		reactive: reactive{
			strategy: PUBLISH,
			events:   publishedEvents,
		},
	}
}

func NewProducer(producedEvents event.EventSet) Producer {
	return &emitter{
		reactive: reactive{
			strategy: PRODUCE,
			events:   producedEvents,
		},
	}
}

func (em *emitter) ValueEvent(e event.Event, payload interface{}) error {

	if !em.events.Declared(e) {
		return event.UndeclaredEvent(e)
	}

	publishEventOccurence(e, payload, em.strategy)
	return nil
}

func (em *emitter) ErrorEvent(e event.Event, err error) error {

	if !em.events.Declared(e) {
		return event.UndeclaredEvent(e)
	}

	publishEventError(e, err, em.strategy)
	return nil
}
