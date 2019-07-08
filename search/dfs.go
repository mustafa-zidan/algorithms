package search

import (
	ds "github.com/mustafa-zidan/algorithms/datastructures"
)

func DFS(node *ds.TreeNode, item int) *ds.TreeNode {
	stack := ds.NewStack()
	stack.Push(node)
	for !stack.IsEmpty() {
		i := stack.Pop().(*ds.TreeNode)
		if i.Val == item {
			return i
		}

		if i.Left != nil {
			stack.Push(i.Left)
		}

		if i.Right != nil {
			stack.Push(i.Right)
		}
	}
	return nil
}

// DFS recursion impel

func DFSR(n *ds.TreeNode, item int) *ds.TreeNode {
	if n.Val == item {
		return n
	}
	var result *ds.TreeNode

	if n.Left != nil {
		result = DFSR(n.Left, item)
	}

	if n.Right != nil && result == nil {
		result = DFSR(n.Right, item)
	}

	return result
}

func TreeDFS(node *ds.Node, item int) *ds.Node {
	stack := ds.NewStack()
	stack.Push(node)
	for !stack.IsEmpty() {
		i := stack.Pop().(*ds.Node)
		if i.Val == item {
			return i
		}
		if i.Children != nil {
			for _, child := range i.Children {
				stack.Push(child)
			}
		}
	}
	return nil
}

func GraphDFTraverse(g *ds.Graph, f func(*ds.Vertex)) {
	stack := ds.NewStack()
	stack.Push(g.Vertices[0])
	visited := make(map[*ds.Vertex]bool)
	for !stack.IsEmpty() {
		i := stack.Pop().(*ds.Vertex)
		visited[i] = true
		if edges, ok := g.Edges[*i]; ok {
			for _, v := range edges {
				if _, ok = visited[v]; !ok {
					stack.Push(v)
					visited[v] = true
				}
			}
		}
		f(i)
	}
}

func GraphDFTraverseRec(g *ds.Graph, v *ds.Vertex, discovered map[*ds.Vertex]bool, f func(*ds.Vertex)) {
	if discovered == nil {
		discovered = make(map[*ds.Vertex]bool)
	}
	if _, ok := discovered[v]; ok {
		return
	}
	if edges, ok := g.Edges[*v]; ok {
		for _, adjecent := range edges {
			GraphDFTraverseRec(g, adjecent, discovered, f)
		}
	}
	f(v)
}
