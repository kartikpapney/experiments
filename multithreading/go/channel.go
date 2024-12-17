package main

import (
	"fmt"
)

func Channel(threadCount int) {
	channel := make(chan int)
	for i := 0; i < threadCount; i++ {
		go func() {
			channel <- i
		}()
	}

	sum := 0
	for i := 0; i < threadCount; i++ {
		data := <-channel
		sum += data
	}

	fmt.Println(sum)
}
