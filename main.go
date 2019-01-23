package main

import (
	al "go-programs/algorithms"
	ds "go-programs/data_structures"
	"fmt"
)

func main() {

	g := ds.CreateNewGraph()

	n1 := &ds.MyNode{Data:0}
	n2 := &ds.MyNode{Data:1}
	n3 := &ds.MyNode{Data:2}
	n4 := &ds.MyNode{Data:3}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)

	g.AddEdge(n1, n2)
	g.AddEdge(n3, n2)
	g.AddEdge(n2, n4)
	g.AddEdge(n1, n3)
	g.PrintGraph()

	fmt.Print("\nDFS: ")
	al.Dfs(g, n1)
	fmt.Print("\nBFS: ")
	al.Bfs(g, n1)
	fmt.Println()
}