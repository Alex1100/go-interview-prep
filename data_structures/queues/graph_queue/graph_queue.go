package graph_queue

import "errors"

type Queue struct {
	Items []string
	Size  int
}

func InitQueue() *Queue {
	return &Queue{
		Items: make([]string, 0),
		Size:  0,
	}
}

func (q *Queue) Enqueue(item string) error {
	q.Items = append(q.Items, item)
	return nil
}

func (q *Queue) Dequeue() (string, error) {
	if len(q.Items) < 1 {
		panic(errors.New("Nothing to Dequeue"))
	}

	removed := q.Items[0]
	q.Items = q.Items[1:]
	return removed, nil
}

func (q *Queue) Peek() string {
	return q.Items[0]
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Contains(data string) bool {
	for _, item := range q.Items {
		if item == data {
			return true
		}
	}
	return false
}
