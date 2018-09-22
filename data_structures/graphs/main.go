package main

import (
	"fmt"
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

}
