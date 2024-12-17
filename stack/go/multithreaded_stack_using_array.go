package main

import (
	"fmt"
	"sync"
)

type stack[T any] struct {
	arr      []T
	capacity int
	lock     sync.RWMutex
}

func New[T any](capacity int) (*stack[T], error) {
	if capacity < 1 {
		return nil, fmt.Errorf("invalid capacity")
	}
	return &stack[T]{
		arr:      []T{},
		capacity: capacity,
	}, nil
}

func (st *stack[T]) Top() (*T, error) {
	st.lock.RLock()
	defer st.lock.RUnlock()
	size := st.Size()
	if size == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	return &st.arr[size-1], nil
}

func (st *stack[T]) Pop() error {
	st.lock.Lock()
	defer st.lock.Unlock()
	size := len(st.arr)
	if size == 0 {
		return fmt.Errorf("stack is empty")
	}
	st.arr = st.arr[:size-1]
	return nil
}

func (st *stack[T]) Push(element T) error {
	st.lock.Lock()
	defer st.lock.Unlock()
	if len(st.arr) == st.capacity {
		return fmt.Errorf("stack overflow")
	}
	st.arr = append(st.arr, element)
	return nil
}

func (st *stack[T]) Size() int {
	st.lock.RLock()
	defer st.lock.RUnlock()
	return len(st.arr)
}

func (st *stack[T]) IsEmpty() bool {
	st.lock.RLock()
	defer st.lock.RUnlock()
	return len(st.arr) == 0
}
