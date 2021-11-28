package event

import "fmt"

type undeclaredEventError struct {
	event Event
}

func (ue undeclaredEventError) Error() string {
	return fmt.Sprintf("event not declared: %v", ue.event)
}

func UndeclaredEvent(e Event) undeclaredEventError {
	return undeclaredEventError{event: e}
}