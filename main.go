package main

import (
	"fmt"

	ds "github.com/mustafa-zidan/algorithms/datastructures"
	"github.com/mustafa-zidan/algorithms/search"
)

func main() {
	queue := ds.NewQueue()
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	for queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}

	stack := ds.NewStack()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	for stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}

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

	g := &ds.Graph{}

	fillGraph(g)
	g.String()
	search.GraphBFTraverse(g, func(v *ds.Vertex) {
		fmt.Printf("%v\n", v)
	})

	search.GraphDFTraverse(g, func(v *ds.Vertex) {
		fmt.Printf("%v\n", v)
	})
	// WARNING STACKOVERFLOW
	// search.GraphDFTraverseRec(g, g.Vertices[0], nil, func(v *ds.Vertex) {
	// 	fmt.Printf("%v\n", v)
	// })
}

func fillGraph(g *ds.Graph) {
	nA := ds.Vertex{"A"}
	nB := ds.Vertex{"B"}
	nC := ds.Vertex{"C"}
	nD := ds.Vertex{"D"}
	nE := ds.Vertex{"E"}
	nF := ds.Vertex{"F"}
	g.AddVertex(&nA)
	g.AddVertex(&nB)
	g.AddVertex(&nC)
	g.AddVertex(&nD)
	g.AddVertex(&nE)
	g.AddVertex(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}
