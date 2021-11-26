package main

import (
	"reactive-go/playground"
)

func main() {
	wait := make(chan int)
	pub := playground.NewTrivialProducer()
	sub1 := playground.NewTrivialConsumer("Con 1")
	sub2 := playground.NewTrivialConsumer("Con 2")

	pub.StartProducing()
	sub1.StartConsuming()
	sub2.StartConsuming()
	<-wait
}