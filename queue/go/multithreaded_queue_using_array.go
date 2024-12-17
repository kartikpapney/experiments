package main

import (
	"fmt"
	"sync"
)

const maxCapacity int = 1e6

type queue[T any] struct {
	arr        []T
	lock       sync.RWMutex
	capacity   int
	frontIndex int
	rearIndex  int
	length     int
}

func New[T any](capacity int) (*queue[T], error) {
	if capacity < 1 {
		return nil, fmt.Errorf("invalid capacity: %d", capacity)
	} else if capacity > maxCapacity {
		return nil, fmt.Errorf("capacity is capped at: %d", maxCapacity)
	}
	return &queue[T]{
		arr:        make([]T, capacity),
		lock:       sync.RWMutex{},
		capacity:   capacity,
		frontIndex: 0,
		rearIndex:  -1,
		length:     0,
	}, nil
}

func (st *queue[T]) Front() (T, error) {
	st.lock.RLock()
	defer st.lock.RUnlock()

	if st.length == 0 {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return st.arr[st.frontIndex], nil
}

func (st *queue[T]) Dequeue() error {
	st.lock.Lock()
	defer st.lock.Unlock()
	if st.length == 0 {
		return fmt.Errorf("queue is empty")
	}
	st.frontIndex = (st.frontIndex + 1) % st.capacity
	st.length--
	return nil
}

func (st *queue[T]) Enqueue(element T) error {
	st.lock.Lock()
	defer st.lock.Unlock()
	if st.length == st.capacity {
		return fmt.Errorf("queue overflow: capacity reached (%d)", st.capacity)
	}
	st.rearIndex = (st.rearIndex + 1) % st.capacity
	st.arr[st.rearIndex] = element
	st.length++
	return nil
}

func (st *queue[T]) Size() int {
	st.lock.RLock()
	defer st.lock.RUnlock()
	return st.length
}

func (st *queue[T]) IsEmpty() bool {
	st.lock.RLock()
	defer st.lock.RUnlock()
	return st.length == 0
}
