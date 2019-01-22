package algorithms

import (
	ds "demo/data_structures"
	"fmt"
)

func DfsRecursive(graph *ds.MyGraph, node *ds.MyNode, vis map[*ds.MyNode]bool) {

	if vis[node] == true {
		return
	}

	fmt.Print(node.Data, " ")
	vis[node] = true

	edges := graph.GetEdges()
	lt := edges[node]
	for j := 0;j<len(lt);j++ {
		DfsRecursive(graph, lt[j], vis)
	}
}

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
				st.Push(neigh.Data)
			}
		}

	}
}