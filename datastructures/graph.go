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

type WeightedGraph struct {
	Vertices 	[]*Vertex
	Edges		[]*Edge
	Lock 		sync.RWMutex
}

type Edge struct{
	V1,V2 *Vertex
	Weight int
}

func (w *WeightedGraph) AddVertex(v * Vertex) {
	if w.Vertices == nil {
		w.Vertices = make([]*Vertex, 0)
	}
	w.Vertices = append(w.Vertices, v)
}

func (w *WeightedGraph) AddEdge(v1, v2 *Vertex, weight int) {
	if w.Edges == nil {
		w.Edges = make([]*Edge, 0)
	}
	// search if edge already exists then just modify weight
	for _, e := range w.Edges {
		if e.V1 == v1 && e.V2 == v2 {
			e.Weight = weight
			return
		}
	}
	w.Edges = append(w.Edges, &Edge{v1,v2, weight})
}

func (w *WeightedGraph) GetEdges(v *Vertex) []*Edge{
	edges := make([]*Edge, 0)
	for _, e := range w.Edges {
		if e.V1 == v {
			edges = append(edges, e)
		}
	}
	return edges
}



func (w *WeightedGraph) Dijkstra(src, dest *Vertex) *Edge {
	path := make(map[*Vertex]int)
	visited := make(map[*Vertex]bool)
	for _, v := range w.Vertices {
		weight := path[v]
		for _,e := range w.GetEdges(v) {
			if _, ok := visited[e.V2]; ok {
				continue
			}
			edgePath := weight + e.Weight
			if  path[e.V2] == 0 || edgePath < path[e.V2] {
				path[e.V2] = edgePath
			}
		}
	}
	return &Edge{src, dest, path[dest]}
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
