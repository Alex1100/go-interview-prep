package main

import (
	"fmt"
	binary_search_tree "go-interview-prep/data_structures/binary_search_trees/binary_search_tree"
	"math"
)

func main() {
	bst := *binary_search_tree.InitEmptyBST()
	bst.AddNode(100)
	bst.AddNode(80)
	bst.AddNode(120)
	bst.AddNode(95)
	bst.AddNode(75)
	bst.AddNode(110)
	bst.AddNode(125)
	bst.AddNode(135)

	fmt.Println("::DEPTH-FIRST-SEARCH::\n")
	pre, err := bst.DepthFirstSearch("pre_order")
	if err == nil {
		fmt.Println("PRE-ORDER: ", pre, "\n")
	} else {
		fmt.Println(err)
	}

	in, err := bst.DepthFirstSearch("in_order")
	if err == nil {
		fmt.Println("IN-ORDER: ", in, "\n")
	} else {
		fmt.Println(err)
	}

	post, err := bst.DepthFirstSearch("post_order")
	if err == nil {
		fmt.Println("POST-ORDER: ", post, "\n")
	} else {
		fmt.Println(err)
	}

	bfs := bst.BreadthFirstSearch()
	fmt.Println("::Breadth-First-Search::\n\n", bfs, "\n")
	fmt.Println(bst.IsValidBST(math.Inf(-1), math.Inf(1)))
	fmt.Println(bst.FindInOrderSuccessor(100).Data.Value)
	fmt.Println(bst.GetMinDepth())
	fmt.Println(bst.GetMaxPathSum())
	fmt.Println(bst.ContainsNode(135))
	found, err := bst.FindNode(130)
	if err == nil {
		fmt.Println(found.Data)
	} else {
		fmt.Println(err)
	}

	fmt.Println(bst.ContainsNode(135))
	fmt.Println(bst.FindMax().Data.Value)
	fmt.Println(bst.FindMin().Data.Value)
}
