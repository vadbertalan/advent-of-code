package graph

type Node[T any] struct {
	Value    T
	Children []*Node[T]
	Parent   *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
}

func (parentNode *Node[T]) addChild(value T) *Node[T] {
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
