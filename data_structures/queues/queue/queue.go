package queue

import "errors"

type Queue struct {
	items []int
	size  int
}

func InitQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
		size:  0,
	}
}

func (q *Queue) Enqueue(item int) error {
	q.items = append(q.items, item)
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) < 1 {
		panic(errors.New("Nothing to Dequeue"))
	}

	removed := q.items[0]
	q.items = q.items[1:]
	return removed, nil
}

func (q *Queue) Peek() int {
	return q.items[0]
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Items() []int {
	return q.items
}
