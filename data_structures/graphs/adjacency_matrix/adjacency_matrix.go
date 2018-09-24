package adjacency_matrix

import (
	"errors"
	graph_queue "go-interview-prep/data_structures/queues/graph_queue"
	graph_stack "go-interview-prep/data_structures/stacks/graph_stack"
)

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
	left := g.Vertexes[from_index].Edges[to_index]
	right := g.Vertexes[to_index].Edges[from_index]
	if (left > 0 || left < 0) && (right > 0 || right < 0) {
		return true
	}

	return false
}

func (g *Graph) SameNodes(from, to string) bool {
	return from == to
}

func (g *Graph) DFSUtil(source_node string, visited map[string]bool, visited_stack graph_stack.Stack) []string {
	if visited_stack.Size == len(g.Vertexes) {
		return visited_stack.Items
	}

	if !visited[source_node] {
		visited_stack.Insert(source_node)
		visited[source_node] = true
	}

	var source_index int

	for i, _ := range g.NodeArray {
		if g.NodeArray[i] == source_node {
			source_index = i
		}
	}

	for i, edge := range g.Vertexes[source_index].Edges {
		if !visited[g.NodeArray[i]] && edge == 1 {
			visited_stack.Items = g.DFSUtil(g.NodeArray[i], visited, visited_stack)
		}
	}

	return visited_stack.Items
}

func (g *Graph) DepthFirstSearch(vertex_key string) []string {
	visited := make(map[string]bool)
	visited_stack := *graph_stack.InitStack()
	visited_stack.Items = g.DFSUtil(vertex_key, visited, visited_stack)

	for _, vertex := range g.NodeArray {
		if !visited[vertex] {
			visited_stack.Items = g.DFSUtil(vertex, visited, visited_stack)
		}
	}

	return visited_stack.Items
}

func (g *Graph) BFSUtil(source_node string, visited map[string]bool, vertex_queue graph_queue.Queue, result_queue graph_queue.Queue) ([]string, error) {
	if result_queue.Size == len(g.NodeArray) {
		return result_queue.Items, nil
	}

	if !visited[source_node] {
		result_queue.Enqueue(source_node)
		visited[source_node] = true
	}

	if vertex_queue.Contains(source_node) {
		_, err := vertex_queue.Dequeue()
		if err != nil {
			return make([]string, 0), errors.New("Error popping queue")
		}
	}

	var source_index int

	for i, _ := range g.NodeArray {
		if g.NodeArray[i] == source_node {
			source_index = i
		}
	}

	for j, edge := range g.Vertexes[source_index].Edges {
		if !visited[g.NodeArray[j]] && edge == 1 {
			vertex_queue.Enqueue(g.NodeArray[j])
		}
	}

	for z := 0; z < len(vertex_queue.Items); z++ {
		vertex, err := vertex_queue.Dequeue()
		if err == nil {
			res, err := g.BFSUtil(vertex, visited, vertex_queue, result_queue)
			if err == nil {
				result_queue.Items = res
			} else {
				return make([]string, 0), errors.New("Error popping vetex")
			}
		} else {
			return make([]string, 0), errors.New("Error popping vertex")
		}
	}

	return result_queue.Items, nil
}

func (g *Graph) BreadthFirstSearch(vertex_key string) ([]string, error) {
	visited := make(map[string]bool)
	vertex_queue := *graph_queue.InitQueue()
	result_queue := *graph_queue.InitQueue()
	res, err := g.BFSUtil(vertex_key, visited, vertex_queue, result_queue)
	if err == nil {
		result_queue.Items = res
	} else {
		return make([]string, 0), errors.New("Error popping vetex")
	}

	for _, vertex := range g.NodeArray {
		if !visited[vertex] {
			result, err := g.BFSUtil(vertex, visited, vertex_queue, result_queue)
			if err == nil {
				result_queue.Items = result
			} else {
				return make([]string, 0), errors.New("Error popping vetex")
			}
		}
	}

	return result_queue.Items, nil
}
