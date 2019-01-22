package algorithms

import (
	ds "demo/data_structures"
	"fmt"
)

func Bfs(graph *ds.MyGraph, src *ds.MyNode) {

	vis := make(map[*ds.MyNode]bool)
	for _,node := range graph.GetNodes() {
		vis[node] = false
	}

	q := ds.MyQueue{}
	q.NewMyQueue()
	q.Push(src.Data)

	for !q.IsEmpty() {

		val, _ := q.Pop()

		var n *ds.MyNode
		flag := false
		for _,node := range graph.GetNodes() {
			if node.Data == val {
				if vis[node] == true {
					flag = true
					break;
				} else {
					n = node
					break
				}
			}
		}
		if flag == true {
			continue
		}

		vis[n] = true
		fmt.Print(val, " ")

		for _,nodes := range graph.GetEdges()[n] {

			flag = false
			var neigh *ds.MyNode = nil
			for _,node := range graph.GetNodes() {
				if nodes.Data == node.Data {
					if vis[node] == true {
						flag = true
						break
					} else {
						neigh = node
						break
					}
				}
			}

			if flag == true {
				continue
			}

			if neigh != nil {
				q.Push(neigh.Data)
			}
		}
	}
}