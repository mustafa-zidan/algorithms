package datastructures

import (
    "sort"
    "sync"
)

type PriorityQueue struct {
    items Paths
    lock  sync.RWMutex
}

func NewPriorityQueue() *PriorityQueue {
    q := &PriorityQueue{}
    q.items = make(Paths, 0)
    return q
}

func (q *PriorityQueue) Push(item *Path) {
    q.lock.Lock()
    if q.items == nil {
        q.items = make(Paths, 0)
    }
    q.items = append(q.items, item)
    sort.Sort(q.items)
    q.lock.Unlock()
}

func (q *PriorityQueue) Pop() *Path {
    q.lock.Lock()
    if len(q.items) == 0 {
        return nil
    }

    item := q.items[0]
    q.items = q.items[1:]
    q.lock.Unlock()
    return item
}

func (q *PriorityQueue) IsEmpty() bool {
    return q.items.Len() == 0
}
