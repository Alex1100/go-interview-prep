package graph

import (
	"errors"
	queue "go-interview-prep/data_structures/queues/queue"
	stack "go-interview-prep/data_structures/stacks/stack"
)

type Edge struct {
	Key        string
	EdgeValues map[string]int
}

type Vertex struct {
	Key   string
	Edges []*Edge
}

type Graph struct {
	Vertexes  map[string]*Vertex
	NodeArray []string
}

func InitGraph() *Graph {
	return &Graph{
		Vertexes:  make(map[string]*Vertex),
		NodeArray: make([]string, 0),
	}
}

func (g *Graph) AddVertex(node_val string) bool {
	if g.Vertexes[node_val] != nil {
		return false
	}

	g.Vertexes[node_val] = &Vertex{Key: node_val, Edges: make([]*Edge, 0)}
	g.NodeArray = append(g.NodeArray, node_val)
	return true
}

func (g *Graph) AddEdge(from, to string) bool {
	if !g.HasVertexes(to, from) || g.HasEdge(from, to) || g.SameEdge(from, to) {
		return false
	}
	var to_edge_vals map[string]int
	var from_edge_vals map[string]int

	to_edge_vals = make(map[string]int)
	to_edge_vals["weight"] = 0
	to_edge_vals["direction"] = 1

	from_edge_vals = make(map[string]int)
	from_edge_vals["weight"] = 0
	from_edge_vals["direction"] = -1

	g.Vertexes[to].Edges = append(g.Vertexes[to].Edges, &Edge{Key: from, EdgeValues: from_edge_vals})
	g.Vertexes[from].Edges = append(g.Vertexes[from].Edges, &Edge{Key: to, EdgeValues: to_edge_vals})
	return true
}

func (g *Graph) HasEdge(from, to string) bool {

	if g.Vertexes[from] != nil {
		for i := 0; i < len(g.Vertexes[from].Edges); i++ {
			if g.Vertexes[from].Edges[i].Key == to {
				return true
			}
		}
	}

	return false
}

func (g *Graph) HasEdges(from, to string) bool {
	return g.HasEdge(from, to) && g.HasEdge(to, from)
}

func (g *Graph) SameEdge(from, to string) bool {
	return from == to
}

func (g *Graph) HasVertexes(from, to string) bool {
	return g.Vertexes[from] != nil && g.Vertexes[to] != nil
}

func (g *Graph) AddEdges(from, to string) bool {
	if !g.HasVertexes(to, from) || g.HasEdge(from, to) || g.SameEdge(from, to) {
		return false
	}
	var to_edge_vals map[string]int
	var from_edge_vals map[string]int

	to_edge_vals = make(map[string]int)
	to_edge_vals["weight"] = 0
	to_edge_vals["direction"] = 1

	from_edge_vals = make(map[string]int)
	from_edge_vals["weight"] = 0
	from_edge_vals["direction"] = 1

	g.Vertexes[to].Edges = append(g.Vertexes[to].Edges, &Edge{Key: from, EdgeValues: from_edge_vals})
	g.Vertexes[from].Edges = append(g.Vertexes[from].Edges, &Edge{Key: to, EdgeValues: to_edge_vals})
	return true
}

func (g *Graph) RemoveVertex(vertex_key string) (*Vertex, error) {
	if g.Vertexes[vertex_key] == nil {
		return g.Vertexes[g.NodeArray[0]], errors.New("Vertex does not exist in Grap")
	}

	removed := g.Vertexes[vertex_key]
	nodes := g.NodeArray[:0]

	for i := 0; i < len(g.NodeArray); i++ {
		if len(g.Vertexes[g.NodeArray[i]].Edges) > 0 {
			current_vertex := g.Vertexes[g.NodeArray[i]].Edges[:0]

			for _, edge := range g.Vertexes[g.NodeArray[i]].Edges {
				if edge.Key != vertex_key {
					current_vertex = append(current_vertex, edge)
				}
			}
			g.Vertexes[g.NodeArray[i]].Edges = current_vertex
		}

		if g.NodeArray[i] == vertex_key {
			delete(g.Vertexes, vertex_key)
		} else {
			nodes = append(nodes, g.NodeArray[i])
		}
	}
	g.NodeArray = nodes
	return removed, nil
}

func (g *Graph) RemoveEdge(from, to string) (bool, error) {
	if !g.HasVertexes(to, from) || !g.HasEdges(from, to) || g.SameEdge(from, to) {
		return false, errors.New("One of more of the given edges don't exist")
	}

	var to_edge_idx int
	var to_edge_val int
	var from_edge_idx int
	var from_edge_val int

	for i := 0; i < len(g.Vertexes[from].Edges); i++ {
		if g.Vertexes[from].Edges[i].Key == to {
			to_edge_idx = i
			to_edge_val = g.Vertexes[from].Edges[to_edge_idx].EdgeValues["direction"]
		}
	}

	for j := 0; j < len(g.Vertexes[to].Edges); j++ {
		if g.Vertexes[to].Edges[j].Key == from {
			from_edge_idx = j
			from_edge_val = g.Vertexes[to].Edges[from_edge_idx].EdgeValues["direction"]
		}
	}

	if to_edge_val > 0 && from_edge_val > 0 {
		g.Vertexes[from].Edges[to_edge_idx].EdgeValues["direction"] = -1
	} else {
		_, err := g.RemoveEdges(from, to)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (g *Graph) RemoveEdges(from, to string) (map[string]*Edge, error) {
	removed_edges := make(map[string]*Edge, 0)

	if !g.HasVertexes(to, from) || !g.HasEdges(from, to) || g.SameEdge(from, to) {
		return removed_edges, errors.New("One of more of the given edges don't exist")
	}

	for i := 0; i < len(g.NodeArray); i++ {
		if len(g.Vertexes[g.NodeArray[i]].Edges) > 0 {
			current_vertex := g.Vertexes[g.NodeArray[i]].Edges[:0]

			for _, edge := range g.Vertexes[g.NodeArray[i]].Edges {
				if edge.Key != from && edge.Key != to {
					current_vertex = append(current_vertex, edge)
				} else {
					removed_edges[edge.Key] = edge
				}
			}
			g.Vertexes[g.NodeArray[i]].Edges = current_vertex
		}
	}

	return removed_edges, nil
}

func (g *Graph) DepthFirstSearch(source_node_key string) *stack.Stack.Items {
  // to-do
}
