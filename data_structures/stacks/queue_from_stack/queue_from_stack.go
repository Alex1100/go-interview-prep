package queue_from_stack

import (
	stack "go-interview-prep/data_structures/stacks/stack"
)

type Queue struct {
	entry_stack stack.Stack
	queue       stack.Stack
	size        int
}

func InitQueue() *Queue {
	return &Queue{
		entry_stack: *stack.InitStack(),
		queue:       *stack.InitStack(),
		size:        0,
	}
}

func (q *Queue) Enqueue(item int) error {
	q.entry_stack.Insert(item)

	for q.entry_stack.Size() > 0 {
		popped, err := q.entry_stack.Pop()
		if err == nil {
			q.queue.Insert(popped)
			q.size++
		} else {
			return err
		}
	}

	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.entry_stack.IsEmpty() {
		for !q.queue.IsEmpty() {
			popped, err := q.queue.Pop()

			if err == nil {
				q.entry_stack.Insert(popped)
			} else {
				panic(err)
			}
		}
	}

	removed, err := q.entry_stack.Pop()

	if err == nil {
		return removed, nil
	} else {
		panic(err)
	}
}

func (q *Queue) Peek() int {
	return q.queue.Front()
}

func (q *Queue) ViewQueue() stack.Stack {
	if q.entry_stack.IsEmpty() {
		return q.queue
	} else {
		return q.entry_stack
	}
}

func (q *Queue) IsEmpty() bool {
	if q.Size() == 0 {
		return true
	}

	return false
}

func (q *Queue) Items() []int {
	current_queue := q.ViewQueue()
	return current_queue.Items()
}

func (q *Queue) Size() int {
	return q.size
}
