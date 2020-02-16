package search

import (
    ds "algorithms/datastructures"
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

func GraphDFTraverse(g *ds.Graph, start string, f func(*ds.Vertex)) {
    stack := ds.NewStack()
    stack.Push(g.Vertices[start])
    visited := make(map[*ds.Vertex]bool)
    for !stack.IsEmpty() {
        i := stack.Pop().(*ds.Vertex)
        for _, v := range i.Edges {
            if _, ok := visited[v]; !ok {
                stack.Push(v)
            }
        }
        visited[i] = true
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
    discovered[v] = true
    for _, adjecent := range v.Edges {
        GraphDFTraverseRec(g, adjecent, discovered, f)
    }

    f(v)

}
