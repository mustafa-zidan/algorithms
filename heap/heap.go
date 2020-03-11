package heap

import (
    "math"
)

type Heap []int

func NewMaxHeap(items ...int) *Heap {
    h := Heap{}
    h = append(h, items...)
    for i := (len(h) / 2); i >= 0; i-- {
        h.MaxHeapify(i)
    }
    return &h
}

func (h Heap) MaxHeapify(index int) {
    left, right := h.LeftChild(index), h.RightChild(index)

    if left != -1 && right != -1 {
        child := left
        if h[left] < h[right] {
            child = right
        }
        if h[child] > h[index] {
            h[index], h[child] = h[child], h[index]
        }
        h.MaxHeapify(child)
    } else if left != -1 && h[left] > h[index] {
        h[index], h[left] = h[left], h[index]
        h.MaxHeapify(left)
    } else if right != -1 && h[right] > h[index] {
        h[index], h[right] = h[right], h[index]
        h.MaxHeapify(right)
    }

}

func (h Heap) LeftChild(index int) (next int) {
    if 2*index >= len(h) {
        return -1
    }
    return 2 * index
}

func (h Heap) RightChild(index int) (next int) {
    if 2*index+1 >= len(h) {
        return -1
    }
    return 2*index + 1
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
