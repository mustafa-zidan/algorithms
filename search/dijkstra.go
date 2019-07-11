package search

import (
	. "github.com/mustafa-zidan/algorithms/datastructures"
)

func Dijkstra(g *WeightedGraph, start, end string) *Path {
	queue := NewPriorityQueue()
	paths := make(map[string]*Path)
	paths[start] = &Path{start, []*WeightedVertex{g.Vertices[start]}, 0 }
	queue.Push(paths[start])
	visited := make(map[string]bool)
	for !queue.IsEmpty() {
		path := queue.Pop()
		if visited[path.ID] {
			continue
		}
		v := g.Vertices[path.ID]
		if path.ID == end {
			return path
		}
		for _, e := range v.Edges{
			if p, ok :=paths[e.Vertex.Val]; !ok || p.Weight > path.Weight +  e.Weight {
				p = &Path{e.Vertex.Val , append(path.Vertices, v), path.Weight +  e.Weight}
				paths[e.Vertex.Val] = p
				queue.Push(p)
			}
		}
		visited[v.Val] = true
	}
	return nil
}
