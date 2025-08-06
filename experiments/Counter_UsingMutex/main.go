/*
	---
	ABOUT
	---
	We are spawning 100 goroutines to update the count of a Counter. Each goroutine will try to update the count by 1. This means that the expected count at the endis 100.


	---
	CHALLENGES
	---
	We have two challenges here:
	(1) Waiting for all the goroutines to finish
	(2) Handling race conditions: multiple goroutines trying to update the count at the same time


	---
	SOLUTION
	---
	(1) We can use a WaitGroup to wait for all the goroutines to finish.
	(2) We can use a Mutex to prevent race conditions. The Mutex is inside the Counter struct. We can lock the Mutex when we want to update the count and unlock it when we are done.

*/

package main

import (
	"fmt"
	"mymodule/experiments/Counter_UsingMutex/Counter"
	"sync"
)

func main() {
	var counter = Counter.New()
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counter.Increase()
		}()
	}

	wg.Wait()

	fmt.Println(counter.Count())
}
