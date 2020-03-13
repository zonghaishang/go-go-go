package main

import (
	"fmt"
	"sync"
)

func main() {
	queue := NewSafeQueue(1)

	var signal sync.WaitGroup
	signal.Add(2)

	go func() {
		defer signal.Done()
		for i := 0; i < 10 ; i++ {
			queue.Put(i)
		}
	}()

	go func() {
		defer signal.Done()
		for i := 0; i < 10 ; i++ {
			_ = queue.Take()
		}
	}()

	signal.Wait()
}

type safeQueue struct {
	cond     *sync.Cond
	capacity int
	elements []interface{}
}

func NewSafeQueue(capacity int) *safeQueue {
	return &safeQueue{
		cond:     sync.NewCond(&sync.Mutex{}),
		capacity: capacity,
		elements: make([]interface{}, 0, capacity),
	}
}

func (q *safeQueue) Put(e interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.elements) == q.capacity {
		q.cond.Wait()
	}

	q.elements = append(q.elements, e)

	fmt.Printf("Put %v\n", e)

	// notify to Take
	if len(q.elements) > 0 {
		q.cond.Signal()
	}
}

func (q *safeQueue) Take() interface{} {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.elements) <= 0 {
		q.cond.Wait()
	}

	element := q.elements[0]
	q.elements = q.elements[1:]

	fmt.Printf("Take %v\n", element)

	// notify to put
	if len(q.elements) < q.capacity {
		q.cond.Signal()
	}

	return element
}
