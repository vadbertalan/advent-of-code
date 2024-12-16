package coordinate

import (
	"aoc/utils/direction"
	"testing"
)

func TestContainsRowColDirCost(t *testing.T) {
	cm := TravelDirCostMap{
		"1-2-U":  5,
		"3-4-R":  10,
		"5-6-DL": 15,
	}

	tests := []struct {
		row      int
		col      int
		dir      direction.Direction
		expected bool
	}{
		{1, 2, direction.Up, true},
		{3, 4, direction.Right, true},
		{5, 6, direction.DownLeft, true},
		{7, 8, direction.Left, false},
		{1, 2, direction.Down, false},
	}

	for _, test := range tests {
		result := cm.ContainsRowColDirCost(test.row, test.col, test.dir)
		if result != test.expected {
			t.Errorf("ContainsRowColDirCost(%d, %d, %v) = %v; want %v", test.row, test.col, test.dir, result, test.expected)
		}
	}
}

func TestContainsCoordDirCost(t *testing.T) {
	cm := TravelDirCostMap{
		"1-2-U": 5,
		"3-4-R": 10,
		"5-6-D": 15,
	}

	tests := []struct {
		coord    Coord
		dir      direction.Direction
		expected bool
	}{
		{Coord{Row: 1, Col: 2}, direction.Up, true},
		{Coord{Row: 3, Col: 4}, direction.Right, true},
		{Coord{Row: 5, Col: 6}, direction.Down, true},
		{Coord{Row: 7, Col: 8}, direction.Left, false},
		{Coord{Row: 1, Col: 2}, direction.Down, false},
	}

	for _, test := range tests {
		result := cm.ContainsCoordDirCost(test.coord, test.dir)
		if result != test.expected {
			t.Errorf("ContainsCoordDirCost(%v, %v) = %v; want %v", test.coord, test.dir, result, test.expected)
		}
	}
}
func TestCoordMapAdd(t *testing.T) {
	cm := make(CoordMap)
	coord := Coord{Row: 1, Col: 2}

	cm.Add(coord)
	if !cm.ContainsCoord(coord) {
		t.Errorf("CoordMap.Add(%v) failed; coord not found in map", coord)
	}
}

func TestTravelMapAdd(t *testing.T) {
	tm := make(TravelMap)
	coord := Coord{Row: 1, Col: 2}
	dir := direction.Up

	tm.Add(coord, dir)
	if !tm.ContainsCoordAndDir(coord, dir) {
		t.Errorf("TravelMap.Add(%v, %v) failed; coord and direction not found in map", coord, dir)
	}
}

func TestTravelDirCostMapAdd(t *testing.T) {
	tdcm := make(TravelDirCostMap)
	coord := Coord{Row: 1, Col: 2}
	dir := direction.Up
	cost := 10

	tdcm.Add(coord, dir, cost)
	if !tdcm.ContainsCoordDirCost(coord, dir) {
		t.Errorf("TravelDirCostMap.Add(%v, %v, %d) failed; coord and direction not found in map", coord, dir, cost)
	}
	if tdcm.Get(coord, dir) != cost {
		t.Errorf("TravelDirCostMap.Add(%v, %v, %d) failed; expected cost %d, got %d", coord, dir, cost, cost, tdcm.Get(coord, dir))
	}
}

func TestRemoveCoordDirCost(t *testing.T) {
	tdcm := make(TravelDirCostMap)
	coord := Coord{Row: 1, Col: 2}
	dir := direction.Up
	cost := 10

	tdcm.Add(coord, dir, cost)
	tdcm.RemoveCoordDirCost(coord, dir, cost)

	if tdcm.ContainsCoordDirCost(coord, dir) {
		t.Errorf("TravelDirCostMap.RemoveCoordDirCost(%v, %v, %d) failed; coord and direction still found in map", coord, dir, cost)
	}
}

func TestTravelDirCostMapGet(t *testing.T) {
	tdcm := make(TravelDirCostMap)
	coord := Coord{Row: 1, Col: 2}
	dir := direction.Up
	cost := 10

	tdcm.Add(coord, dir, cost)

	tests := []struct {
		coord    Coord
		dir      direction.Direction
		expected int
	}{
		{Coord{Row: 1, Col: 2}, direction.Up, 10},
		{Coord{Row: 3, Col: 4}, direction.Right, 0}, // not added, should return 0
	}

	for _, test := range tests {
		result := tdcm.Get(test.coord, test.dir)
		if result != test.expected {
			t.Errorf("TravelDirCostMap.Get(%v, %v) = %d; want %d", test.coord, test.dir, result, test.expected)
		}
	}
}

func TestCoordMapClear(t *testing.T) {
	cm := make(CoordMap)
	cm.Add(Coord{Row: 1, Col: 2})
	cm.Add(Coord{Row: 3, Col: 4})

	cm.Clear()

	if len(cm) != 0 {
		t.Errorf("CoordMap.Clear() failed; expected map to be empty, got %d elements", len(cm))
	}
}

func TestTravelMapClear(t *testing.T) {
	tm := make(TravelMap)
	tm.Add(Coord{Row: 1, Col: 2}, direction.Up)
	tm.Add(Coord{Row: 3, Col: 4}, direction.Right)

	tm.Clear()

	if len(tm) != 0 {
		t.Errorf("TravelMap.Clear() failed; expected map to be empty, got %d elements", len(tm))
	}
}

func TestTravelDirCostMapClear(t *testing.T) {
	tdcm := make(TravelDirCostMap)
	tdcm.Add(Coord{Row: 1, Col: 2}, direction.Up, 10)
	tdcm.Add(Coord{Row: 3, Col: 4}, direction.Right, 20)

	tdcm.Clear()

	if len(tdcm) != 0 {
		t.Errorf("TravelDirCostMap.Clear() failed; expected map to be empty, got %d elements", len(tdcm))
	}
}
