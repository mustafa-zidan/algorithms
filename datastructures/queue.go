package datastructures

import (
	"sync"
)

type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

//NewQueue returns new instance pointer of Queue struct
func NewQueue() *Queue {
	q := Queue{}
	q.items = make([]interface{}, 0)
	return &q
}

// Push adds an Item to the end of the queue
func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	q.items = append(q.items, item)
	q.lock.Unlock()
}

// Pop removes first item added to the queue
func (q *Queue) Pop() interface{} {
	q.lock.Lock()
	item := q.items[0]
	if len(q.items) > 1 {
		q.items = q.items[1:]
	} else {
		q.items = make([]interface{}, 0)
	}
	q.lock.Unlock()
	return item
}
