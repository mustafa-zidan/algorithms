package search

import (
	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

func BFS(node *ds.TreeNode, item int) *ds.TreeNode {
	queue := ds.NewQueue()
	queue.Push(node)
	for !queue.IsEmpty() {
		i := queue.Pop().(*ds.TreeNode)
		if i.Val == item {
			return i
		}

		if i.Left != nil {
			queue.Push(i.Left)
		}

		if i.Right != nil {
			queue.Push(i.Right)
		}
	}
	return nil
}

func TreeBFS(node *ds.Node, item int) *ds.Node {
	queue := ds.NewQueue()
	queue.Push(node)
	for !queue.IsEmpty() {
		i := queue.Pop().(*ds.Node)
		if i.Val == item {
			return i
		}
		if i.Children != nil {
			for _, child := range i.Children {
				queue.Push(child)
			}
		}
	}
	return nil
}

func GraphBFTraverse(g *ds.Graph, f func(*ds.Vertex)) {
	queue := ds.NewQueue()
	queue.Push(g.Vertices[0])
	visited := make(map[*ds.Vertex]bool)
	for !queue.IsEmpty() {
		i := queue.Pop().(*ds.Vertex)
		visited[i] = true
		if edges, ok := g.Edges[*i]; ok {
			for _, v := range edges {
				if _, ok = visited[v]; !ok {
					queue.Push(v)
					visited[v] = true
				}
			}
		}
		f(i)
	}
}
