package algorithms

import (
	ds "go-programs/data_structures"
	"fmt"
)

func Dfs(graph *ds.MyGraph, src *ds.MyNode) {

	vis := make(map[*ds.MyNode]bool)
	for _,node := range graph.GetNodes() {
		vis[node] = false;
	}

	st := ds.MyStack{}
	st.NewMyStack()
	st.Push(src.Data)

	for !st.IsEmpty() {

		val,_ := st.Pop()

		var n *ds.MyNode
		flag := false
		for _,node := range graph.GetNodes() {
			if node.Data == val {
				if vis[node] == true {
					flag = true
				} else {
					n = node
				}
				break
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
					} else {
						neigh = node
					}
					break
				}
			}
			if flag == true {
				continue
			}

			if neigh != nil {
				st.Push(neigh.Data)
			}
		}

	}
}