package main

import (
	"fmt"
	singly_linked_list "go-interview-prep/data_structures/linked_lists/singly_linked_list"
)

func main() {
	sll := *singly_linked_list.InitLinkedList("head")
	sll.Insert(20)
	sll.Insert(33)
	sll.Insert(44)
	sll.Insert(55)
	sll.Insert(66)
	fmt.Println(sll.Data(sll.Head()))
	fmt.Println(sll.Next(sll.Head()))
	sll.AppendToHead(30)
	sll.PrependToTail(120)
	fmt.Println(sll.Next(sll.Head()))
	fmt.Println(sll.ListToSlice())
	fmt.Println(sll.Clone().ListToSlice())
}
