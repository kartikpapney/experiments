package main

import (
	"sync"
)

type AtomicInt struct {
	Val   int
	Mutex sync.Mutex
}

func (val *AtomicInt) Increament(threadCount int) {
	var wg sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val.Mutex.Lock()
			val.Val++
			val.Mutex.Unlock()
		}()
	}
	wg.Wait()
}
