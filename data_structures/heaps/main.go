package main

import (
	"fmt"
	max_heap "go-interview-prep/data_structures/heaps/max_heap"
)

func main() {
	mh := *max_heap.InitMaxHeap()
	mh.Add(10)
	mh.Add(9)
	mh.Add(14)
	mh.Add(13)
	mh.Add(8)
	mh.Add(23)
	mh.Add(15)
	mh.Add(30)
	mh.Add(26)
	fmt.Println(mh.HeapArray())
	fmt.Println(mh.HeapArray()[mh.GetParentIndex(mh.Size()-1)])
	fmt.Println(mh.HeapArray()[mh.GetLeftChild(mh.GetParentIndex(mh.Size()-1))])
	fmt.Println(mh.HeapArray()[mh.GetRightChild(mh.GetParentIndex(mh.Size()-1))])
	fmt.Println(mh.Contains(mh.HeapArray()[0]))
	val, err := mh.RemoveHead()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	fmt.Println(mh.Contains(val))
	fmt.Println(mh.HasNode(val))
}
