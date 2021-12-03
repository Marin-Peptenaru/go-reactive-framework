package event

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventSetCreation(t *testing.T) {
	declaredEvents := []Event{Event("test1"), Event("test2"), Event("test3")}
	undeclaredEvent := Event("undeclared-event")

	eventSet := NewEventSet(declaredEvents...)

	for _, event := range declaredEvents {
		assert.True(t, eventSet.Declared(event), fmt.Sprintf("%s should be declared in the event set.\n", event))
	}

	assert.False(t, eventSet.Declared(undeclaredEvent), fmt.Sprintf("%s should not be declared in the event set.\n", undeclaredEvent))

}