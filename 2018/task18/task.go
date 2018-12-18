package task18

import (
	"fmt"
	"sort"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

type Tile string

var OpenGround = Tile(".")
var Tree = Tile("|")
var Lumberyard = Tile("#")

func NewTile(input rune) Tile {
	switch string(input) {
	case ".":
		return OpenGround
	case "|":
		return Tree
	case "#":
		return Lumberyard
	default:
		panic(fmt.Sprintf("Unexpected input: %s", string(input)))
	}
}

type Point struct {
	y, x int
}

type ConcurrentGrid struct {
	grid map[Point]Tile
}

type NorthPole struct {
	grids                 [2]*ConcurrentGrid
	currentGrid, nextGrid int
}

func (np *NorthPole) SetupCurrentGrid(lines []string) {
	for y, line := range lines {
		for x, r := range line {
			np.grids[np.currentGrid].grid[Point{y, x}] = NewTile(r)
		}
	}
}

func (np *NorthPole) Grow() {
	current := np.currentGrid
	for point, tile := range np.grids[current].grid {
		tileInNext := tile
		switch tile {
		case OpenGround:
			if np.ShouldBecomeTree(point) {
				tileInNext = Tree
			}
			break
		case Tree:
			if np.ShouldBecomeLumberyard(point) {
				tileInNext = Lumberyard
			}
			break
		case Lumberyard:
			if np.ShouldBecomeOpenGround(point) {
				tileInNext = OpenGround
			}
			break
		default:
			panic("Unexpected tile")
		}
		np.SetInNext(point, tileInNext)
	}

	np.currentGrid = np.nextGrid
	np.nextGrid = current
}

func (np *NorthPole) SetInNext(point Point, tile Tile) {
	np.grids[np.nextGrid].grid[point] = tile
}

func (np *NorthPole) ShouldBecomeTree(point Point) bool {
	treesCount := 0
	for y := point.y - 1; y <= point.y+1; y++ {
		for x := point.x - 1; x <= point.x+1; x++ {
			if x != point.x || y != point.y {
				if np.grids[np.currentGrid].grid[Point{y, x}] == Tree {
					treesCount++
				}
			}
		}
	}
	return treesCount >= 3
}

func (np *NorthPole) ShouldBecomeLumberyard(point Point) bool {
	lumberyardCount := 0
	for y := point.y - 1; y <= point.y+1; y++ {
		for x := point.x - 1; x <= point.x+1; x++ {
			if x != point.x || y != point.y {
				tile := np.grids[np.currentGrid].grid[Point{y, x}]
				if tile == Lumberyard {
					lumberyardCount++
				}
			}
		}
	}
	return lumberyardCount >= 3
}

func (np *NorthPole) ShouldBecomeOpenGround(point Point) bool {
	lumberyardCount := 0
	treeCount := 0
	for y := point.y - 1; y <= point.y+1; y++ {
		for x := point.x - 1; x <= point.x+1; x++ {
			if x != point.x || y != point.y {
				if np.grids[np.currentGrid].grid[Point{y, x}] == Lumberyard {
					lumberyardCount++
				}
				if np.grids[np.currentGrid].grid[Point{y, x}] == Tree {
					treeCount++
				}
			}
		}
	}
	return lumberyardCount < 1 || treeCount < 1
}

func (np *NorthPole) ShowCurrent() {
	var sortedPoints []Point
	for key := range np.grids[np.currentGrid].grid {
		sortedPoints = append(sortedPoints, key)
	}
	sort.Slice(sortedPoints, func(i, j int) bool {
		if sortedPoints[i].y != sortedPoints[j].y {
			return sortedPoints[i].y < sortedPoints[j].y
		}
		return sortedPoints[i].x < sortedPoints[j].x
	})
	for _, point := range sortedPoints {
		if point.x == 0 {
			fmt.Print("\n")
		}
		fmt.Print(np.grids[np.currentGrid].grid[point])
	}
	fmt.Print("\n")
}

func (np *NorthPole) CountTrees() int {
	count := 0
	for _, tile := range np.grids[np.currentGrid].grid {
		if tile == Tree {
			count++
		}
	}
	return count
}

func (np *NorthPole) CountLumberyards() int {
	count := 0
	for _, tile := range np.grids[np.currentGrid].grid {
		if tile == Lumberyard {
			count++
		}
	}
	return count
}

func NewNorthPole() *NorthPole {
	return &NorthPole{
		grids:       [2]*ConcurrentGrid{{make(map[Point]Tile)}, {make(map[Point]Tile)}},
		currentGrid: 0,
		nextGrid:    1,
	}
}

func Solution1(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	northPole := NewNorthPole()
	northPole.SetupCurrentGrid(lines)
	//fmt.Println("Initially")
	//northPole.ShowCurrent()
	for i := 1; i <= 10; i++ {
		//fmt.Printf("\nAfter %d min\n", i)
		northPole.Grow()
		//northPole.ShowCurrent()
	}
	woods := northPole.CountTrees()
	lumberyards := northPole.CountLumberyards()

	return woods * lumberyards
}

type PositionValue struct {
	positions []int
	count     int
	value     int
}

func Solution2(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	northPole := NewNorthPole()
	northPole.SetupCurrentGrid(lines)

	positionValues := make(map[int]*PositionValue)
	for i := 1; i <= 1000; i++ {
		northPole.Grow()
		woods := northPole.CountTrees()
		lumberyards := northPole.CountLumberyards()
		if _, exists := positionValues[woods*lumberyards]; exists {
			positionValues[woods*lumberyards].positions = append(positionValues[woods*lumberyards].positions, i)
			positionValues[woods*lumberyards].count++
		} else {
			positionValues[woods*lumberyards] = &PositionValue{
				positions: []int{i},
				count:     1,
				value:     woods * lumberyards,
			}
		}
	}
	var bestPositionValue PositionValue
	bestCount := 0
	for _, positionValue := range positionValues {
		if positionValue.count > bestCount {
			bestPositionValue = *positionValue
			bestCount = positionValue.count
		}
	}

	distance := (bestPositionValue.positions[len(bestPositionValue.positions)-1] - bestPositionValue.positions[0]) / (len(bestPositionValue.positions) - 1)

	firstOccurrenceAfter1000 := bestPositionValue.positions[len(bestPositionValue.positions)-1]
	for firstOccurrenceAfter1000 <= 1000 {
		firstOccurrenceAfter1000 += distance
	}
	for i := 1001; i <= firstOccurrenceAfter1000; i++ {
		northPole.Grow()
	}

	justBelow1000000000 := firstOccurrenceAfter1000
	for justBelow1000000000 < 1000000000 {
		justBelow1000000000 += distance
	}
	justBelow1000000000 -= distance

	for i := justBelow1000000000 + 1; i <= 1000000000; i++ {
		northPole.Grow()
	}

	woods := northPole.CountTrees()
	lumberyards := northPole.CountLumberyards()
	return woods * lumberyards
}
