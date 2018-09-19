package singly_linked_list

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type HashTableLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func InitLinkedList(Data interface{}) *LinkedList {
	Head := &Node{
		Data: Data,
		Next: nil,
	}

	Tail := &Node{
		Data: Data,
		Next: nil,
	}
	return &LinkedList{
		Head: Head,
		Tail: Tail,
		Size: 1,
	}
}

func InitEmptyLinkedList() *LinkedList {
	Head := &Node{
		Data: nil,
		Next: nil,
	}

	Tail := &Node{
		Data: nil,
		Next: nil,
	}
	return &LinkedList{
		Head: Head,
		Tail: Tail,
		Size: 0,
	}
}

func (sll *LinkedList) Insert(Data interface{}) {
	current := sll.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{
		Data: Data,
		Next: nil,
	}

	sll.Tail = current.Next
	sll.Size++
}

func (sll *LinkedList) AppendToHead(Data interface{}) {
	Next := sll.Head.Next

	sll.Head.Next = &Node{
		Data: Data,
		Next: Next,
	}
	sll.Size++
}

func (sll *LinkedList) PrependToTail(Data interface{}) {
	current := sll.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	temp := current.Next
	current.Next = &Node{
		Data: Data,
		Next: temp,
	}

	sll.Size++
}

func (sll *LinkedList) GetByKey(key interface{}) (*Node, error) {
	current := sll.Head

	if current.Data == key {
		return current, nil
	}

	for current.Next != nil {
		current = current.Next
		if current.Data == key {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *LinkedList) GetByIndex(index int) (*Node, error) {
	current := sll.Head
	counter := 0
	if counter == index {
		return current, nil
	}

	for current.Next != nil {
		current = current.Next
		counter++
		if counter == index {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *LinkedList) RemoveNode(key interface{}) *Node {
	current := sll.Head
	prev := sll.Head
	var removed *Node

	if current.Data == key {
		sll.Head = current.Next
		sll.Head.Next = current.Next.Next
		sll.Size--
		return current
	} else {
		for current.Next != nil {
			prev = current
			current = current.Next
			if current.Data == key {
				prev.Next = current.Next
				removed = current
				sll.Size--
				return removed
			}
		}

		sll.Tail = prev
		sll.Size--
		return removed
	}
}

func (sll *LinkedList) ListToSlice() []interface{} {
	list := make([]interface{}, 0)
	current := sll.Head
	list = append(list, current.Data)

	for current.Next != nil {
		current = current.Next
		list = append(list, current.Data)
	}

	return list
}

func (sll *LinkedList) Clone() *LinkedList {
	cloned_linked_list := &LinkedList{
		Head: sll.Head,
		Tail: sll.Tail,
		Size: 1,
	}

	current := sll.Head
	current_cloned := cloned_linked_list.Head
	for current.Next != nil {
		current = current.Next
		current_cloned.Next = current
		cloned_linked_list.Size++
		current_cloned = current_cloned.Next
	}

	return cloned_linked_list
}

// HASH TABLE LINKED LIST
func InitHashTableLinkedList(Data interface{}) *HashTableLinkedList {
	Head := &Node{
		Data: Data,
		Next: nil,
	}

	Tail := &Node{
		Data: Data,
		Next: nil,
	}
	return &HashTableLinkedList{
		Head: Head,
		Tail: Tail,
		Size: 1,
	}
}

func InitEmptyHashTableLinkedList() *HashTableLinkedList {
	Head := &Node{
		Data: nil,
		Next: nil,
	}

	Tail := &Node{
		Data: nil,
		Next: nil,
	}
	return &HashTableLinkedList{
		Head: Head,
		Tail: Tail,
		Size: 0,
	}
}

func (sll *HashTableLinkedList) Insert(Data []interface{}) {
	current := sll.Head
	if current.Data == nil {
		current.Data = Data
		current.Next = nil
		sll.Tail = current
		sll.Size++
	} else {

		for current.Next != nil {
			current = current.Next
		}

		current.Next = &Node{
			Data: Data,
			Next: nil,
		}
		sll.Tail = current.Next
		sll.Size++
	}
}

func (sll *HashTableLinkedList) AppendToHead(Data []interface{}) {
	Next := sll.Head.Next

	sll.Head.Next = &Node{
		Data: Data,
		Next: Next,
	}
	sll.Size++
}

func (sll *HashTableLinkedList) PrependToTail(Data []interface{}) {
	current := sll.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	temp := current.Next
	current.Next = &Node{
		Data: Data,
		Next: temp,
	}

	sll.Size++
}

func (sll *HashTableLinkedList) GetByKey(key []interface{}) (*Node, error) {
	current := sll.Head
	temp1 := make([]interface{}, 0)
	temp2 := make([]interface{}, 0)
	temp1 = append(temp1, current.Data)
	temp2 = append(temp2, key)
	if sliceToString(temp1) == sliceToString(temp2) {
		return current, nil
	}

	for current.Next != nil {
		current = current.Next
		temp3 := make([]interface{}, 0)
		temp4 := make([]interface{}, 0)
		temp3 = append(temp3, current.Data)
		temp4 = append(temp4, key)
		if sliceToString(temp3) == sliceToString(temp4) {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *HashTableLinkedList) GetByIndex(index int) (*Node, error) {
	current := sll.Head
	counter := 0
	if counter == index {
		return current, nil
	}

	for current.Next != nil {
		current = current.Next
		counter++
		if counter == index {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *HashTableLinkedList) RemoveNode(key []interface{}) (interface{}, *HashTableLinkedList) {
	current := sll.Head
	prev := sll.Head
	var removed *Node
	iteration_counter := 0

	temp1 := make([]interface{}, 0)
	temp1 = append(temp1, current.Data)

	if sliceToString(temp1) == sliceToString(key) {
		sll.Head = prev.Next
		sll.Size--
		return current.Data, sll
	} else {
		for iteration_counter < sll.Size-1 {
			prev = current
			current = current.Next
			temp2 := make([]interface{}, 0)
			temp2 = append(temp2, current.Data)

			if sliceToString(temp2) == sliceToString(key) {
				prev.Next = current.Next
				removed = current
				if iteration_counter == sll.Size-1 {
					sll.Tail = prev
				}
				sll.Size--
				return removed, sll
			}
		}

		sll.Tail = prev
		sll.Size--
		return removed, sll
	}
}

func (sll *HashTableLinkedList) ListToSlice() []interface{} {
	list := make([]interface{}, 0)
	current := sll.Head
	list = append(list, current.Data)

	for current.Next != nil {
		current = current.Next
		list = append(list, current.Data)
	}

	return list
}

func (sll *HashTableLinkedList) ListOfKeys() []interface{} {
	list := make([]interface{}, 0)
	current := sll.Head
	data1 := current.Data.([]interface{})[0]
	list = append(list, data1)

	for current.Next != nil {
		current = current.Next
		data2 := current.Data.([]interface{})[0]
		list = append(list, data2)
	}

	return list
}

func (sll *HashTableLinkedList) ListOfValues() []interface{} {
	list := make([]interface{}, 0)
	current := sll.Head
	data1 := current.Data.([]interface{})[1]
	list = append(list, data1)

	for current.Next != nil {
		current = current.Next
		data2 := current.Data.([]interface{})[1]
		list = append(list, data2)
	}

	return list
}

func (sll *HashTableLinkedList) Clone() *HashTableLinkedList {
	cloned_linked_list := &HashTableLinkedList{
		Head: sll.Head,
		Tail: sll.Tail,
		Size: 1,
	}

	current := sll.Head
	current_cloned := cloned_linked_list.Head
	for current.Next != nil {
		current = current.Next
		current_cloned.Next = current
		cloned_linked_list.Size++
		current_cloned = current_cloned.Next
	}

	return cloned_linked_list
}

func sliceToString(values []interface{}) string {
	s := make([]string, len(values)) // Pre-allocate the right size
	for index := range values {
		s[index] = fmt.Sprintf("%v", values[index])
	}
	return strings.Join(s, ",")
}
