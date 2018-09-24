package adjacency_matrix

type Vertex struct {
	Edges []int
}

type Graph struct {
	Vertexes  []*Vertex
	NodeArray []string
}

func InitAdjacencyMatrix() *Graph {
	return &Graph{
		Vertexes:  make([]*Vertex, 0),
		NodeArray: make([]string, 0),
	}
}

func (g *Graph) AddVertex(vertex_key string) bool {
	if g.ContainsVertex(vertex_key) {
		return false
	}

	g.NodeArray = append(g.NodeArray, vertex_key)
	g.Vertexes = append(g.Vertexes, &Vertex{Edges: make([]int, 0)})

	for i, _ := range g.NodeArray {
		new_vertex := make([]int, 0)
		if len(g.Vertexes[i].Edges) > 0 {
			for _, edge := range g.Vertexes[i].Edges {
				new_vertex = append(new_vertex, edge)
			}
			new_vertex = append(new_vertex, 0)

		} else {
			for j := 0; j < len(g.NodeArray); j++ {
				new_vertex = append(new_vertex, 0)
			}
		}

		g.Vertexes[i].Edges = new_vertex
	}

	return true
}

func (g *Graph) ContainsVertex(vertex_key string) bool {
	for _, vertex := range g.NodeArray {
		if vertex == vertex_key {
			return true
		}
	}

	return false
}

func (g *Graph) AddEdge(from, to string) bool {
	if g.SameNodes(from, to) {
		return false
	} else if g.HasEdge(to, from) {
		var from_index int
		var to_index int

		for idx, vertex := range g.NodeArray {
			if vertex == from {
				from_index = idx
			}

			if vertex == to {
				to_index = idx
			}
		}

		g.Vertexes[from_index].Edges[to_index] = 1
		g.Vertexes[to_index].Edges[from_index] = 1
		return true
	} else {
		var from_index int
		var to_index int

		for idx, vertex := range g.NodeArray {
			if vertex == from {
				from_index = idx
			}

			if vertex == to {
				to_index = idx
			}
		}

		g.Vertexes[from_index].Edges[to_index] = 1
		g.Vertexes[to_index].Edges[from_index] = -1
		return true
	}
	return false
}

func (g *Graph) AddEdges(from, to string) bool {
	if g.SameNodes(from, to) {
		return false
	} else if g.HasEdge(to, from) {
		return false
	} else {
		var from_index int
		var to_index int

		for idx, vertex := range g.NodeArray {
			if vertex == from {
				from_index = idx
			}

			if vertex == to {
				to_index = idx
			}
		}

		g.Vertexes[from_index].Edges[to_index] = 1
		g.Vertexes[to_index].Edges[from_index] = 1
		return true
	}
}

func (g *Graph) HasEdge(from, to string) bool {
	var left bool = false
	var right bool = false

	for i, _ := range g.Vertexes {
		for j, _ := range g.Vertexes[i].Edges {
			if g.NodeArray[j] == from && g.Vertexes[i].Edges[j] > 0 {
				left = true
			}

			if g.NodeArray[j] == to && (g.Vertexes[i].Edges[j] > 0 || g.Vertexes[i].Edges[j] < 0) {
				right = true
			}
		}
	}
	return left && right
}

func (g *Graph) SameNodes(from, to string) bool {
	return from == to
}
