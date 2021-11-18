package events

import (
	"context"

	"github.com/reactivex/rxgo/v2"
)

type coreImpl struct {
	eventSources map[Event] rxgo.Observable
	eventEmmiters map[Event] chan<- rxgo.Item
}

var core *coreImpl = &coreImpl{
	eventSources: make(map[Event]rxgo.Observable),
	eventEmmiters: make(map[Event]chan<- rxgo.Item),
}

func eventDrivenApplication() EventDrivenApplicationCore{
	return core
}

func (core *coreImpl) newEventSource(event Event){
	eventsChannel := make(chan rxgo.Item, 100)
	core.eventEmmiters[event] = eventsChannel
	core.eventSources[event] = rxgo.FromChannel(eventsChannel, rxgo.WithPublishStrategy())
	core.eventSources[event].Connect(context.TODO())
}

func (core *coreImpl) NewEventBehaviour(event Event, behaviour *Behaviour) {
	eventSource := core.EventSource(event)
	
	if behaviour.onEvent != nil {eventSource.DoOnNext(behaviour.onEvent)}
	if behaviour.onError != nil {eventSource.DoOnError(behaviour.onError)}
	if behaviour.onDisposed != nil {eventSource.DoOnCompleted(behaviour.onDisposed)}
}
func (core *coreImpl) EventSource(event Event) rxgo.Observable {
	if core.eventSources[event] == nil {
		core.newEventSource(event)
	}
	return core.eventSources[event]
}

func (core *coreImpl) publishEventOccurence(event Event, payload interface{}){
	go func() {
		core.eventEmmiters[event] <- rxgo.Of(payload)
	}()
}

func (core *coreImpl) publishEventError(event Event, err error){
	go func() {
		core.eventEmmiters[event] <- rxgo.Error(err)
	}()
}


