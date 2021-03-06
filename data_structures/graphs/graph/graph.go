package graph

import (
	"errors"
	queue "go-interview-prep/data_structures/queues/graph_queue"
	stack "go-interview-prep/data_structures/stacks/graph_stack"
	"math"
	"math/rand"
	"time"
)

type StringIntMap struct {
	Val map[string]int
}

type Prev struct {
	Prev *StringIntMap
}

type DijkstraValue struct {
	ShortestDistance float64
	ShortestPath     []string
}

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

func random(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}

func (g *Graph) AddEdge(from, to string) bool {
	if !g.HasVertexes(to, from) || g.HasEdge(from, to) || g.SameEdge(from, to) {
		return false
	}
	var to_edge_vals map[string]int
	var from_edge_vals map[string]int
	random_weight := random(rand.New(rand.NewSource(time.Now().Unix())), 1, 200)

	to_edge_vals = make(map[string]int)
	to_edge_vals["weight"] = random_weight
	to_edge_vals["direction"] = 1

	from_edge_vals = make(map[string]int)
	from_edge_vals["weight"] = random_weight
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
	random_weight := random(rand.New(rand.NewSource(time.Now().Unix())), 1, 200)

	to_edge_vals = make(map[string]int)
	to_edge_vals["weight"] = random_weight
	to_edge_vals["direction"] = 1

	from_edge_vals = make(map[string]int)
	from_edge_vals["weight"] = random_weight
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

func (g *Graph) DFSUtil(
	source_node string,
	visited map[string]bool,
	visited_stack *stack.Stack,
) *stack.Stack {
	if visited_stack.Size < len(g.NodeArray) {
		if !visited[source_node] {
			visited[source_node] = true
			visited_stack.Insert(source_node)
		}

		for _, edge := range g.Vertexes[source_node].Edges {
			if !visited[edge.Key] {
				visited_stack = g.DFSUtil(edge.Key, visited, visited_stack)
			}
		}
	}
	return visited_stack
}

func (g *Graph) DepthFirstSearch(source_node_key string) (*stack.Stack, error) {
	if g.Vertexes[source_node_key] == nil {
		panic(errors.New("Vertex is not a member of the graph"))
	}

	visited := make(map[string]bool)

	for _, vertex := range g.NodeArray {
		visited[vertex] = false
	}

	visited_stack := &stack.Stack{Items: make([]string, 0), Size: 0}
	visited_stack = g.DFSUtil(source_node_key, visited, visited_stack)

	for _, vertex := range g.NodeArray {
		visited_stack = g.DFSUtil(vertex, visited, visited_stack)
	}

	return visited_stack, nil
}

func (g *Graph) BFSUtil(
	source_node string,
	result *queue.Queue,
	visited map[string]bool,
	node_queue *queue.Queue,
) *queue.Queue {
	if result.Size < len(g.NodeArray) {
		if !visited[source_node] {
			visited[source_node] = true
			result.Enqueue(source_node)
		}

		if node_queue.Contains(source_node) {
			node_queue.Dequeue()
		}
		for _, edge := range g.Vertexes[source_node].Edges {
			if !visited[edge.Key] {
				node_queue.Enqueue(edge.Key)
			}
		}

		for _, node := range node_queue.Items {
			result = g.BFSUtil(node, result, visited, node_queue)
		}
	}

	return result
}

func (g *Graph) BreadthFirstSearch(source_node_key string) (*queue.Queue, error) {
	if g.Vertexes[source_node_key] == nil {
		panic(errors.New("Vertex is not a member of the graph"))
	}

	visited := make(map[string]bool)
	result := &queue.Queue{Items: make([]string, 0), Size: 0}
	node_queue := &queue.Queue{Items: make([]string, 0), Size: 0}

	for _, vertex := range g.NodeArray {
		visited[vertex] = false
	}

	result = g.BFSUtil(source_node_key, result, visited, node_queue)

	for _, vertex := range g.NodeArray {
		result = g.BFSUtil(vertex, result, visited, node_queue)
	}

	return result, nil
}

func (g *Graph) CostLength(u, v string) int {
	for _, edge := range g.Vertexes[u].Edges {
		if edge.Key == v {
			return edge.EdgeValues["weight"]
		}
	}

	return 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (g *Graph) ExtractDijkstraMin(
	dj_set map[string]*Vertex,
	dist map[string]float64,
	deleted_set []string,
) string {
	minimum_distance := math.Inf(1)
	var node_with_min_distance string

	for vertex, _ := range dj_set {
		if contains(deleted_set, vertex) == false && dist[vertex] <= minimum_distance {
			minimum_distance = dist[vertex]
			node_with_min_distance = vertex
		}
	}

	return node_with_min_distance
}

func (g *Graph) GetShortestPath(
	destination_node string,
	shortest_path map[string][]string,
	prev map[string]string,
	dist map[string]float64,
) []string {
	count := len(prev)
	path := shortest_path[destination_node]
	node := destination_node

	for count > 0 {
		if len(node) >= 1 {
			path = append(path, node)
			node = prev[node]
		}

		count--
	}

	return ReverseStringSlice(path)
}

func ReverseStringSlice(slice []string) []string {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	return slice
}

func (g *Graph) Dijkstra(source_node, destination_node string) *DijkstraValue {
	dj_set := make(map[string]*Vertex, 0)
	shortest_path := make(map[string][]string)
	dist := make(map[string]float64)
	prev := make(map[string]string)
	deleted_set := make([]string, 0)

	for _, vertex := range g.NodeArray {
		dist[vertex] = math.Inf(0)
		dj_set[vertex] = g.Vertexes[vertex]
	}

	dist[source_node] = float64(0)
	shortest_path[destination_node] = make([]string, 0)
	counter := len(dj_set)

	for len(deleted_set) != counter {
		u := g.ExtractDijkstraMin(dj_set, dist, deleted_set)

		if g.Vertexes[u] != nil {
			for _, edge := range g.Vertexes[u].Edges {
				alt := dist[u] + float64(g.CostLength(u, edge.Key))

				if alt < dist[edge.Key] {
					dist[edge.Key] = alt
					prev[edge.Key] = u
				}
			}
		}

		deleted_set = append(deleted_set, u)
	}

	result := g.GetShortestPath(destination_node, shortest_path, prev, dist)

	return &DijkstraValue{
		ShortestDistance: dist[destination_node],
		ShortestPath:     result,
	}
}

func (g *Graph) HasCycleUtil(
	source_node string,
	visited map[string]bool,
	visited_stack *stack.Stack,
) bool {
	if !visited[source_node] {
		visited[source_node] = true
		visited_stack.Insert(source_node)

		for _, edge := range g.Vertexes[source_node].Edges {
			if !visited[edge.Key] && edge.EdgeValues["direction"] == 1 && g.HasCycleUtil(edge.Key, visited, visited_stack) {
				return true
			} else if visited_stack.Contains(edge.Key) && edge.EdgeValues["direction"] == 1 {
				return true
			}
		}
	}

	visited_stack.Pop()
	return false
}

func (g *Graph) HasCycle() bool {
	visited := make(map[string]bool)
	visited_stack := *stack.InitStack()
	return g.HasCycleUtil(g.NodeArray[0], visited, &visited_stack)
}

func (g *Graph) TopologicalSort() ([]string, error) {
	if g.HasCycle() {
		return make([]string, 0), errors.New("Can't sort a Cyclic Graph")
	}

	visited := make(map[string]bool)
	non_dependent_vertexes := *stack.InitStack()
	vertex_stack := *stack.InitStack()
	result := make([]string, 0)

	for _, vertex := range g.NodeArray {
		for _, edge := range g.Vertexes[vertex].Edges {
			if edge.EdgeValues["direction"] == 1 {
				visited[vertex] = true
				non_dependent_vertexes.Insert(vertex)
				break
			}
		}
	}

	for non_dependent_vertexes.Size > 0 {
		for _, edge := range g.Vertexes[non_dependent_vertexes.Front()].Edges {
			if !visited[edge.Key] && edge.EdgeValues["direction"] == 1 {
				vertex_stack.Insert(edge.Key)
				visited[edge.Key] = true
			}
		}

		popped, err := non_dependent_vertexes.Pop()
		if err == nil {
			vertex_stack.Insert(popped)
		} else {
			return make([]string, 0), errors.New("Errored Popping item off stack")
		}
	}

	for vertex_stack.Size > 0 {
		popped, err := vertex_stack.Pop()
		if err == nil {
			result = append(result, popped)
		} else {
			return make([]string, 0), errors.New("Errored Popping item off stack")
		}
	}

	return result, nil
}
