package ds

import "github.com/pkg/errors"

type MyStack struct {
	arr []int
}

func (s *MyStack) NewMyStack() *MyStack {
	s.arr = []int{}
	return s
}

func (s *MyStack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *MyStack) Pop() (*int, error) {
	if(len(s.arr) == 0) {
		return nil , errors.New("EmptyStackException: Stack is empty")
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return &val, nil
}

func (s *MyStack) Top() (*int, error) {
	if(len(s.arr) == 0) {
		return nil, errors.New("EmptyStackException: Stack is empty")
	}
	val := s.arr[len(s.arr)-1]
	return &val, nil
}

func (s *MyStack) IsEmpty() bool {
	if(len(s.arr) == 0) {
		return true
	}
	return false
}