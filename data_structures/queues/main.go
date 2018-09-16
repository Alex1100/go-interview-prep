package main

import (
	"fmt"
	queue "go-interview-prep/data_structures/queues/queue"
)

func main() {
	q := *queue.InitQueue()
	q.Enqueue(10)
	q.Enqueue(12)
	q.Enqueue(24)
	fmt.Println(q.Items())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Items())
}
