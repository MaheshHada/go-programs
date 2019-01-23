package data_structures

import "github.com/pkg/errors"

var (
	ErrEmptyQueue = errors.New("Queue is empty")
)

type MyQueue struct {
	arr []int
}

func (q *MyQueue) NewMyQueue() *MyQueue {
	return &MyQueue{arr : make([]int, 0)}
}

func (q *MyQueue) Push(val int) {
	q.arr = append(q.arr, val)
}

func (q *MyQueue) Pop() (int, error) {
	if len(q.arr) == 0 {
		return 0, ErrEmptyQueue
	}
	val := q.arr[0]
	if len(q.arr) == 1 {
		q.arr = make([]int, 0)
		return val, nil
	}
	q.arr = q.arr[1:]
	return val, nil
}

func (q *MyQueue) Top() (int, error) {
	if len(q.arr) == 0 {
		return 0, ErrEmptyQueue
	}
	return q.arr[0], nil
}

func (q *MyQueue) IsEmpty() bool {
	if len(q.arr) == 0 {
		return true
	}
	return false
}
