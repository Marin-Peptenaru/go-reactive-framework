package pubsub

import (
	"context"
	"fmt"
	"reactive-go/behavior"
	"reactive-go/event"

	"github.com/reactivex/rxgo/v2"
)

var eventSources map[event.Event]rxgo.Observable = make(map[event.Event]rxgo.Observable)
var eventEmmitters map[event.Event]chan<- rxgo.Item = make(map[event.Event]chan<- rxgo.Item)


func newEventSource(event event.Event){
	eventsChannel := make(chan rxgo.Item, 100)
	eventEmmitters[event] = eventsChannel
	eventSources[event] = rxgo.FromChannel(eventsChannel, rxgo.WithPublishStrategy())
	eventSources[event].Connect(context.TODO())
}

func newEventBehaviour(event event.Event, behavior *behavior.Behavior) {
	eventSource := eventSource(event)
	
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
func eventSource(event event.Event) rxgo.Observable {
	if eventSources[event] == nil {
	   newEventSource(event)
	}
	return eventSources[event]
}

func publishEventOccurence(event event.Event, payload interface{}){
	go func() {
		eventEmmitters[event] <- rxgo.Of(payload)
	}()
}

func publishEventError(event event.Event, err error){
	go func() {
		eventEmmitters[event] <- rxgo.Error(err)
	}()
}


