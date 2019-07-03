package pkg

import "fmt"

type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}
type Queue struct {
	len int
	Elem []*Node
}
func NewQueue() *Queue {
	return &Queue{0, []*Node{}}
}

func (q *Queue) AddQueue(item *Node) {
	q.len++
	q.Elem = append(q.Elem, item)
}

func (q *Queue) IsEmpty() bool {
	return q.len==0
}

func (q *Queue) Top() *Node {
	return q.Elem[0]
}

func (q *Queue) DeQueue() (bool,*Node) {
	if q.IsEmpty() {
		return false, &Node{}
	} else {
		q.len--
		elem := q.Elem[0]
		q.Elem = q.Elem[1:]
		return true, elem
	}
}


type Stack struct {
	Top int
	Elem []*Node
}

func NewStack() *Stack {
	var elem []*Node
	s := &Stack{-1, elem}
	return s
}



func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}
func (s *Stack) push(i *Node) {
	s.Top++
	s.Elem = append(s.Elem, i)
}
func (s *Stack) pop() (bool,*Node) {
	if s.IsEmpty() {
		return false, &Node{}
	} else {
		fmt.Print(s.Top)
		elem := s.Elem[s.Top]
		s.Top--
		return true, elem
	}
}