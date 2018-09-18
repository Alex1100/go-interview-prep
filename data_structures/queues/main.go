package main

import (
	"fmt"
	pq "go-interview-prep/data_structures/queues/priority_queue"
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
	fmt.Println("STARTING MIN HEAP OPERATIONS::::")
	min_pq := *pq.InitMinPriorityQueue()
	min_pq.Enqueue(1000)
	min_pq.Enqueue(2000)
	min_pq.Enqueue(3000)
	fmt.Println(min_pq.Items())
	fmt.Println(min_pq.Dequeue())
	fmt.Println(min_pq.Items())
	fmt.Println("STARTING MAX HEAP OPERATIONS::::")
	max_pq := *pq.InitMaxPriorityQueue()
	max_pq.Enqueue(8000)
	max_pq.Enqueue(9000)
	max_pq.Enqueue(12000)
	fmt.Println(max_pq.Items())
	fmt.Println(max_pq.Dequeue())
	fmt.Println(max_pq.Items())
}
