package graph

import (
	"aoc/utils"
	"aoc/utils/collections"
	"fmt"
)

type UGraph struct {
	Neighbors map[string][]string
}

func NewUGraph() *UGraph {
	g := &UGraph{}
	g.Neighbors = make(map[string][]string)
	return g
}

func (g *UGraph) Println() {
	fmt.Println(g.Neighbors)
}

// Idempotent: does not insert same node twice
func (g *UGraph) AddNode(node string) {
	_, ok := g.Neighbors[node]
	if !ok {
		g.Neighbors[node] = []string{}
	}
}

func (g *UGraph) AddEdge(node1, node2 string) {
	_, ok := g.Neighbors[node1]
	if !ok {
		g.AddNode(node1)
	}

	_, ok = g.Neighbors[node2]
	if !ok {
		g.AddNode(node2)
	}

	g.Neighbors[node1] = append(g.Neighbors[node1], node2)
	g.Neighbors[node2] = append(g.Neighbors[node2], node1)
}

func (g *UGraph) GetEdges() (edges [][2]string) {
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

func (g *UGraph) RemoveEdge(edgeKey string) {
	node1, node2 := ParseEdgeKey(edgeKey)
	g.Neighbors[node1] = utils.RemoveItemFromArray(g.Neighbors[node1], node2)
	g.Neighbors[node2] = utils.RemoveItemFromArray(g.Neighbors[node2], node1)
}

func (g *UGraph) CountComponents(startNode string) int {
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

func (g *UGraph) TraverseBFS(startNode string) (path []string, edges []string) {
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
