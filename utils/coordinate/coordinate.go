package coordinate

import (
	"aoc/utils/direction"
	"fmt"
	"math"
)

type Coord struct {
	Row, Col int
}

func (c Coord) IsEqual(c2 Coord) bool {
	return c.Row == c2.Row && c.Col == c2.Col
}

func (oldCoord Coord) GetNewCoord(dir direction.Direction) Coord {
	m := map[direction.Direction]([2]int){
		direction.Up:    [2]int{-1, 0},
		direction.Right: [2]int{0, 1},
		direction.Down:  [2]int{1, 0},
		direction.Left:  [2]int{0, -1},
	}

	rowAdd := m[dir][0]
	colAdd := m[dir][1]
	return Coord{Row: oldCoord.Row + rowAdd, Col: oldCoord.Col + colAdd}
}

func (c1 Coord) ManhattanDist(c2 Coord) int {
	return int(math.Abs(float64(c1.Row)-float64(c2.Row)) + math.Abs(float64(c1.Col)-float64(c2.Col)))
}

//   ____                    _ __  __
//  / ___|___   ___  _ __ __| |  \/  | __ _ _ __
// | |   / _ \ / _ \| '__/ _` | |\/| |/ _` | '_ \
// | |__| (_) | (_) | | | (_| | |  | | (_| | |_) |
//  \____\___/ \___/|_|  \__,_|_|  |_|\__,_| .__/
//                                         |_|

type CoordMap map[string]bool

func (cm CoordMap) ContainsRowCol(row, col int) bool {
	_, ok := cm[fmt.Sprintf("%d-%d", row, col)]
	return ok
}

func (cm CoordMap) ContainsCoord(c Coord) bool {
	_, ok := cm[fmt.Sprintf("%d-%d", c.Row, c.Col)]
	return ok
}

func (cm *CoordMap) Add(c Coord) {
	(*cm)[fmt.Sprintf("%d-%d", c.Row, c.Col)] = true
}

func (cm *CoordMap) GetAllCoordValues() []Coord {
	coords := []Coord{}
	for k := range *cm {
		var r, c int
		fmt.Sscanf(k, "%d-%d", &r, &c)
		coords = append(coords, Coord{Row: r, Col: c})
	}
	return coords
}

//  ____  _       ___   __  __          _
// |  _ \(_)_ __ / _ \ / _|/ _|___  ___| |_
// | | | | | '__| | | | |_| |_/ __|/ _ \ __|
// | |_| | | |  | |_| |  _|  _\__ \  __/ |_
// |____/|_|_|   \___/|_| |_| |___/\___|\__|

type DirOffset struct {
	Dir       direction.Direction
	RowOffset int
	ColOffset int
}

var dirOffsetsMap map[direction.Direction]DirOffset = map[direction.Direction]DirOffset{
	direction.Up:        {Dir: direction.Up, RowOffset: -1, ColOffset: 0},
	direction.UpRight:   {Dir: direction.UpRight, RowOffset: -1, ColOffset: 1},
	direction.Right:     {Dir: direction.Right, RowOffset: 0, ColOffset: 1},
	direction.RightDown: {Dir: direction.RightDown, RowOffset: 1, ColOffset: 1},
	direction.Down:      {Dir: direction.Down, RowOffset: 1, ColOffset: 0},
	direction.DownLeft:  {Dir: direction.DownLeft, RowOffset: 1, ColOffset: -1},
	direction.Left:      {Dir: direction.Left, RowOffset: 0, ColOffset: -1},
	direction.LeftUp:    {Dir: direction.LeftUp, RowOffset: -1, ColOffset: -1},
}

func GetOffsetForDir(dir direction.Direction) DirOffset {
	return dirOffsetsMap[dir]
}

func GetOffsetsArray(diagonal bool) []DirOffset {
	if diagonal {
		return []DirOffset{
			dirOffsetsMap[direction.Up],
			dirOffsetsMap[direction.UpRight],
			dirOffsetsMap[direction.Right],
			dirOffsetsMap[direction.RightDown],
			dirOffsetsMap[direction.Down],
			dirOffsetsMap[direction.DownLeft],
			dirOffsetsMap[direction.Left],
			dirOffsetsMap[direction.LeftUp],
		}
	}
	return []DirOffset{
		dirOffsetsMap[direction.Up],
		dirOffsetsMap[direction.Right],
		dirOffsetsMap[direction.Down],
		dirOffsetsMap[direction.Left],
	}
}
