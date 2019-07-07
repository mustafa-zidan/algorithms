package main

import (
	"fmt"

	ds "mustafa-zidan/algorithms/datastructures"
 )


func main() {
	queue := ds.NewQueue()
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())


	stack := ds.NewStack()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())


}