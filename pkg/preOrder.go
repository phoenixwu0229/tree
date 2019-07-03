package pkg

import "fmt"

func PreOrder(root *Node) {
	s := NewStack()
	p := root
	s.push(root)
	for p != nil || !s.IsEmpty() {
		_, p = s.pop()
		fmt.Print(p.Val.(int))
		if p.Right != nil {
			s.push(p.Right)
			p = p.Right
		} else if p.Left != nil{
			s.push(p.Left)
			p = p.Left
		}
	}
}