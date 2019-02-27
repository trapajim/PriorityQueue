package PriorityQueue

import (
	"container/heap"
	"errors"
	"math"
	"sync"
)

type PriorityQueue struct {
	q *Queue
	sync.Mutex
	autoLock bool
	count    int
}

// New generates a new PriorityQueue
func New(autoMutexLock bool) PriorityQueue {
	pq := make(Queue, 0)
	heap.Init(&pq)
	return PriorityQueue{q: &pq, autoLock: autoMutexLock}
}

// Push add a new element with a priority (0 lowest)
func (pq *PriorityQueue) Push(value interface{}, priority uint8) (bool, error) {
	if pq.autoLock {
		pq.Lock()
		defer pq.Unlock()
	}

	if pq.q.Len() == math.MaxUint8 {
		return false, errors.New("Queue is full")
	}
	pq.count++
	item := Item{
		index:    pq.count,
		priority: priority,
		value:    value,
	}
	heap.Push(pq.q, &item)
	return true, nil
}

// Pull returns the highest priority value
func (pq *PriorityQueue) Pull() (value interface{}, err error) {
	if pq.Len() == 0 {
		return nil, errors.New("Queue is empty")
	}
	if pq.autoLock {
		pq.Lock()
		defer pq.Unlock()
	}
	return heap.Pop(pq.q).(*Item).value, nil
}

func (pq *PriorityQueue) Len() int {
	if pq.autoLock {
		pq.Lock()
		defer pq.Unlock()
	}
	return pq.count
}
