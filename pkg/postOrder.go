package pkg

import "fmt"

func PostOrder(root *Node) {
	s := NewStack()
	curP := root
	preVisit := &Node{}
	for curP.Left != nil {
		s.push(curP)
		curP = curP.Left
	}
	for curP != nil || !s.IsEmpty() {
		if preVisit == nil || preVisit == curP.Right {
			_,p := s.pop()
			fmt.Print(p.Val.(string))
			preVisit = p
		}
		if curP.Right != nil {
			s.push(curP.Right)
			preVisit = curP.Right
		} else if curP.Left != nil {
			for curP.Left != nil {
				curP = curP.Left
				s.push(curP)
				preVisit = nil
			}
		}
	}
}