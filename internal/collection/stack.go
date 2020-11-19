package collection

import "errors"

type Stack struct {
	contents []interface{}
	length   int
}

func NewStack() *Stack {
	return &Stack{
		contents: make([]interface{}, 0),
		length:   0,
	}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() (interface{}, error) {
	if s.length < 1 {
		return nil, errors.New("stack is empty")
	}
	return s.contents[s.length-1], nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length < 1 {
		return nil, errors.New("stack is empty")
	}
	popped := s.contents[s.length-1]
	s.length--
	return popped, nil
}

func (s *Stack) Push(element interface{}) {
	s.length++
	if len(s.contents) < s.length {
		// there is not enough space in s.inner. we need to append to the end of it to cause a re-allocation
		s.contents = append(s.contents, element)
	} else {
		// the s.inner slice already has enough space for this element
		s.contents[s.length-1] = element
	}
}
