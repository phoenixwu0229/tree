package main

import (
	"fmt"
	"tree/pkg"
)

func main() {
	elem := []int{1,2,3,4,5,6}
	q := pkg.NewQueue()
	root := &pkg.Node{elem[0], nil, nil}
	q.AddQueue(root)
	for _, e := range elem {
		_, node := q.DeQueue()
		node.Val = e
		fmt.Println(e)
		nodeL := &pkg.Node{0, nil, nil}
		nodeR := &pkg.Node{0, nil, nil}
		node.Left = nodeL
		node.Right = nodeR
		q.AddQueue(nodeL)
		q.AddQueue(nodeR)
	}
	for _, v := range q.Elem {
		fmt.Println(v.Val.(int))
	}
	fmt.Print(q)
	//pkg.PreOrder(root)
}