package graph

import (
	"aoc/utils"
	"fmt"
)

type UGraph[T any] struct {
	Graph[T]
}

func NewUGraph[T any]() *UGraph[T] {
	g := &UGraph[T]{}
	g.Neighbors = make(map[string][]string)
	g.Values = make(map[string]T)
	return g
}

func (g *UGraph[T]) AddEdge(node1, node2 string) {
	_, ok := g.Neighbors[node1]
	if !ok {
		panic(fmt.Sprintf("Node %s does not exist in graph", node1))
	}

	_, ok = g.Neighbors[node2]
	if !ok {
		panic(fmt.Sprintf("Node %s does not exist in graph", node2))
	}

	g.Neighbors[node1] = append(g.Neighbors[node1], node2)
	g.Neighbors[node2] = append(g.Neighbors[node2], node1)
}

func (g *UGraph[T]) RemoveEdge(edgeKey string) {
	node1, node2 := ParseEdgeKey(edgeKey)
	g.Neighbors[node1] = utils.RemoveItemFromArray(g.Neighbors[node1], node2)
	g.Neighbors[node2] = utils.RemoveItemFromArray(g.Neighbors[node2], node1)
}

func (g *UGraph[T]) GetIncomingNodesOf(node string) (incomingNodes []string) {
	return g.Neighbors[node]
}
