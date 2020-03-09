package heap

import (
    "math"
)

type Heap []int

func NewMaxHeap(items ...int) *Heap {
    h := Heap{}
    h = append(h, items...)
    for i := ((len(h) / 2) - 1); i >= 0; i-- {
        h.MaxHeapify(i)
    }
    return &h
}

func (h Heap) MaxHeapify(index int) {
    li, left := h.LeftChild(index)
    ri, right := h.RightChild(index)
    child := li
    if left < right {
        child = ri
    }
    if h[child] > h[index] {
        h[index], h[child] = h[child], h[index]
    }
}

func (h Heap) LeftChild(index int) (next int, value int) {
    if 2*index > len(h) {
        return -1, 0
    }
    return 2 * index, h[2*index]
}

func (h Heap) RightChild(index int) (next int, value int) {
    if 2*index+1 > len(h) {
        return -1, 0
    }
    return 2*index + 1, h[2*index+1]
}

func (h Heap) Size() int {
    return len(h)
}

func (h Heap) Depth() int {
    return int(math.Floor(math.Log2(float64(len(h) + 1))))
}

func (h *Heap) Height() int {
    return 0
}

func (h *Heap) Width() {

}

func (h *Heap) Pop() int {
    return 0
}

func (h *Heap) Push(item int) {

}
