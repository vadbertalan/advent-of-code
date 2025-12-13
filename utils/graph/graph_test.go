package graph

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to create a graph for testing
func setupTestGraphForPaths() *Graph[int] {
	g := &Graph[int]{
		Neighbors: make(map[string][]string),
		Values:    make(map[string]int),
	}
	g.AddNode("A", nil)
	g.AddNode("B", nil)
	g.AddNode("C", nil)
	g.AddNode("D", nil)
	g.AddNode("E", nil)
	g.AddNode("F", nil)

	g.Neighbors["A"] = []string{"B", "C"}
	g.Neighbors["B"] = []string{"D"}
	g.Neighbors["C"] = []string{"D", "E"}
	g.Neighbors["D"] = []string{"F"}
	g.Neighbors["E"] = []string{"F"}
	return g
}

// Helper for sorting paths for comparison
func sortPaths(paths [][]string) {
	sort.Slice(paths, func(i, j int) bool {
		path1 := paths[i]
		path2 := paths[j]
		for k := 0; k < len(path1) && k < len(path2); k++ {
			if path1[k] != path2[k] {
				return path1[k] < path2[k]
			}
		}
		return len(path1) < len(path2)
	})
}

func TestGetAllPaths(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph[int]
		from          string
		to            string
		expectedPaths [][]string
	}{
		{
			name:  "simple paths",
			graph: setupTestGraphForPaths(),
			from:  "A",
			to:    "D",
			expectedPaths: [][]string{
				{"A", "B", "D"},
				{"A", "C", "D"},
			},
		},
		{
			name:  "longer paths",
			graph: setupTestGraphForPaths(),
			from:  "A",
			to:    "F",
			expectedPaths: [][]string{
				{"A", "B", "D", "F"},
				{"A", "C", "D", "F"},
				{"A", "C", "E", "F"},
			},
		},
		{
			name:  "no path",
			graph: setupTestGraphForPaths(),
			from:  "A",
			to:    "A", // Path to self is defined by the function as returning {{to}} when from == to
			expectedPaths: [][]string{
				{"A"},
			},
		},
		{
			name:          "no path to unreachable node",
			graph:         setupTestGraphForPaths(),
			from:          "F",
			to:            "A",
			expectedPaths: nil, // F has no outgoing edges, so no path to A.
		},
		{
			name:  "path to an immediate neighbor",
			graph: setupTestGraphForPaths(),
			from:  "A",
			to:    "B",
			expectedPaths: [][]string{
				{"A", "B"},
			},
		},
		{
			name: "graph with a cycle (should not loop due to 'seen' set)",
			graph: func() *Graph[int] {
				g := &Graph[int]{
					Neighbors: make(map[string][]string),
					Values:    make(map[string]int),
				}
				g.AddNode("X", nil)
				g.AddNode("Y", nil)
				g.AddNode("Z", nil)
				g.Neighbors["X"] = []string{"Y"}
				g.Neighbors["Y"] = []string{"X", "Z"} // Cycle X-Y
				g.Neighbors["Z"] = []string{}
				return g
			}(),
			from: "X",
			to:   "Z",
			expectedPaths: [][]string{
				{"X", "Y", "Z"},
			},
		},
		{
			name: "single node graph, path to self",
			graph: func() *Graph[int] {
				g := &Graph[int]{
					Neighbors: make(map[string][]string),
					Values:    make(map[string]int),
				}
				g.AddNode("S", nil)
				return g
			}(),
			from:          "S",
			to:            "S",
			expectedPaths: [][]string{{"S"}},
		},
		{
			name: "empty graph",
			graph: func() *Graph[int] {
				return &Graph[int]{
					Neighbors: make(map[string][]string),
					Values:    make(map[string]int),
				}
			}(),
			from:          "A",
			to:            "B",
			expectedPaths: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualPaths := tt.graph.GetAllPaths(tt.from, tt.to)

			// Sort both actual and expected paths for order-independent comparison
			sortPaths(actualPaths)
			sortPaths(tt.expectedPaths)

			if !reflect.DeepEqual(actualPaths, tt.expectedPaths) {
				t.Errorf("GetAllPaths(%s, %s) got %v, want %v", tt.from, tt.to, actualPaths, tt.expectedPaths)
			}
		})
	}
}

func TestCountComponents_SingleComponent(t *testing.T) {
	g := &Graph[int]{
		Neighbors: map[string][]string{
			"a": {"b"},
			"b": {"a", "c"},
			"c": {"b"},
		},
		Values: map[string]int{},
	}

	count, components := g.CountComponents("a")

	if count != 1 {
		t.Errorf("CountComponents() count = %v; want 1", count)
	}
	if len(components) != 1 {
		t.Errorf("CountComponents() components length = %v; want 1", len(components))
	}
	if len(components[0]) != 3 {
		t.Errorf("CountComponents() component[0] length = %v; want 3", len(components[0]))
	}
}

func TestCountComponents_MultipleComponents(t *testing.T) {
	g := &Graph[int]{
		Neighbors: map[string][]string{
			"a": {"b"},
			"b": {"a"},
			"c": {"d"},
			"d": {"c"},
			"e": {},
		},
		Values: map[string]int{},
	}

	count, components := g.CountComponents("a")

	if count != 3 {
		t.Errorf("CountComponents() count = %v; want 3", count)
	}
	if len(components) != 3 {
		t.Errorf("CountComponents() components length = %v; want 3", len(components))
	}
	if len(components[0]) != 2 {
		t.Errorf("CountComponents() component[0] length = %v; want 2", len(components[0]))
	}
	if len(components[1]) != 2 {
		t.Errorf("CountComponents() component[1] length = %v; want 2", len(components[1]))
	}
	if len(components[2]) != 1 {
		t.Errorf("CountComponents() component[2] length = %v; want 1", len(components[2]))
	}
}

func TestCountComponents_SingleNode(t *testing.T) {
	g := &Graph[int]{
		Neighbors: map[string][]string{
			"a": {},
		},
		Values: map[string]int{},
	}

	count, components := g.CountComponents("a")

	if count != 1 {
		t.Errorf("CountComponents() count = %v; want 1", count)
	}
	if len(components[0]) != 1 {
		t.Errorf("CountComponents() component[0] length = %v; want 1", len(components[0]))
	}
}
