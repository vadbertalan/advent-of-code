package graph

import (
	"aoc/utils"
	"fmt"
)

type DGraph[T any] struct {
	Graph[T]
}

func NewDGraph[T any]() *DGraph[T] {
	g := &DGraph[T]{}
	g.Neighbors = make(map[string][]string)
	g.Values = make(map[string]T)
	return g
}

func (g *DGraph[T]) AddEdge(node1, node2 string) {
	_, ok := g.Neighbors[node1]
	if !ok {
		panic(fmt.Sprintf("Node %s does not exist in graph", node1))
	}

	_, ok = g.Neighbors[node2]
	if !ok {
		panic(fmt.Sprintf("Node %s does not exist in graph", node2))
	}

	g.Neighbors[node1] = append(g.Neighbors[node1], node2)
}

func (g *DGraph[T]) RemoveEdge(edgeKey string) {
	node1, node2 := ParseEdgeKey(edgeKey)
	g.Neighbors[node1] = utils.RemoveItemFromArray(g.Neighbors[node1], node2)
}

func (g *DGraph[T]) GetIncomingNodesOf(node string) (incomingNodes []string) {
	for fromNode, toNodes := range g.Neighbors {
		for _, toNode := range toNodes {
			if toNode == node {
				incomingNodes = append(incomingNodes, fromNode)

				// Assuming there's max one edge from one node to another
				break
			}
		}
	}
	return incomingNodes
}
