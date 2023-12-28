package graph

import (
	"aoc/utils"
	"fmt"
)

type Node[T any] struct {
	Value    T
	Children []*Node[T]
	Parent   *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
}

func (parentNode *Node[T]) AddChild(value T) *Node[T] {
	newNode := &Node[T]{Value: value}
	parentNode.Children = append(parentNode.Children, newNode)
	newNode.Parent = parentNode
	return newNode
}

func (tree Tree[T]) FindAllPaths() [][]T {
	if tree.Root == nil {
		return [][]T{}
	}

	return tree.Root.FindAllPaths()
}

func (node Node[T]) FindAllPaths() [][]T {
	allPathsRetArrays := [][]T{}

	if len(node.Children) == 0 {
		return append(allPathsRetArrays, []T{node.Value})
	}

	for _, childNode := range node.Children {
		for _, path := range childNode.FindAllPaths() {
			extendedPath := []T{node.Value}
			extendedPath = append(extendedPath, path...)
			allPathsRetArrays = append(allPathsRetArrays, extendedPath)
		}
	}
	return allPathsRetArrays
}

func GetEdgeKey(node1, node2 string) string {
	n1, n2 := orderStrings(node1, node2)
	return fmt.Sprintf("%s-%s", n1, n2)
}

func ParseEdgeKey(edgeKey string) (node1, node2 string) {
	node1, node2 = utils.SplitIn2(edgeKey, "-")
	return node1, node2
}

func orderStrings(str1, str2 string) (string, string) {
	if str1 < str2 {
		return str1, str2
	}
	return str2, str1
}
