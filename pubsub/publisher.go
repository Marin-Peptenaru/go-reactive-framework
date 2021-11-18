package events

import "reactive-go/event"

type Publisher interface {
	ValueEvent(e event.Event, payload interface{})
	ErrorEvent(e event.Event, err error)
}

type publisherImpl struct{}

func NewPublisher(publishedEvents ...event.Event) Publisher {
	return &publisherImpl{}
}

func (publisher *publisherImpl) ValueEvent(e event.Event, payload interface{}) {
	publishEventOccurence(e, payload)
}

func (publisher *publisherImpl) ErrorEvent(e event.Event, err error) {
	publishEventError(e, err)
}