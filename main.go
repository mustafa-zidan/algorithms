package main

import (
	"fmt"

	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

// func main() {
// 	queue := ds.NewQueue()
// 	queue.Push(2)
// 	queue.Push(3)
// 	queue.Push(4)
// 	fmt.Println(queue.Pop())
// 	fmt.Println(queue.Pop())
// 	fmt.Println(queue.Pop())

// 	stack := ds.NewStack()
// 	stack.Push(2)
// 	stack.Push(3)
// 	stack.Push(4)
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())
// 	fmt.Println(stack.Pop())

// }

func main() {
	var root *ds.AVLNode
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	for _, key := range keys {
		root = ds.Insert(root, key)
	}

	tree := &ds.BinarySearchTree{}
	for _, key := range keys {
		tree.Insert(key)
	}

	fmt.Println(tree.InOrderTreeWalk())

	ds.InOrder(root, func(node *ds.AVLNode) {
		fmt.Println(node.Val, node.Height)
	})
}
