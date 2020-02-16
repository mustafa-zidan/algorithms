package main

import (
    "fmt"

    "algorithms/datastructures"
    "algorithms/search"
    "algorithms/sort"
)

func main() {
    sorts()
    stack()
    queue()
    binarySearchTree()
    avlTree()
    // graph()
    weightedGraph()
}

func sorts() {
    arr := []int{11, 16, 2, 8, 1, 9, 4, 7}
    fmt.Printf("Quick Sort Input: %v\n", arr)
    sort.QuickSort(arr)
    fmt.Printf("Quick Sort Output: %v\n", arr)

    arr = []int{11, 16, 2, 8, 1, 9, 4, 7}
    fmt.Printf("Merge Sort Input: %v\n", arr)
    arr = sort.MergeSort(arr)
    fmt.Printf("Merge Sort Output: %v\n", arr)

    arr = []int{11, 16, 2, 8, 1, 9, 4, 7}
    fmt.Printf("Insertion Sort Input: %v\n", arr)
    sort.InsertionSort(arr)
    fmt.Printf("Insertion Sort Output: %v\n", arr)
}

func avlTree() {
    var root *datastructures.AVLNode
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    for _, key := range keys {
        root = datastructures.Insert(root, key)
    }

    datastructures.AVLInOrder(root, func(node *datastructures.AVLNode) {
        fmt.Printf("AVL In Order Node Value %v, Height: %v \n", node.Val, node.Height)
    })
    datastructures.AVLPreOrder(root, func(node *datastructures.AVLNode) {
        fmt.Printf("AVL Pre Order Node Value %v, Height: %v \n", node.Val, node.Height)
    })
    datastructures.AVLPostOrder(root, func(node *datastructures.AVLNode) {
        fmt.Printf("AVL Post Order Node Value %v, Height: %v \n", node.Val, node.Height)
    })
}

func binarySearchTree() {
    keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
    tree := &datastructures.BinarySearchTree{}
    for _, key := range keys {
        tree.Insert(key)
    }
    fmt.Printf("Binary Search Tree In Order Walk: %v\n", tree.InOrderTreeWalk())
}

func graph() {
    g := &datastructures.Graph{}
    g.AddEdge("A", "B", true)
    g.AddEdge("A", "C", true)
    g.AddEdge("B", "E", true)
    g.AddEdge("C", "E", true)
    g.AddEdge("E", "F", true)
    g.AddEdge("D", "A", true)

    fmt.Println("Graph:")
    g.String()
    search.GraphBFTraverse(g, "A", func(v *datastructures.Vertex) {
        fmt.Printf("Graph BFT raverse: %v\n", v)
    })

    search.GraphDFTraverse(g, "A", func(v *datastructures.Vertex) {
        fmt.Printf("Graph DFT raverse: %v\n", v)
    })
    // WARNING STACKOVERFLOW
    // search.GraphDFTraverseRec(g, g.Vertices[0], nil, func(v *datastructures.Vertex) {
    //     fmt.Printf("Graph DFT raverseRec %v\n", v)
    // })
}

func stack() {
    arr := []int{2, 3, 4, 5, 6}
    stack := datastructures.NewStack()

    for _, v := range arr {
        fmt.Printf("Stack Push %v\n", v)
        stack.Push(v)
    }

    for !stack.IsEmpty() {
        fmt.Printf("Stack pop %v\n", stack.Pop())
    }
}

func queue() {
    arr := []int{2, 3, 4, 5, 6}
    queue := datastructures.NewQueue()

    for _, v := range arr {
        fmt.Printf("Queue Push %v\n", v)
        queue.Push(v)
    }

    for !queue.IsEmpty() {
        fmt.Printf("Queue pop %v\n", queue.Pop())
    }
}

func weightedGraph() {
    g := &datastructures.WeightedGraph{}

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

    fmt.Printf("Dijkstra WeightedGraph search S -> E %v \n", search.Dijkstra(g, "S", "E"))
}
