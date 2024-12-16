package coordinate

import (
	"aoc/utils"
	"aoc/utils/direction"
	"fmt"
	"math"
	"strconv"
)

type Coord struct {
	Row, Col int
}

func (c Coord) IsEqual(c2 Coord) bool {
	return c.Row == c2.Row && c.Col == c2.Col
}

func (oldCoord Coord) GetNewCoord(dir direction.Direction) Coord {
	m := map[direction.Direction]([2]int){
		direction.Up:        [2]int{-1, 0},
		direction.UpRight:   [2]int{-1, 1},
		direction.Right:     [2]int{0, 1},
		direction.RightDown: [2]int{1, 1},
		direction.Down:      [2]int{1, 0},
		direction.DownLeft:  [2]int{1, -1},
		direction.Left:      [2]int{0, -1},
		direction.LeftUp:    [2]int{-1, -1},
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

func (cm *CoordMap) RemoveCoord(c Coord) {
	delete(*cm, fmt.Sprintf("%d-%d", c.Row, c.Col))
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

func (cm *CoordMap) Clear() {
	*cm = make(CoordMap)
}

func (cm CoordMap) Copy() CoordMap {
	newMap := make(CoordMap)
	for k := range cm {
		newMap[k] = true
	}
	return newMap
}

//   _______                  _ __  __
//  |__   __|                | |  \/  |
//     | |_ __ __ ___   _____| | \  / | __ _ _ __
//     | | '__/ _` \ \ / / _ \ | |\/| |/ _` | '_ \
//     | | | | (_| |\ V /  __/ | |  | | (_| | |_) |
//     |_|_|  \__,_| \_/ \___|_|_|  |_|\__,_| .__/
//                                          | |
//                                          |_|

type TravelMap map[string]bool

func getDirString(d direction.Direction) string {
	switch d {
	case direction.Up:
		return "U"
	case direction.UpRight:
		return "UR"
	case direction.Right:
		return "R"
	case direction.RightDown:
		return "RD"
	case direction.Down:
		return "D"
	case direction.DownLeft:
		return "DL"
	case direction.Left:
		return "L"
	case direction.LeftUp:
		return "LU"
	}
	panic(fmt.Sprintf("Invalid direction %v", d))
}

func (cm TravelMap) ContainsRowColDir(row, col int, dir direction.Direction) bool {
	_, ok := cm[fmt.Sprintf("%d-%d-%s", row, col, getDirString(dir))]
	return ok
}

func (cm TravelMap) ContainsCoordAndDir(c Coord, dir direction.Direction) bool {
	_, ok := cm[fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(dir))]
	return ok
}

func (cm *TravelMap) Add(c Coord, d direction.Direction) {
	(*cm)[fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(d))] = true
}

func (cm *TravelMap) RemoveCoordAndDir(c Coord, d direction.Direction) {
	delete(*cm, fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(d)))
}

func (cm *TravelMap) Clear() {
	*cm = make(TravelMap)
}

//   _______                  _ _____  _       _____          _   __  __
//  |__   __|                | |  __ \(_)     / ____|        | | |  \/  |
//     | |_ __ __ ___   _____| | |  | |_ _ __| |     ___  ___| |_| \  / | __ _ _ __
//     | | '__/ _` \ \ / / _ \ | |  | | | '__| |    / _ \/ __| __| |\/| |/ _` | '_ \
//     | | | | (_| |\ V /  __/ | |__| | | |  | |___| (_) \__ \ |_| |  | | (_| | |_) |
//     |_|_|  \__,_| \_/ \___|_|_____/|_|_|   \_____\___/|___/\__|_|  |_|\__,_| .__/
//                                                                            | |
//                                                                            |_|

type TravelDirCostMap map[string]int

func (cm TravelDirCostMap) ContainsRowColDirCost(row, col int, dir direction.Direction) bool {
	_, ok := cm[fmt.Sprintf("%d-%d-%s", row, col, getDirString(dir))]
	return ok
}

func (cm TravelDirCostMap) ContainsCoordDirCost(c Coord, dir direction.Direction) bool {
	_, ok := cm[fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(dir))]
	return ok
}

func (cm *TravelDirCostMap) Add(c Coord, d direction.Direction, cost int) {
	(*cm)[fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(d))] = cost
}

func (cm *TravelDirCostMap) RemoveCoordDirCost(c Coord, d direction.Direction, cost int) {
	delete(*cm, fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(d)))
}

func (cm TravelDirCostMap) Get(c Coord, d direction.Direction) int {
	return cm[fmt.Sprintf("%d-%d-%s", c.Row, c.Col, getDirString(d))]
}

func (cm *TravelDirCostMap) Clear() {
	*cm = make(TravelDirCostMap)
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

func GetOnlyDiagonalOffsets() []DirOffset {
	return []DirOffset{
		dirOffsetsMap[direction.UpRight],
		dirOffsetsMap[direction.RightDown],
		dirOffsetsMap[direction.DownLeft],
		dirOffsetsMap[direction.LeftUp],
	}
}

func GetPerpendicularOffsets(dirOffset DirOffset) []DirOffset {
	if !utils.Contains[DirOffset](GetOnlyDiagonalOffsets(), dirOffset) {
		panic(fmt.Sprintf("Invalid direction, input must be diagonal direction %v", dirOffset))
	}

	if dirOffset.Dir == direction.UpRight || dirOffset.Dir == direction.DownLeft {
		return []DirOffset{dirOffsetsMap[direction.LeftUp], dirOffsetsMap[direction.RightDown]}
	}
	return []DirOffset{dirOffsetsMap[direction.UpRight], dirOffsetsMap[direction.DownLeft]}
}

func GetCounterClockwise90DegreeNeighborOffset(dirOffset DirOffset) DirOffset {
	if dirOffset.Dir == direction.Up {
		return dirOffsetsMap[direction.Left]
	}
	if dirOffset.Dir == direction.Right {
		return dirOffsetsMap[direction.Up]
	}
	if dirOffset.Dir == direction.Down {
		return dirOffsetsMap[direction.Right]
	}
	// dirOffset.Dir == direction.Left
	return dirOffsetsMap[direction.Down]
}

func GetClockwise90DegreeNeighborOffset(dirOffset DirOffset) DirOffset {
	if dirOffset.Dir == direction.Up {
		return dirOffsetsMap[direction.Right]
	}
	if dirOffset.Dir == direction.Right {
		return dirOffsetsMap[direction.Down]
	}
	if dirOffset.Dir == direction.Down {
		return dirOffsetsMap[direction.Left]
	}
	if dirOffset.Dir == direction.Left {
		return dirOffsetsMap[direction.Up]
	}

	if dirOffset.Dir == direction.UpRight {
		return dirOffsetsMap[direction.RightDown]
	}
	if dirOffset.Dir == direction.RightDown {
		return dirOffsetsMap[direction.DownLeft]
	}
	if dirOffset.Dir == direction.DownLeft {
		return dirOffsetsMap[direction.LeftUp]
	}
	// dirOffset.Dir == direction.LeftUp
	return dirOffsetsMap[direction.UpRight]
}

func GetClockwise90DegreeDirection(dir direction.Direction) direction.Direction {
	if dir == direction.Up {
		return direction.Right
	}
	if dir == direction.Right {
		return direction.Down
	}
	if dir == direction.Down {
		return direction.Left
	}
	// dir == direction.Left
	return direction.Up
}

func GetCounterClockwise90DegreeDirection(dir direction.Direction) direction.Direction {
	if dir == direction.Up {
		return direction.Left
	}
	if dir == direction.Left {
		return direction.Down
	}
	if dir == direction.Down {
		return direction.Right
	}
	// dir == direction.Right
	return direction.Up
}

func ParseCoordStr(str, separator string) Coord {
	i, j := utils.SplitIn2(str, separator)
	row, errRow := strconv.Atoi(i)
	col, errCol := strconv.Atoi(j)
	if errRow != nil || errCol != nil {
		panic(fmt.Sprintf("Invalid input %s", str))
	}
	return Coord{Row: row, Col: col}
}
