package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	queue := NewSafeQueue(1)

	var signal sync.WaitGroup
	signal.Add(3)

	go func() {
		defer signal.Done()
		for i := 0; i < 10; i++ {
			queue.Put(i)
			//fmt.Printf("Put %v\n", i)
		}
	}()

	go func() {
		defer signal.Done()
		for i := 10; i < 20; i++ {
			queue.Put(i)
			//fmt.Printf("Put %v\n", i)
		}
	}()

	go func() {
		defer signal.Done()
		for i := 0; i < 20; i++ {
			_ = queue.Take()
			//fmt.Printf("Take %v\n", v)
		}
	}()

	signal.Wait()
}

type safeQueue struct {
	putLock  *sync.Mutex
	takLock  *sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
	len      int32
	capacity int32
	elements []interface{}
}

func NewSafeQueue(capacity int32) *safeQueue {
	putLock := &sync.Mutex{}
	takLock := &sync.Mutex{}

	return &safeQueue{
		putLock:  putLock,
		takLock:  takLock,
		notFull:  sync.NewCond(putLock),
		notEmpty: sync.NewCond(takLock),
		len:      0,
		capacity: capacity,
		elements: make([]interface{}, 0, capacity),
	}
}

func (q *safeQueue) Put(e interface{}) {
	q.putLock.Lock()

	for q.currentLength() == q.capacity {
		q.notFull.Wait()
	}
	c := q.currentLength()

	// not thread safe
	q.elements = append(q.elements, e)

	fmt.Printf("Put %v\n", e)

	len := q.updateLength(1)

	if len < q.capacity {
		q.notFull.Signal()
	}

	q.putLock.Unlock()

	// notify
	if c == 0 {
		q.notEmpty.Signal()
	}
}

func (q *safeQueue) Take() interface{} {
	q.takLock.Lock()

	for q.currentLength() == 0 {
		q.notEmpty.Wait()
	}

	c := q.currentLength()

	element := q.elements[0]
	// not thread safe
	q.elements = q.elements[1:]

	fmt.Printf("Take %v\n", element)

	len := q.updateLength(-1)

	// notify to take
	if len > 1 {
		q.notEmpty.Signal()
	}

	q.takLock.Unlock()

	// notify
	if c == q.capacity {
		q.notFull.Signal()
	}

	return element
}

func (q *safeQueue) updateLength(offset int32) int32 {
	prev := atomic.LoadInt32(&q.len)
	for !atomic.CompareAndSwapInt32(&q.len, prev, prev+offset) {
		prev = atomic.LoadInt32(&q.len)
	}

	return prev + offset
}

func (q *safeQueue) currentLength() int32 {
	return atomic.LoadInt32(&q.len)
}
