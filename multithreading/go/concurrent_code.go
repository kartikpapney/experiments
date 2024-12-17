package main

import (
	"sync"
)

type NonAtomicInt struct {
	Val int
}

func (val *NonAtomicInt) Increament(threadCount int) {
	var wg sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val.Val++
		}()
	}
	wg.Wait()
}
