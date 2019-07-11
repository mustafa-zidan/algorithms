package search

import (
	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

func BFS(node *ds.TreeNode, item int) *ds.TreeNode {
	queue := make([]*ds.TreeNode, 0)
	queue = append(queue, node)
	for len(queue)  > 0 {
		i := queue[0]
		if i.Val == item {
			return i
		}

		if i.Left != nil {
			queue = append(queue, i.Left)
		}

		if i.Right != nil {
			queue = append(queue, i.Right)
		}
	}
	return nil
}

func TreeBFS(node *ds.Node, item int) *ds.Node {
	queue := make([]*ds.Node, 0)
	queue = append(queue, node)
	for len(queue)  > 0 {
		i := queue[0]
		if i.Val == item {
			return i
		}
		if i.Children != nil {
			queue = append(queue, i.Children...)
		}
	}
	return nil
}

func GraphBFTraverse(g *ds.Graph, start string, f func(*ds.Vertex)) {
	queue := make([]*ds.Vertex, 0)
	queue = append(queue,g.Vertices[start])
	visited := make(map[*ds.Vertex]bool)
	for len(queue) != 0 {
		i := queue[0]
		for _, v := range i.Edges {
			if _, ok := visited[v]; !ok {
				queue= append(queue, v )
				visited[v] = true
			}
		}
		f(i)
		visited[i] = true
	}
}
