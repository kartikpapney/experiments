package main

import (
	"fmt"
	"sync"
)

func main() {
	q, err := New[int](100000)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {

				err := q.Enqueue(j*1000 + i)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	wg.Wait()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {

				err := q.Enqueue(j*1000 + i)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
}
