package queue_from_stack

import (
	stack "go-interview-prep/data_structures/stacks/stack"
)

type Queue struct {
	EntryStack stack.Stack
	Queue      stack.Stack
	Size       int
}

func InitQueue() *Queue {
	return &Queue{
		EntryStack: *stack.InitStack(),
		Queue:      *stack.InitStack(),
		Size:       0,
	}
}

func (q *Queue) Enqueue(item int) error {
	q.EntryStack.Insert(item)

	for q.EntryStack.Size() > 0 {
		popped, err := q.EntryStack.Pop()
		if err == nil {
			q.Queue.Insert(popped)
			q.Size++
		} else {
			return err
		}
	}

	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.EntryStack.IsEmpty() {
		for !q.Queue.IsEmpty() {
			popped, err := q.Queue.Pop()

			if err == nil {
				q.EntryStack.Insert(popped)
			} else {
				panic(err)
			}
		}
	}

	removed, err := q.EntryStack.Pop()

	if err == nil {
		return removed, nil
	} else {
		panic(err)
	}
}

func (q *Queue) Peek() int {
	return q.Queue.Front()
}

func (q *Queue) ViewQueue() stack.Stack {
	if q.EntryStack.IsEmpty() {
		return q.Queue
	} else {
		return q.EntryStack
	}
}

func (q *Queue) IsEmpty() bool {
	if q.Size == 0 {
		return true
	}

	return false
}
