package Timer

import "time"

func After(duration time.Duration) <-chan struct{} {
	var channel = make(chan struct{})

	go func() {
		time.Sleep(duration)
		channel <- struct{}{}
	}()

	return channel
}

/*
**Firma de la función:
Timer.Tick(2 * Time.Seconds) <-chan struct{}

**Usage Example:
package main

import (

	"fmt"
	"mymodule/experiments/concurrency/Timer"
	"time"

)

	func main() {
		for {
			<-Timer.Tick(2 * time.Second)
			fmt.Println("Another 2s has passed")
		}
	}
*/
func Tick(duration time.Duration) <-chan struct{} {
	var channel = make(chan struct{})

	go func() {
		for {
			time.Sleep(duration)
			channel <- struct{}{}
		}
	}()

	return channel
}
