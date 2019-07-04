package main

import (
	"fmt"
	"sync"
)

// Stack
type Stack struct {
	items []interface{}
	lock  sync.RWMutex
}

//New returns new instance pointer of Stack struct
func New() *Stack {
	s := &Stack{}
	s.items = make([]interface{}, 0)
	return s
}

// Push adds an Item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	s.items = append(s.items, item)
	s.lock.Unlock()
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return item
}

func (s *Stack) Len() int {
	return len(s.items)
}

func main() {
	stack := New()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}