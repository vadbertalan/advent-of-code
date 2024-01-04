package graph

import (
	"aoc/utils/collections"
	"fmt"
)

type EdgeKey = string

type GraphMethods[T any] interface {
	Println()
	AddNode(name string, value T)
	AddEdge(node1Name, node2Name string)
	GetEdges() [2]string
	GetIncomingNodesOf(node string) []string
	RemoveEdge(edgeKey EdgeKey)
	CountComponents(startNode string) int
	TraverseBFS(startNode string) (path []string, edges []EdgeKey)
}

type Graph[T any] struct {
	Neighbors map[string][]string
	Values    map[string]T
}

func (g *Graph[T]) Println() {
	fmt.Println(g.Neighbors)
	fmt.Println(g.Values)
}

// Idempotent: does not insert same node twice
func (g *Graph[T]) AddNode(node string, value *T) {
	_, ok := g.Neighbors[node]
	if !ok {
		g.Neighbors[node] = []string{}
		if value != nil {
			g.Values[node] = *value
		}
	}
}

func (g *Graph[T]) GetEdges() (edges [][2]string) {
	edgeSet := collections.NewSet[string]()
	for node, neighbors := range g.Neighbors {
		for _, neighborNode := range neighbors {
			key1 := fmt.Sprintf("%s-%s", node, neighborNode)
			key2 := fmt.Sprintf("%s-%s", neighborNode, node)
			if !edgeSet.Has(key1) && !edgeSet.Has(key2) {
				edgeSet.Add(key1)
			}
		}
	}
	for _, edgeKey := range edgeSet.GetValues() {
		node1, node2 := ParseEdgeKey(edgeKey)
		edges = append(edges, [2]string{node1, node2})
	}
	return edges
}

func (g *Graph[T]) CountComponents(startNode string) int {
	componentCount := 0

	nodeMap := map[string]int{}

	visitedNodes := 0

	q := collections.NewQueue[string]()

	for visitedNodes != len(g.Neighbors) {
		q.Append(startNode)

		// Traverse current component of startNode
		for !q.IsEmpty() {
			node := q.Pop()

			// Nodes can be inserted into the q many times, so
			// checking their visited status has to be done twice
			_, didVisitSinceGotInQueue := nodeMap[node]
			if didVisitSinceGotInQueue {
				continue
			}

			visitedNodes++
			nodeMap[node] = componentCount

			for _, n := range g.Neighbors[node] {
				_, didVisit := nodeMap[n]
				if !didVisit {
					q.Append(n)
				}
			}
		}

		// Get new start node
		for node := range g.Neighbors {
			_, didVisit := nodeMap[node]
			if !didVisit {
				startNode = node
				continue
			}
		}

		componentCount++
	}

	return componentCount
}

func (g *Graph[T]) TraverseBFS(startNode string) (path []string, edges []string) {
	nodeMap := map[string]bool{}

	type tf struct {
		to, from string
	}
	q := collections.NewQueue[tf]()
	q.Append(tf{startNode, ""})

	for !q.IsEmpty() {
		toFrom := q.Pop()
		toNode := toFrom.to
		fromNode := toFrom.from

		// Nodes can be inserted into the q many times, so
		// checking their visited status has to be done twice
		_, didVisitSinceGotInQueue := nodeMap[toNode]
		if didVisitSinceGotInQueue {
			continue
		}

		if toNode != startNode && fromNode != "" {
			edges = append(edges, GetEdgeKey(toNode, fromNode))
		}
		path = append(path, toNode)
		nodeMap[toNode] = true

		for _, newNode := range g.Neighbors[toNode] {
			_, didVisit := nodeMap[newNode]
			if !didVisit {
				q.Append(tf{to: newNode, from: toNode})
			}
		}
	}
	return path, edges
}
