package matrix

import (
	"aoc/utils-go/coordinate"
	"aoc/utils-go/direction"
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestCountPathsBetween(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		start  coordinate.Coord
		end    coordinate.Coord
		test   func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool
		want   int
	}{
		{
			name: "simple path",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: 2,
		},
		{
			name: "no path",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: 0,
		},
		{
			name: "multiple paths",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: 12,
		},
		{
			name: "multiple paths 2",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{2, 3, 4},
					{3, 4, 5},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue-currentValue == 1
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.CountPathsBetween(tt.start, tt.end, tt.test); got != tt.want {
				t.Errorf("CountPathsBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPathBetween(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		start  coordinate.Coord
		end    coordinate.Coord
		test   func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool
		want   bool
	}{
		{
			name: "simple path",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: true,
		},
		{
			name: "no path",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: false,
		},
		{
			name: "multiple paths",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue == 1
			},
			want: true,
		},
		{
			name: "path with increasing values",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{2, 3, 4},
					{3, 4, 5},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue-currentValue == 1
			},
			want: true,
		},
		{
			name: "no valid path with increasing values",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{2, 3, 5},
					{3, 4, 6},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			end:   coordinate.Coord{Row: 2, Col: 2},
			test: func(currentCoord, nextCoord coordinate.Coord, currentValue, nextValue int) bool {
				return nextValue-currentValue == 1
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.IsPathBetween(tt.start, tt.end, tt.test); got != tt.want {
				t.Errorf("IsPathBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCoordsWhich(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		test   func(value int) bool
		want   []coordinate.Coord
	}{
		{
			name: "all ones",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(value int) bool {
				return value == 1
			},
			want: []coordinate.Coord{
				{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2},
				{Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 1, Col: 2},
				{Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2},
			},
		},
		{
			name: "no match",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(value int) bool {
				return value == 0
			},
			want: []coordinate.Coord{},
		},
		{
			name: "some match",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(value int) bool {
				return value%2 == 0
			},
			want: []coordinate.Coord{
				{Row: 0, Col: 1}, {Row: 1, Col: 0}, {Row: 1, Col: 2}, {Row: 2, Col: 1},
			},
		},
		{
			name: "single match",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(value int) bool {
				return value == 5
			},
			want: []coordinate.Coord{
				{Row: 1, Col: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.GetAllCoordsWhich(tt.test)
			if len(got) != len(tt.want) {
				t.Errorf("GetAllCoordsWhich() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("GetAllCoordsWhich() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestParseStringMatrix(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  Matrix[string]
	}{
		{
			name:  "empty input",
			lines: []string{},
			want:  Matrix[string]{},
		},
		{
			name: "single line",
			lines: []string{
				"abc",
			},
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
				},
				RowCount:    1,
				ColumnCount: 3,
			},
		},
		{
			name: "multiple lines",
			lines: []string{
				"abc",
				"def",
				"ghi",
			},
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "lines with different lengths",
			lines: []string{
				"abc",
				"de",
				"fghi",
			},
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
					{"d", "e"},
					{"f", "g", "h", "i"},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStringMatrix(tt.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDigitMatrix(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  Matrix[int]
	}{
		{
			name:  "empty input",
			lines: []string{},
			want:  Matrix[int]{},
		},
		{
			name: "single line",
			lines: []string{
				"123",
			},
			want: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
				},
				RowCount:    1,
				ColumnCount: 3,
			},
		},
		{
			name: "multiple lines",
			lines: []string{
				"123",
				"456",
				"789",
			},
			want: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "lines with different lengths",
			lines: []string{
				"123",
				"45",
				"6789",
			},
			want: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5},
					{6, 7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDigitMatrix(tt.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDigitMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstValidNeighbor(t *testing.T) {
	tests := []struct {
		name     string
		matrix   Matrix[int]
		start    coordinate.Coord
		test     func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool
		diagonal bool
		want     *coordinate.Coord
	}{
		{
			name: "valid neighbor exists",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value == 6
			},
			diagonal: false,
			want:     &coordinate.Coord{Row: 1, Col: 2},
		},
		{
			name: "no valid neighbor",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value == 10
			},
			diagonal: false,
			want:     nil,
		},
		{
			name: "valid diagonal neighbor exists",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value == 9
			},
			diagonal: true,
			want:     &coordinate.Coord{Row: 2, Col: 2},
		},
		{
			name: "valid neighbor at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value == 2
			},
			diagonal: false,
			want:     &coordinate.Coord{Row: 0, Col: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.GetFirstValidNeighbor(tt.start, tt.test, tt.diagonal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstValidNeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValidNeighborCoords(t *testing.T) {
	tests := []struct {
		name     string
		matrix   Matrix[int]
		start    coordinate.Coord
		test     func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool
		diagonal bool
		want     []coordinate.Coord
	}{
		{
			name: "valid neighbors exist",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value%2 == 0
			},
			diagonal: false,
			want: []coordinate.Coord{
				{Row: 0, Col: 1}, {Row: 1, Col: 2}, {Row: 2, Col: 1}, {Row: 1, Col: 0},
			},
		},
		{
			name: "no valid neighbors",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value == 10
			},
			diagonal: false,
			want:     []coordinate.Coord{},
		},
		{
			name: "valid diagonal neighbors exist",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 1, Col: 1},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value%2 == 0
			},
			diagonal: true,
			want: []coordinate.Coord{
				{Row: 0, Col: 1}, {Row: 1, Col: 2}, {Row: 2, Col: 1}, {Row: 1, Col: 0},
			},
		},
		{
			name: "valid neighbors at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value%2 == 0
			},
			diagonal: false,
			want: []coordinate.Coord{
				{Row: 0, Col: 1}, {Row: 1, Col: 0},
			},
		},
		{
			name: "valid diagonal neighbors at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 6, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			start: coordinate.Coord{Row: 0, Col: 0},
			test: func(value int, neighborCoord coordinate.Coord, dir direction.Direction) bool {
				return value%2 == 0
			},
			diagonal: true,
			want: []coordinate.Coord{
				{Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 1, Col: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrix.GetValidNeighborCoords(tt.start, tt.test, tt.diagonal)
			if !(len(got) == 0 && len(tt.want) == 0) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidNeighborCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCount(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		value  int
		want   int
	}{
		{
			name: "count ones",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			value: 1,
			want:  8,
		},
		{
			name: "count zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 1, 1},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			value: 0,
			want:  4,
		},
		{
			name: "count twos",
			matrix: Matrix[int]{
				Values: [][]int{
					{2, 2, 2},
					{2, 2, 2},
					{2, 2, 2},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			value: 2,
			want:  9,
		},
		{
			name: "count non-existent value",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			value: 10,
			want:  0,
		},
		{
			name: "count mixed values",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			value: 3,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Count(tt.value); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		name      string
		matrix    Matrix[int]
		coord     coordinate.Coord
		value     int
		wantPanic bool
		want      Matrix[int]
	}{
		{
			name: "valid set",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 1, Col: 1},
			value: 10,
			want: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 10, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "set at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 0, Col: 0},
			value: 10,
			want: Matrix[int]{
				Values: [][]int{
					{10, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "invalid set out of bounds",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord:     coordinate.Coord{Row: 3, Col: 3},
			value:     10,
			wantPanic: true,
		},
		{
			name: "invalid set negative index",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord:     coordinate.Coord{Row: -1, Col: -1},
			value:     10,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("Set() panicked unexpectedly: %v", r)
					}
				} else {
					if tt.wantPanic {
						t.Errorf("Set() did not panic")
					}
				}
			}()

			tt.matrix.Set(tt.coord, tt.value)
			if !tt.wantPanic && !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("Set() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}

func TestAt(t *testing.T) {
	tests := []struct {
		name      string
		matrix    Matrix[int]
		coord     coordinate.Coord
		want      int
		wantPanic bool
	}{
		{
			name: "valid coordinate",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 1, Col: 1},
			want:  5,
		},
		{
			name: "coordinate at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 0, Col: 0},
			want:  1,
		},
		{
			name: "invalid coordinate out of bounds",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord:     coordinate.Coord{Row: 3, Col: 3},
			wantPanic: true,
		},
		{
			name: "invalid coordinate negative index",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord:     coordinate.Coord{Row: -1, Col: -1},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantPanic {
						t.Errorf("At() panicked unexpectedly: %v", r)
					}
				} else {
					if tt.wantPanic {
						t.Errorf("At() did not panic")
					}
				}
			}()

			got := tt.matrix.At(tt.coord)
			if !tt.wantPanic && got != tt.want {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidCoord(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		coord  coordinate.Coord
		want   bool
	}{
		{
			name: "valid coordinate",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 1, Col: 1},
			want:  true,
		},
		{
			name: "coordinate at edge",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 0, Col: 0},
			want:  true,
		},
		{
			name: "invalid coordinate out of bounds",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: 3, Col: 3},
			want:  false,
		},
		{
			name: "invalid coordinate negative index",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			coord: coordinate.Coord{Row: -1, Col: -1},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.IsValidCoord(tt.coord); got != tt.want {
				t.Errorf("IsValidCoord() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestClone(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.matrix.Clone()
			if !reflect.DeepEqual(clone, tt.matrix) {
				t.Errorf("Clone() = %v, want %v", clone, tt.matrix)
			}

			// Modify the original matrix and ensure the clone does not change
			if tt.matrix.RowCount > 0 && tt.matrix.ColumnCount > 0 {
				tt.matrix.Values[0][0] = -1
				if reflect.DeepEqual(clone, tt.matrix) {
					t.Errorf("Clone() changed after modifying the original matrix")
				}
			}
		})
	}
}

func TestPrintWithSmallSpacing(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: " 1 2 3\n 4 5 6\n 7 8 9\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: " 0 0 0\n 0 0 0\n 0 0 0\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: " 1 2\n 3 4\n 5 6\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.PrintWithSmallSpacing()
			})
			if got != tt.want {
				t.Errorf("PrintWithSmallSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPrintWithSpacing(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "    1    2    3\n    4    5    6\n    7    8    9\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "    0    0    0\n    0    0    0\n    0    0    0\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: "    1    2\n    3    4\n    5    6\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.PrintWithSpacing()
			})
			if got != tt.want {
				t.Errorf("PrintWithSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "123\n456\n789\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "000\n000\n000\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: "12\n34\n56\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.Print()
			})
			if got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintln(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "123\n456\n789\n\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "000\n000\n000\n\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "\n",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: "12\n34\n56\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.Println()
			})
			if got != tt.want {
				t.Errorf("Println() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintlnWithSpacing(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "    1    2    3\n    4    5    6\n    7    8    9\n\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: "    0    0    0\n    0    0    0\n    0    0    0\n\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "\n",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: "    1    2\n    3    4\n    5    6\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.PrintlnWithSpacing()
			})
			if got != tt.want {
				t.Errorf("PrintlnWithSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintlnWithSmallSpacing(t *testing.T) {
	tests := []struct {
		name   string
		matrix Matrix[int]
		want   string
	}{
		{
			name: "simple matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: " 1 2 3\n 4 5 6\n 7 8 9\n\n",
		},
		{
			name: "matrix with zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			want: " 0 0 0\n 0 0 0\n 0 0 0\n\n",
		},
		{
			name: "empty matrix",
			matrix: Matrix[int]{
				Values:      [][]int{},
				RowCount:    0,
				ColumnCount: 0,
			},
			want: "\n",
		},
		{
			name: "non-square matrix",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
				RowCount:    3,
				ColumnCount: 2,
			},
			want: " 1 2\n 3 4\n 5 6\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.PrintlnWithSmallSpacing()
			})
			if got != tt.want {
				t.Errorf("PrintlnWithSmallSpacing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintlnWithOverride(t *testing.T) {
	tests := []struct {
		name     string
		matrix   Matrix[int]
		test     func(i, j int, value int) bool
		override string
		want     string
	}{
		{
			name: "override zeros",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 0, 3},
					{4, 0, 6},
					{7, 8, 0},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(i, j int, value int) bool {
				return value == 0
			},
			override: "X",
			want:     " 1 X 3\n 4 X 6\n 7 8 X\n",
		},
		{
			name: "override even numbers",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(i, j int, value int) bool {
				return value%2 == 0
			},
			override: "E",
			want:     " 1 E 3\n E 5 E\n 7 E 9\n",
		},
		{
			name: "override all",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(i, j int, value int) bool {
				return true
			},
			override: "A",
			want:     " A A A\n A A A\n A A A\n",
		},
		{
			name: "override none",
			matrix: Matrix[int]{
				Values: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			test: func(i, j int, value int) bool {
				return false
			},
			override: "N",
			want:     " 1 2 3\n 4 5 6\n 7 8 9\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureOutput(func() {
				tt.matrix.PrintlnWithOverride(tt.test, tt.override)
			})
			if got != tt.want {
				t.Errorf("PrintlnWithOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStringMatrixAndGetStartingPoint(t *testing.T) {
	tests := []struct {
		name      string
		lines     []string
		test      func(value string) bool
		want      Matrix[string]
		wantCoord *coordinate.Coord
	}{
		{
			name:      "empty input",
			lines:     []string{},
			test:      func(value string) bool { return value == "S" },
			want:      Matrix[string]{},
			wantCoord: nil,
		},
		{
			name: "single line with starting point",
			lines: []string{
				"abcS",
			},
			test: func(value string) bool { return value == "S" },
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c", "S"},
				},
				RowCount:    1,
				ColumnCount: 4,
			},
			wantCoord: &coordinate.Coord{Row: 0, Col: 3},
		},
		{
			name: "multiple lines with starting point",
			lines: []string{
				"abc",
				"def",
				"gSh",
			},
			test: func(value string) bool { return value == "S" },
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "S", "h"},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			wantCoord: &coordinate.Coord{Row: 2, Col: 1},
		},
		{
			name: "lines with different lengths and starting point",
			lines: []string{
				"abc",
				"deS",
				"fghi",
			},
			test: func(value string) bool { return value == "S" },
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
					{"d", "e", "S"},
					{"f", "g", "h", "i"},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			wantCoord: &coordinate.Coord{Row: 1, Col: 2},
		},
		{
			name: "no starting point",
			lines: []string{
				"abc",
				"def",
				"ghi",
			},
			test: func(value string) bool { return value == "S" },
			want: Matrix[string]{
				Values: [][]string{
					{"a", "b", "c"},
					{"d", "e", "f"},
					{"g", "h", "i"},
				},
				RowCount:    3,
				ColumnCount: 3,
			},
			wantCoord: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotCoord := ParseStringMatrixAndGetStartingPoint(tt.lines, tt.test)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringMatrixAndGetStartingPoint() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotCoord, tt.wantCoord) {
				t.Errorf("ParseStringMatrixAndGetStartingPoint() gotCoord = %v, want %v", gotCoord, tt.wantCoord)
			}
		})
	}
}

// captureOutput captures the output of a function that writes to stdout.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
