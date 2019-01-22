package ds

import "github.com/pkg/errors"

var (
	ErrEmptyStack = errors.New("Stack is empty")
)

type MyStack struct {
	arr []int
}

func (s *MyStack) NewMyStack() *MyStack {
	return &MyStack{arr: make([]int, 0)}
}

func (s *MyStack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *MyStack) Pop() (int, error) {
	if(len(s.arr) == 0) {
		return 0 , ErrEmptyStack
	}
	val := s.arr[len(s.arr)-1]
	if len(s.arr) == 1 {
		s = &MyStack{arr : make([]int, 0)}
		return val, nil
	}
	s.arr = s.arr[: len(s.arr)-1]
	return val, nil
}

func (s *MyStack) Top() (int, error) {
	if(len(s.arr) == 0) {
		return 0, ErrEmptyStack
	}
	return s.arr[len(s.arr)-1], nil
}

func (s *MyStack) IsEmpty() bool {
	if(len(s.arr) == 0) {
		return true
	}
	return false
}