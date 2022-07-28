// DATE: 2022/5/11
// DESC: 队列

package structure

import (
	"sync"
)

type ArrayQueue struct {
	arr []interface{}
	sync.RWMutex
}

func NewArrayQueue() *ArrayQueue {
	return new(ArrayQueue)
}

func (q *ArrayQueue) Size() int {
	return len(q.arr)
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ArrayQueue) GetFront() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.arr[0]
}

func (q *ArrayQueue) GetBack() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.arr[q.Size()-1]
}

func (q *ArrayQueue) PopBack() interface{} {
	q.Lock()
	defer q.Unlock()
	if q.IsEmpty() {
		return nil
	}
	ret := q.arr[q.Size()-1]
	q.arr = q.arr[:q.Size()-1]
	return ret
}

func (q *ArrayQueue) DelQueue() interface{} {
	q.Lock()
	defer q.Unlock()
	if len(q.arr) == 0 {
		return nil
	}
	first := q.arr[0]
	q.arr = q.arr[1:]
	return first
}

func (q *ArrayQueue) EnQueue(item interface{}) {
	q.Lock()
	defer q.Unlock()
	q.arr = append(q.arr, item)
}

// ---------------- without lock through like link_list ------------------

type queueItem struct {
	item interface{}
	prev *queueItem
}

type queue struct {
	curr *queueItem
	last *queueItem
	size uint64
}

func NewQueue() *queue {
	return new(queue)
}

// Enqueue puts a given item into Queue
func (q *queue) Enqueue(item interface{}) {
	if q.size == 0 {
		q.curr = &queueItem{item: item, prev: nil}
		q.last = q.curr
		q.size++
		return
	}
	qq := &queueItem{item: item, prev: nil}
	q.last.prev = qq
	q.last = qq
	q.size++
}

// Dequeue Extracts first item from the Queue
func (q *queue) Dequeue() interface{} {
	if q.size > 0 {
		item := q.curr.item
		q.curr = q.curr.prev
		q.size--
		return item
	}
	return nil
}
