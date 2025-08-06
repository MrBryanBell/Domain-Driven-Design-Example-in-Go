package Counter

import "sync"

type Counter struct {
	mutex sync.Mutex
	count int
}

func New() Counter {
	return Counter{
		count: 0,
		mutex: sync.Mutex{},
	}
}

func (counter *Counter) Increase() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	counter.count += 1
}

func (counter *Counter) Count() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	return counter.count
}
