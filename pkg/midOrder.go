package pkg

import "fmt"

func MidOrder(root *Node) {
	s := NewStack()
	p := root
	s.push(p)
	p = p.Left
	for p!=nil||!s.IsEmpty() {
		if p!=nil {
			s.push(p)
			p = p.Left
		} else {
			_,p = s.pop()
			fmt.Print(p.Val.(string))
			p=p.Right
		}
	}
}