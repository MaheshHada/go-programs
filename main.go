package main

import (
	al "demo/algorithms"
	ds "demo/data_structures"
	"fmt"
)

func main() {

	//s  :=  ds.MyStack{}
	//s.NewMyStack()
	//
	//s.Push(10)
	//s.Push(20)
	//val,err := s.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//val,err = s.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//val,err = s.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//
	//q := ds.MyQueue{}
	//q.NewMyQueue()
	//
	//q.Push(11)
	//q.Push(12)
	//val,err = q.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//val,err = q.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}
	//
	//val,err = q.Pop()
	//if err == nil {
	//	fmt.Println(*val)
	//} else {
	//	fmt.Println(err)
	//}

	g := ds.MyGraph{}
	n1 := ds.MyNode{Data:0}
	n2 := ds.MyNode{Data:1}
	n3 := ds.MyNode{Data:2}
	n4 := ds.MyNode{Data:3}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n3, &n2)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n1, &n3)
	g.PrintGraph()

	fmt.Print("\nDFS: ")
	al.Dfs(g, n1)
	fmt.Print("\nBFS: ")
	al.Bfs(g, n1)
	fmt.Println()
}