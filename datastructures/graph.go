package datastructures

import (
	"fmt"
	"sync"
)

// Vertex a single Vertex that composes the tree
type Vertex struct {
	Value string
}

func (n *Vertex) String() string {
	return fmt.Sprintf("%v", n.Value)
}

// Graph the Items graph
type Graph struct {
	Vertices []*Vertex
	Edges    map[Vertex][]*Vertex
	Lock     sync.RWMutex
}

// AddVertex adds a node to the graph
func (g *Graph) AddVertex(n *Vertex) {
	g.Lock.Lock()
	g.Vertices = append(g.Vertices, n)
	g.Lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Vertex) {
	g.Lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Vertex][]*Vertex)
	}
	g.Edges[*n1] = append(g.Edges[*n1], n2)
	g.Edges[*n2] = append(g.Edges[*n2], n1)
	g.Lock.Unlock()
}

// String Prints String Presentation of the Graph
func (g *Graph) String() {
	g.Lock.RLock()
	s := ""
	for i := 0; i < len(g.Vertices); i++ {
		s += g.Vertices[i].String() + " -> "
		near := g.Edges[*g.Vertices[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.Lock.RUnlock()
}

// BFS
// DFS
// A*
// Dijkstra
