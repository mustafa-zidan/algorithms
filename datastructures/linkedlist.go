package datastructures

type LinkedList struct {
    Previous *LinkedList
    Val      interface{}
    Next     *LinkedList
}

func (l *LinkedList) Add(item interface{}) {
    next := l.Next
    if next != nil {
        for next.Next != nil {
            next = next.Next
        }
    }
    next = &LinkedList{l, item, nil}
    l.Next = next
}

func (l *LinkedList) Remove(i int) {
    node := l.Get(i)
    if node.Next != nil && node.Previous != nil {
        node.Previous.Next = node.Next
        node.Next.Previous = node.Previous
    } else if node.Next != nil && node.Previous == nil {
        node.Next.Previous = nil
    } else if node.Previous != nil && node.Next == nil {
        node.Previous.Next = nil
    }

}

func (l *LinkedList) Get(i int) *LinkedList {
    count := 1
    item := l.Next
    for count <= i {
        item = item.Next
        count++
    }
    return item
}

func (l *LinkedList) Len() int {
    count := 1
    next := l.Next
    for next != nil {
        next = next.Next
        count++
    }
    return count
}
