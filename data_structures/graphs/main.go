package main

import (
	"fmt"
	adjacency_matrix "go-interview-prep/data_structures/graphs/adjacency_matrix"
	graph "go-interview-prep/data_structures/graphs/graph"
)

func main() {
	g := *graph.InitGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddVertex("G")
	g.AddVertex("H")
	fmt.Println(g)
	g.AddEdge("A", "H")
	g.AddEdge("H", "F")
	g.AddEdges("E", "D")
	g.AddEdge("G", "C")
	fmt.Println("A EDGES: ", g.Vertexes["A"].Edges[0])
	fmt.Println("H EDGES: ", g.Vertexes["H"].Edges[0])
	removed, err := g.RemoveVertex("A")

	if err == nil {
		fmt.Println(removed, g.Vertexes["H"])
	} else {
		fmt.Println(err)
	}

	removed_edges, err := g.RemoveEdges("H", "F")

	if err == nil {
		fmt.Println(removed_edges)
	} else {
		fmt.Println(err)
	}

	fmt.Println(g.Vertexes["E"].Edges[0], g.Vertexes["D"].Edges[0])
	stat, err := g.RemoveEdge("E", "D")
	if err == nil {
		fmt.Println(stat)
	} else {
		fmt.Println(err)
	}

	fmt.Println(g.Vertexes["E"].Edges[0], g.Vertexes["D"].Edges[0])
	statu, err := g.RemoveEdge("E", "D")
	if err == nil {
		fmt.Println(statu)
	} else {
		fmt.Println(err)
	}

	fmt.Println(g.Vertexes["E"].Edges, g.Vertexes["D"].Edges)

	g.AddEdge("C", "E")
	g.AddEdge("C", "F")
	g.AddEdges("D", "H")
	g.AddEdges("H", "C")
	g.AddEdges("B", "E")
	g.AddEdges("E", "F")
	g.AddEdge("B", "H")
	g.AddEdges("E", "D")
	dfs, err := g.DepthFirstSearch("E")

	if err == nil {
		fmt.Println(dfs.Items)
	} else {
		fmt.Println(err)
	}

	bfs, err := g.BreadthFirstSearch("D")

	if err == nil {
		fmt.Println(bfs.Items)
	} else {
		fmt.Println(err)
	}

	dijkstra := g.Dijkstra("C", "B")
	fmt.Println("PATH COSTS ARE: ", dijkstra.ShortestDistance)
	fmt.Println("PATHS ARE: ", len(dijkstra.ShortestPath), dijkstra.ShortestPath)
	fmt.Println("CYCLIC???: ", g.HasCycle())
	gr := *graph.InitGraph()
	gr.AddVertex("1234")
	gr.AddVertex("098")
	gr.AddEdge("1234", "098")
	gr.AddVertex("333")
	gr.AddEdge("098", "333")
	gr.AddVertex("433")
	gr.AddEdge("098", "433")
	gr.AddVertex("998")
	gr.AddEdge("1234", "998")

	fmt.Println("CYCLIC???: ", gr.HasCycle())
	ordering, err := g.TopologicalSort()

	if err == nil {
		fmt.Println(ordering)
	} else {
		fmt.Println(err)
	}

	toppped, err := gr.TopologicalSort()

	if err == nil {
		fmt.Println(toppped)
	} else {
		fmt.Println(err)
	}

	adj_matrix := *adjacency_matrix.InitAdjacencyMatrix()

	adj_matrix.AddVertex("A")
	adj_matrix.AddVertex("B")
	adj_matrix.AddVertex("C")
	adj_matrix.AddVertex("D")
	adj_matrix.AddVertex("E")
	adj_matrix.AddVertex("F")
	adj_matrix.AddVertex("G")
	adj_matrix.AddVertex("H")
	fmt.Println(adj_matrix.Vertexes[0], adj_matrix.Vertexes[1], adj_matrix.Vertexes[2])
	adj_matrix.AddEdge("A", "B")
	adj_matrix.AddEdge("C", "A")
	adj_matrix.AddEdge("A", "F")
	adj_matrix.AddEdge("F", "H")
	adj_matrix.AddEdge("H", "G")
	fmt.Println(adj_matrix.HasEdge("A", "B"))
	fmt.Println(adj_matrix.HasEdge("C", "A"))
	fmt.Println(adj_matrix.HasEdge("C", "B"))
	fmt.Println(adj_matrix.Vertexes[0], adj_matrix.Vertexes[1], adj_matrix.Vertexes[2])
	fmt.Println(adj_matrix.Vertexes[3], adj_matrix.Vertexes[4], adj_matrix.Vertexes[5])
	fmt.Println(adj_matrix.Vertexes[6], adj_matrix.Vertexes[7])
	fmt.Println(adj_matrix.DepthFirstSearch("A"))
	fmt.Println(adj_matrix.DepthFirstSearch("B"))
	fmt.Println(adj_matrix.DepthFirstSearch("C"))
	fmt.Println(adj_matrix.BreadthFirstSearch("A"))
	fmt.Println(adj_matrix.BreadthFirstSearch("B"))
	fmt.Println(adj_matrix.BreadthFirstSearch("C"))
	fmt.Println(adj_matrix.BreadthFirstSearch("G"))
	fmt.Println(adj_matrix.BreadthFirstSearch("H"))
	fmt.Println(adj_matrix.HasCycle())
	fmt.Println(adj_matrix.FindCycle())

	adj_matrix_2 := *adjacency_matrix.InitAdjacencyMatrix()
	adj_matrix_2.AddVertex("A")
	adj_matrix_2.AddVertex("B")
	adj_matrix_2.AddVertex("C")
	adj_matrix_2.AddEdge("A", "B")
	adj_matrix_2.AddEdge("A", "C")
	fmt.Println(adj_matrix_2.HasCycle())
	fmt.Println(adj_matrix_2.FindCycle())
	adj_matrix_2.AddEdge("C", "A")
	fmt.Println(adj_matrix_2.HasCycle())
	fmt.Println(adj_matrix_2.FindCycle())
}
