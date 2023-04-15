package queue

import (
	"fmt"
)

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return -1, fmt.Errorf("Queue is empty")
	}
	if len(q.elements) == 1 {
		dequeuedElement := q.elements[0]
		q.elements = []int{}
		return dequeuedElement, nil
	}
	dequeuedElement := q.elements[0]
	q.elements = q.elements[1:]
	return dequeuedElement, nil
}

func (q *Queue) Capacity() (int, int) {
	return len(q.elements), cap(q.elements)
}

func (q *Queue) Display() {
	fmt.Println("Elements in queue:", q.elements)
}
