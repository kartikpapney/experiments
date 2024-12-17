package main

import (
	"fmt"
	"sync"
)

func main() {
	st, _ := New[int](1000)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			st.Push(i)
		}()
	}
	wg.Wait()
	sum := 0
	for !st.IsEmpty() {
		val, err := st.Top()
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += *val
		st.Pop()
	}
	fmt.Println(sum)
}
