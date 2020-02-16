package search

import (
    "algorithms/datastructures"
)

func BFS(node *datastructures.TreeNode, item int) *datastructures.TreeNode {
    queue := make([]*datastructures.TreeNode, 0)
    queue = append(queue, node)
    for len(queue) > 0 {
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

func TreeBFS(node *datastructures.Node, item int) *datastructures.Node {
    queue := make([]*datastructures.Node, 0)
    queue = append(queue, node)
    for len(queue) > 0 {
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

func GraphBFTraverse(g *datastructures.Graph, start string, f func(*datastructures.Vertex)) {
    queue := make([]*datastructures.Vertex, 0)
    queue = append(queue, g.Vertices[start])
    visited := make(map[*datastructures.Vertex]bool)
    for len(queue) != 0 {
        i := queue[0]
        for _, v := range i.Edges {
            if _, ok := visited[v]; !ok {
                queue = append(queue, v)
                visited[v] = true
            }
        }
        f(i)
        visited[i] = true
    }
}
