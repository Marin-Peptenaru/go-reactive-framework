package pubsub

import (
	"context"
	"fmt"
	"reactive-go/behavior"
	"reactive-go/event"

	"github.com/reactivex/rxgo/v2"
)

type PropagationStrategy int

// Enum with values used to differentiate between when to create producer observable and when to create publisher observables
const (
	PUBLISH PropagationStrategy = iota
	PRODUCE
)

// maps of observables and channels for emitting and propagating published events
var publishers map[event.Event]rxgo.Observable = make(map[event.Event]rxgo.Observable)
var publishChannels map[event.Event]chan<- rxgo.Item = make(map[event.Event]chan<- rxgo.Item)


// maps of observable and channels for emitting and propagating produced events
var producers map[event.Event]rxgo.Observable = make(map[event.Event]rxgo.Observable)
var produceChannels map[event.Event]chan<- rxgo.Item = make(map[event.Event]chan<- rxgo.Item)


var observables = map[PropagationStrategy]map[event.Event]rxgo.Observable{
	PUBLISH: publishers,
	PRODUCE: producers,
}

var channels = map[PropagationStrategy]map[event.Event]chan<- rxgo.Item{
	PUBLISH: publishChannels,
	PRODUCE: produceChannels,
}

func newEventSource(event event.Event, strategy PropagationStrategy){

	eventsChannel := make(chan rxgo.Item, 100)
	if strategy == PUBLISH {
		publishChannels[event] = eventsChannel
		publishers[event] = rxgo.FromChannel(eventsChannel, rxgo.WithPublishStrategy())
		publishers[event].Connect(context.TODO())
	} else if strategy == PRODUCE {
		produceChannels[event] = eventsChannel
		producers[event] = rxgo.FromChannel(eventsChannel)
	}
}

func newEventBehaviour(event event.Event, behavior *behavior.Behavior, strategy PropagationStrategy) {
	eventSource := eventSource(event, strategy)
	
	if behavior.OnEvent != nil {
		fmt.Println(eventSource != nil)
		eventSource.DoOnNext(behavior.OnEvent)
	}

	if behavior.OnError != nil {
		eventSource.DoOnError(behavior.OnError)
	}

	if behavior.OnDisposed != nil {
		eventSource.DoOnCompleted(behavior.OnDisposed)
	}
}

func eventSource(event event.Event, strategy PropagationStrategy) rxgo.Observable {
	observable := observables[strategy][event]

	if observable == nil {
	   newEventSource(event, strategy)
	}
	return observables[strategy][event]
}

func publishEventOccurence(event event.Event, payload interface{}, strategy PropagationStrategy){
	go func() {
		channels[strategy][event] <- rxgo.Of(payload)
	}()
}

func publishEventError(event event.Event, err error, strategy PropagationStrategy){
	go func() {
		channels[strategy][event] <- rxgo.Error(err)
	}()
}


