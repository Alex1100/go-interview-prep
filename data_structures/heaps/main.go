package main

import (
	"fmt"
	max_heap "go-interview-prep/data_structures/heaps/max_heap"
	min_heap "go-interview-prep/data_structures/heaps/min_heap"
)

func main() {
	fmt.Println("STARTING MAX HEAP OPERATIONS::::")
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

	fmt.Println("\n\n\nSTARTING MIN HEAP OPERATIONS::::")
	mnh := *min_heap.InitMinHeap()
	mnh.Add(10)
	mnh.Add(9)
	mnh.Add(14)
	mnh.Add(13)
	mnh.Add(8)
	mnh.Add(23)
	mnh.Add(15)
	mnh.Add(30)
	mnh.Add(26)
	fmt.Println(mnh.HeapArray())
	fmt.Println(mnh.HeapArray()[mnh.GetParentIndex(mnh.Size()-1)])
	fmt.Println(mnh.HeapArray()[mnh.GetLeftChild(mnh.GetParentIndex(mnh.Size()-1))])
	fmt.Println(mnh.HeapArray()[mnh.GetRightChild(mnh.GetParentIndex(mnh.Size()-1))])
	fmt.Println(mnh.Contains(mnh.HeapArray()[0]))
	value, er := mnh.RemoveHead()
	if er != nil {
		panic(er)
	}
	fmt.Println(value)
	fmt.Println(mnh.Contains(value))
	fmt.Println(mnh.HasNode(value))
}
