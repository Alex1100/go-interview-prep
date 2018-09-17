package singly_linked_list

import "errors"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func InitLinkedList(data interface{}) *LinkedList {
	head := &Node{
		data: data,
		next: nil,
	}

	tail := &Node{
		data: data,
		next: nil,
	}
	return &LinkedList{
		head: head,
		tail: tail,
		size: 1,
	}
}

func (sll *LinkedList) Head() *Node {
	return sll.head
}

func (sll *LinkedList) Data(node *Node) interface{} {
	return node.data
}

func (sll *LinkedList) Next(node *Node) *Node {
	return node.next
}

func (sll *LinkedList) Insert(data interface{}) {
	current := sll.head

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{
		data: data,
		next: nil,
	}

	sll.tail = current.next
	sll.size++
}

func (sll *LinkedList) AppendToHead(data interface{}) {
	next := sll.head.next

	sll.head.next = &Node{
		data: data,
		next: next,
	}
	sll.size++
}

func (sll *LinkedList) PrependToTail(data interface{}) {
	current := sll.head

	for current.next.next != nil {
		current = current.next
	}

	temp := current.next
	current.next = &Node{
		data: data,
		next: temp,
	}

	sll.size++
}

func (sll *LinkedList) GetByKey(key interface{}) (*Node, error) {
	current := sll.head

	if current.data == key {
		return current, nil
	}

	for current.next != nil {
		current = current.next
		if current.data == key {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *LinkedList) GetByIndex(index int) (*Node, error) {
	current := sll.head
	counter := 0
	if counter == index {
		return current, nil
	}

	for current.next != nil {
		current = current.next
		counter++
		if counter == index {
			return current, nil
		}
	}

	panic(errors.New("Does Not Exist"))
}

func (sll *LinkedList) RemoveNode(key interface{}) *Node {
	current := sll.head
	prev := sll.head
	var removed *Node

	if current.data == key {
		sll.head = current.next
		sll.head.next = current.next.next
		sll.size--
		return current
	} else {
		for current.next != nil {
			prev = current
			current = current.next
			if current.data == key {
				prev.next = current.next
				removed = current
				sll.size--
				return removed
			}
		}

		sll.tail = prev
		sll.size--
		return removed
	}
}

func (sll *LinkedList) ListToSlice() []interface{} {
	list := make([]interface{}, 0)
	current := sll.head
	list = append(list, current.data)

	for current.next != nil {
		current = current.next
		list = append(list, current.data)
	}

	return list
}

func (sll *LinkedList) Clone() *LinkedList {
	cloned_linked_list := &LinkedList{
		head: sll.head,
		tail: sll.tail,
		size: 1,
	}

	current := sll.head
	current_cloned := cloned_linked_list.head
	for current.next != nil {
		current = current.next
		current_cloned.next = current
		cloned_linked_list.size++
		current_cloned = current_cloned.next
	}

	return cloned_linked_list
}

func (sll *LinkedList) Size() int {
	return sll.size
}
