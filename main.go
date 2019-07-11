package main

import (
	"fmt"

	ds "github.com/mustafa-zidan/algorithms/datastructures"
	"github.com/mustafa-zidan/algorithms/search"
)

func main() {
	// stack()
	// queue()
	// binarySearchTree()
	// avlTree()
	// graph()
	weightedGraph()
}


func avlTree() {
	var root *ds.AVLNode
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	for _, key := range keys {
		root = ds.Insert(root, key)
	}

	ds.AVLInOrder(root, func(node *ds.AVLNode) {
		fmt.Println(node.Val, node.Height)
	})
	ds.AVLPostOrder(root, func(node *ds.AVLNode) {
		fmt.Println(node.Val, node.Height)
	})
	ds.AVLPostOrder(root, func(node *ds.AVLNode) {
		fmt.Println(node.Val, node.Height)
	})
}

func binarySearchTree() {
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	tree := &ds.BinarySearchTree{}
	for _, key := range keys {
		tree.Insert(key)
	}
	fmt.Println(tree.InOrderTreeWalk())
}

func graph() {
	g := &ds.Graph{}
	g.AddEdge("A", "B", true)
	g.AddEdge("A", "C", true)
	g.AddEdge("B", "E", true)
	g.AddEdge("C", "E", true)
	g.AddEdge("E", "F", true)
	g.AddEdge("D", "A", true)

	g.String()
	search.GraphBFTraverse(g, "A", func(v *ds.Vertex) {
		fmt.Printf("%v\n", v)
	})

	search.GraphDFTraverse(g, "A", func(v *ds.Vertex) {
		fmt.Printf("%v\n", v)
	})
	// WARNING STACKOVERFLOW
	// search.GraphDFTraverseRec(g, g.Vertices[0], nil, func(v *ds.Vertex) {
	// 	fmt.Printf("%v\n", v)
	// })
}


func stack() {
	stack := ds.NewStack()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	for stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}

func queue() {
	queue := ds.NewQueue()
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	for queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}
}


func weightedGraph() {
	g := &ds.WeightedGraph{}

	g.AddEdge("S", "A", 7, true)
	g.AddEdge("S", "B", 2, true)
	g.AddEdge("S", "C", 3, true)

	g.AddEdge("A", "B", 3, true)
	g.AddEdge("A", "D", 4, true)

	g.AddEdge("B", "D", 4, true)
	g.AddEdge("B", "H", 1, true)

	g.AddEdge("C", "L", 2, true)

	g.AddEdge("D", "F", 5, true)

	g.AddEdge("E", "G", 2, true)
	g.AddEdge("E", "K", 5, true)

	g.AddEdge("F", "H", 3, true)

	g.AddEdge("G", "H", 2, true)

	g.AddEdge("I", "J", 6, true)
	g.AddEdge("I", "K", 4, true)
	g.AddEdge("I", "L", 4, true)

	g.AddEdge("J", "K", 4, true)
	g.AddEdge("J", "L", 4, true)

	fmt.Println(search.Dijkstra(g, "S", "E"))
}