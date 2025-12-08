package graph

import (
	"testing"
)

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
