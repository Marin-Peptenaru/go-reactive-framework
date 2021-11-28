package playground

import "reactive-go/event"

var TrivialEvent = event.Event("Trivial")
var TrivialEventSet = event.NewEventSet(TrivialEvent)
