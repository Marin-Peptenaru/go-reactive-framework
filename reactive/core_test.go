package reactive

import (
	"reactive-go/event"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type coreTestSuite struct {
	suite.Suite
	testEvent event.Event
}

func (s *coreTestSuite) SetupTest() {
	s.testEvent = event.Event("test")
}

// test the creation of an event source with the Produce propagation strategy.
// the test passes if a both a channel for emitting events and an observable for propagating them
// have been created and stored in the correct map under the correct propagation strategy
func (s *coreTestSuite) Test_newEventSource_ProduceStrategyEvent() {
	newEventSource(s.testEvent, PRODUCE)

	s.NotNil(s.T(), channels[PRODUCE][s.testEvent], "Channel for emitting events shoul be created.\n")
	s.NotNil(s.T(), observables[PRODUCE][s.testEvent], "Observable for propagating events should be created.\n")
}

// test the creation of an event source with the Publish propagation strategy.
// the test passes if a both a channel for emitting events and an observable for propagating them
// have been created and stored in the correct map under the correct propagation strategy
func (s *coreTestSuite) Test_newEventSource_PublishStrategy() {
	newEventSource(s.testEvent, PUBLISH)

	assert.NotNil(s.T(), channels[PUBLISH][s.testEvent], "Channel for emitting events shoul be created.\n")
	assert.NotNil(s.T(), observables[PUBLISH][s.testEvent], "Observable for propagating events should be created.\n")
}

//tests the eventSource function, which returns the observable corresponding to an event for a given propagation strategy.
//this test is meant for the case where the observable for a new event is requested, case in which a new channel and observable should be
// created
func (s *coreTestSuite) Test_eventSource_ProduceStrategy_newEvent() {
	e := event.Event("new-produce-event")
	obs := eventSource(e, PRODUCE)

	s.NotNil(s.T(), obs, "The returned observable should never be nil.\n")
	s.NotNil(s.T(), channels[PRODUCE][e], "Channel for emitting events shoul be created.\n")
	s.NotNil(s.T(), observables[PRODUCE][e], "Observable for propagating events should be created.\n")
	s.True(observables[PRODUCE][e] == obs, "The returned observable should be the same as the newly created one.\n")
}

//tests the eventSource function, which returns the observable corresponding to an event for a given propagation strategy.
//this test is meant for the case where the observable for a new event is requested, case in which a new channel and observable should be
// created
func (s *coreTestSuite) Test_eventSource_PublishStrategy_newEvent() {
	e := event.Event("new-produce-event")
	obs := eventSource(e, PUBLISH)

	s.NotNil(s.T(), obs, "The returned observable should never be nil.\n")
	s.NotNil(s.T(), channels[PUBLISH][e], "Channel for emitting events shoul be created.\n")
	s.NotNil(s.T(), observables[PUBLISH][e], "Observable for propagating events should be created.\n")
	s.True(observables[PUBLISH][e] == obs, "The returned observable should be the same as the newly created one.\n")
}

func (s *coreTestSuite) Test_emitEvent_EventReceived_ProduceStrategy() {
	emittedValue := 42
	publishEventOccurence(s.testEvent, emittedValue, PRODUCE)
	obs := observables[PRODUCE][s.testEvent].Observe()
	timeOut :=  time.Tick(30 * time.Second)
	
	select {
	case receivedValue := <-obs:
		s.NotNil(receivedValue.V)
		s.Nil(receivedValue.E)
		s.Equal(emittedValue, receivedValue.V.(int), "The received value should be the same received value.\n")
	case <-timeOut:
		s.Fail("timed out: value not received")
	}

}
func TestCoreSuite(t *testing.T) {
	suite.Run(t, new(coreTestSuite))
}
