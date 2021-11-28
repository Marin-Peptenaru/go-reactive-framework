package main

import (
	"reactive-go/playground"
)

func main() {
	wait := make(chan int)
	pub := playground.NewTrivialPublisher()
	sub1 := playground.NewTrivialSubscriber("Con 1")
	sub2 := playground.NewTrivialSubscriber("Con 2")

	pub.StartPublishing()
	sub1.StartSubscribing()
	sub2.StartSubscribing()
	<-wait
}