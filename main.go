package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int){
	i := 0
	for {
		ch <- i
		time.Sleep(1 * time.Second)
		i++
	}
}

func consumer(name string, ch <-chan int){
	for{
		fmt.Printf("%s : %d\n", name, <-ch)
	}
}





func main() {

	ch := make(chan int)
	wait := make(chan interface{})

	go producer(ch)
	go consumer("Consumer 1", ch)
	go consumer("Consumer 2", ch)

	<-wait

}