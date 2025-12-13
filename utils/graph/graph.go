package graph

import (
	"aoc/utils/collections"
	"fmt"
	"slices"
)

type EdgeKey = string

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

func (g *Graph[T]) GetNodes() []string {
	nodes := []string{}
	for node := range g.Neighbors {
		nodes = append(nodes, node)
	}
	return nodes
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

func (g *Graph[T]) CountComponents(startNode string) (int, [][]string) {
	componentsCount := 0
	nodeMap := map[string]int{}
	visitedNodes := 0
	components := [][]string{}

	q := collections.NewQueue[string]()

	for visitedNodes != len(g.Neighbors) {
		q.Append(startNode)
		currentComponent := []string{}

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
			nodeMap[node] = componentsCount
			currentComponent = append(currentComponent, node)

			for _, n := range g.Neighbors[node] {
				_, didVisit := nodeMap[n]
				if !didVisit {
					q.Append(n)
				}
			}
		}

		components = append(components, currentComponent)

		// Get new start node
		for node := range g.Neighbors {
			_, didVisit := nodeMap[node]
			if !didVisit {
				startNode = node
				break
			}
		}

		componentsCount++
	}

	return componentsCount, components
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

func (g *Graph[T]) getAllPaths(from, to string, seen collections.Set[string]) (paths [][]string) {
	if from == to {
		return [][]string{{to}}
	}

	if seen.Has(from) {
		return nil
	}

	// Create a new set for the current path to avoid modifying the parent's 'seen' set
	// for other branches.
	currentPathSeen := seen.Copy()
	currentPathSeen.Add(from) // Add the current node to this path's seen set

	for _, nb := range g.Neighbors[from] {
		// Recursively find paths from neighbor to 'to', passing the updated 'seen' set
		for _, path := range g.getAllPaths(nb, to, *currentPathSeen) {
			paths = append(paths, slices.Concat([]string{from}, path))
		}
	}
	return paths
}

func (g *Graph[T]) GetAllPaths(from, to string) (paths [][]string) {
	return g.getAllPaths(from, to, *collections.NewSet[string]())
}
