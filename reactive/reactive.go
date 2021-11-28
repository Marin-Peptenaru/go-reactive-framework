package reactive

import "reactive-go/event"

type reactive struct {
	events   event.EventSet
	strategy PropagationStrategy
}


