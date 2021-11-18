package events

type Publisher interface {
	ValueEvent(e Event, payload interface{})
	ErrorEvent(e Event, err error)
}

type publisherImpl struct{}

func NewPublisher(publishedEvents ...Event) Publisher {
	return &publisherImpl{}
}

func (publisher *publisherImpl) ValueEvent(e Event, payload interface{}) {
	eventDrivenApplication().publishEventOccurence(e, payload)
}

func (publisher *publisherImpl) ErrorEvent(e Event, err error) {
	eventDrivenApplication().publishEventError(e, err)
}