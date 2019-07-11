package datastructures

import (
	"fmt"
	"sync"
)

type WeightedVertex struct {
	Val string
	Edges map[string]*WeightedEdge
}

type WeightedEdge struct {
	Vertex *WeightedVertex
	Weight int
}

type WeightedGraph struct {
	Vertices map[string]*WeightedVertex
	lock sync.RWMutex
}


type Vertex struct {
	Val string
	Edges map[string]*Vertex
}

type Graph struct {
	Vertices map[string]*Vertex
	lock sync.RWMutex
}


type Path struct {
	ID string
	Vertices []*WeightedVertex
	Weight int
}

type Paths  []*Path

func (p Paths) Less(i,j int) bool { return p[i].Weight < p[j].Weight}
func (p Paths) Swap(i,j int)  { p[i],p[j] = p[j],p[i]}
func (p Paths) Len() int { return len(p)}

func (g *WeightedGraph) AddVertex(value string ) *WeightedVertex{
	if g.Vertices == nil {
		g.Vertices = make(map[string]*WeightedVertex, 0)
	}
	g.Vertices[value] = &WeightedVertex{value, make(map[string]*WeightedEdge)}
	return g.Vertices[value]
}

func (g *WeightedGraph) AddEdge(a,b string,weight int,  bidirictional bool) {
	g.lock.Lock()
	if _,ok := g.Vertices[a]; !ok {
		g.AddVertex(a)
	}

	if _,ok := g.Vertices[b]; !ok {
		g.AddVertex(b)
	}

	A := g.Vertices[a]
	B := g.Vertices[b]
	if _, ok := A.Edges[b]; ok {
		A.Edges[b] = &WeightedEdge{B, weight}
	}

	if _, ok := B.Edges[a]; ok && bidirictional{
		B.Edges[a] = &WeightedEdge{A, weight}
	}

	g.lock.Unlock()
}

func (n *WeightedVertex) String() string {
	return fmt.Sprintf("%v", n.Val)
}

// String Prints String Presentation of the Graph
func (g *WeightedGraph) String() {
	g.lock.RLock()
	s := ""
	for _, v := range g.Vertices{
		s += v.String() + " -> "
		for _, e := range v.Edges{
			s += e.Vertex.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
// String Prints String Presentation of the Graph
func (g *Graph) String() {
	g.lock.RLock()
	s := ""
	for _, v := range g.Vertices{
		s += v.String() + " -> "
		for _, e := range v.Edges{
			s += e.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

func (n *Vertex) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (g *Graph) AddVertex(value string ) *Vertex{
	g.lock.Lock()
	if g.Vertices == nil {
		g.Vertices = make(map[string]*Vertex, 0)
	}
	g.Vertices[value] = &Vertex{value, make(map[string]*Vertex, 0)}
	g.lock.Unlock()
	return g.Vertices[value]
}

func (g *Graph) AddEdge(a,b string, bidirictional bool) {
	g.lock.Lock()
	if _,ok := g.Vertices[a]; !ok {
		g.AddVertex(a)
	}

	if _,ok := g.Vertices[b]; !ok {
		g.AddVertex(b)
	}

	A := g.Vertices[a]
	B := g.Vertices[b]
	if _, ok := A.Edges[b]; ok {
		A.Edges[b] = B
	}

	if _, ok := B.Edges[a]; ok && bidirictional{
		B.Edges[a] = A
	}

	g.lock.Unlock()
}


