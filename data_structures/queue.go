package ds

import "github.com/pkg/errors"

type MyQueue struct {
	arr []int
}

func (q *MyQueue) NewMyQueue() *MyQueue {
	q.arr = []int{}
	return q
}

func (q *MyQueue) Push(val int) {
	q.arr = append(q.arr, val)
}

func (q *MyQueue) Pop() (*int, error) {
	if(len(q.arr) == 0) {
		return nil, errors.New("EmptyQueueException: Queue is empty")
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	return &val, nil
}

func (q *MyQueue) Top() (*int, error) {
	if(len(q.arr) == 0) {
		return nil, errors.New("EmptyQueueException: Queue is empty")
	}
	val := q.arr[0]
	return &val, nil
}

func (q *MyQueue) IsEmpty() bool {
	if(len(q.arr) == 0) {
		return true
	}
	return false
}
