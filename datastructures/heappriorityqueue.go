package datastructures

import (
    "container/heap"
)

type PriorityQueueItem struct {
    Val   interface{}
    Weight int
    Index int // The index of the item in the heap.
}

type HeapPriorityQueue []*PriorityQueueItem


func NewHeapPriorityQueue()  *HeapPriorityQueue{
	pq := make(HeapPriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func (pq HeapPriorityQueue) Len() int { return len(pq) }

func (pq HeapPriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the lowest based on Weight number as the priority
    // The lower the Weight, the higher the priority
    return pq[i].Weight < pq[j].Weight
}

// We just implement the pre-defined function in interface of heap.

func (pq *HeapPriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.Index = -1
    *pq = old[0 : n-1]
    return item
}

func (pq *HeapPriorityQueue) Push(i interface{}) {
    n := len(*pq)
    item := i.(*PriorityQueueItem)
    item.Index = n
    *pq = append(*pq, item)
}

func (pq HeapPriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].Index = i
    pq[j].Index = j
}

func (pq HeapPriorityQueue) IsEmpty() bool {
    return len(pq) == 0
}

func (pq HeapPriorityQueue) Next() interface{} {
	i := heap.Pop(&pq).(*PriorityQueueItem)
    return i.Val
}

