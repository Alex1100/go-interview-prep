package priority_queue

import (
	max_heap "go-interview-prep/data_structures/heaps/max_heap"
	min_heap "go-interview-prep/data_structures/heaps/min_heap"
)

type MaxPriorityQueue struct {
	queue max_heap.MaxHeap
}

type MinPriorityQueue struct {
	queue min_heap.MinHeap
}

func InitMaxPriorityQueue() *MaxPriorityQueue {
	return &MaxPriorityQueue{
		queue: *max_heap.InitMaxHeap(),
	}
}

func InitMinPriorityQueue() *MinPriorityQueue {
	return &MinPriorityQueue{
		queue: *min_heap.InitMinHeap(),
	}
}

// Max Priority Queue Funcs
func (max_pq *MaxPriorityQueue) Enqueue(item int) {
	max_pq.queue.Add(item)
}

func (max_pq *MaxPriorityQueue) Dequeue() (int, error) {
	removed, err := max_pq.queue.RemoveHead()

	if err != nil {
		panic(err)
	}

	return removed, nil
}

func (max_pq *MaxPriorityQueue) Front() int {
	return max_pq.queue.HeapArray()[0]
}

func (max_pq *MaxPriorityQueue) Size() int {
	return max_pq.queue.Size()
}

func (max_pq *MaxPriorityQueue) IsEmpty() bool {
	return max_pq.Size() == 0
}

func (max_pq *MaxPriorityQueue) Items() []int {
	return max_pq.queue.HeapArray()
}

// Min Priority Queue Funcs

func (min_pq *MinPriorityQueue) Enqueue(item int) {
	min_pq.queue.Add(item)
}

func (min_pq *MinPriorityQueue) Dequeue() (int, error) {
	removed, err := min_pq.queue.RemoveHead()

	if err != nil {
		panic(err)
	}

	return removed, nil
}

func (min_pq *MinPriorityQueue) Front() int {
	return min_pq.queue.HeapArray()[0]
}

func (min_pq *MinPriorityQueue) Size() int {
	return min_pq.queue.Size()
}

func (min_pq *MinPriorityQueue) IsEmpty() bool {
	return min_pq.Size() == 0
}

func (min_pq *MinPriorityQueue) Items() []int {
	return min_pq.queue.HeapArray()
}
