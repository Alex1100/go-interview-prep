package main

import (
	"fmt"
	queue "go-interview-prep/data_structures/stacks/queue_from_stack"
	stack "go-interview-prep/data_structures/stacks/stack"
)

func main() {
	oo := *stack.InitStack()
	oo.Insert(99)
	oo.Insert(100)
	oo.Insert(120)
	fmt.Println(oo.Pop())
	fmt.Println(oo.Items)

	fmt.Println("NOW WORKING ON QUEUE FROM STACKS:::")
	qq := *queue.InitQueue()
	qq.Enqueue(10)
	qq.Enqueue(30)
	qq.Enqueue(150)
	fmt.Println(qq.Dequeue())
	fmt.Println(qq.Items)
	qq.Enqueue(55)
	fmt.Println(qq.ViewQueue())
}
