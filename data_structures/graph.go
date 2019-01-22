package ds

import (
	"fmt"
)

type MyNode struct {
	Data int
}

type MyGraph struct {
	nodes []*MyNode
	edges map[MyNode][]*MyNode
}

func (graph *MyGraph) Size() int {
	return len(graph.nodes)
}

func (graph *MyGraph) AddNode(node *MyNode) {
	graph.nodes = append(graph.nodes, node)
}

func (graph *MyGraph) GetNodes() []*MyNode {
	return graph.nodes
}

func (graph *MyGraph) AddEdge(src, dest *MyNode) {
	if(graph.edges == nil) {
		graph.edges = make(map[MyNode][]*MyNode)
	}
	graph.edges[*src] = append(graph.edges[*src], dest)
}

func (graph *MyGraph) GetEdges() map[MyNode][]*MyNode {
	return graph.edges
}

func (graph *MyGraph) PrintGraph() {
	fmt.Println("Printing the Graph:")
	i := 0
	for ;i<len(graph.nodes);i++ {
		fmt.Print(graph.nodes[i].Data, ": ")
		j := 0
		lt := graph.edges[*graph.nodes[i]]
		for ;j<len(lt);j++ {
			fmt.Print(lt[j].Data, " ")
		}
		fmt.Println()
	}
}
