package task17

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/tools/container/intsets"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

type Grid struct {
	minX, minY, maxX, maxY int
	flowMap                [][]string
}

func (g *Grid) SetMaxAndMin(x, y int) {
	if x < g.minX {
		g.minX = x
	}
	if x > g.maxX {
		g.maxX = x
	}
	if y < g.minY {
		g.minY = y
	}
	if y > g.maxY {
		g.maxY = y
	}
}

func (g *Grid) SetupMap() {
	g.flowMap = make([][]string, g.maxY+2)
	for y := range g.flowMap {
		g.flowMap[y] = make([]string, g.maxX+2)
	}
}

func NewGrid() *Grid {
	return &Grid{
		minX: intsets.MaxInt,
		minY: intsets.MaxInt,
		maxX: intsets.MinInt,
		maxY: intsets.MinInt,
	}
}

func (g *Grid) SetMinsAndMaxes(lines []string) {
	lineRegex := regexp.MustCompile(`^(.)=(\d+), .=(\d+)\.{2}(\d+)$`)
	for _, line := range lines {
		ranges := lineRegex.FindAllStringSubmatch(line, -1)[0]
		lineCoord := ranges[1]
		lineCoordNum, _ := strconv.Atoi(ranges[2])
		rangeCoordMin, _ := strconv.Atoi(ranges[3])
		rangeCoordMax, _ := strconv.Atoi(ranges[4])

		if lineCoord == "x" {
			g.SetMaxAndMin(lineCoordNum, rangeCoordMin)
			g.SetMaxAndMin(lineCoordNum, rangeCoordMax)
		} else {
			g.SetMaxAndMin(rangeCoordMin, lineCoordNum)
			g.SetMaxAndMin(rangeCoordMax, lineCoordNum)
		}
	}
}

func (g *Grid) SetupClay(lines []string) {
	lineRegex := regexp.MustCompile(`^(.)=(\d+), .=(\d+)\.{2}(\d+)$`)
	for _, line := range lines {
		ranges := lineRegex.FindAllStringSubmatch(line, -1)[0]
		lineCoord := ranges[1]
		lineCoordNum, _ := strconv.Atoi(ranges[2])
		rangeCoordMin, _ := strconv.Atoi(ranges[3])
		rangeCoordMax, _ := strconv.Atoi(ranges[4])

		if lineCoord == "x" {
			for y := rangeCoordMin; y <= rangeCoordMax; y++ {
				g.SetToClay(y, lineCoordNum)
			}
		} else {
			for x := rangeCoordMin; x <= rangeCoordMax; x++ {
				g.SetToClay(lineCoordNum, x)
			}
		}
	}
}

func (g *Grid) Fill(y, x int) {
	if y > g.maxY {
		return
	} else if !g.IsEmptyOrFall(y, x) {
		return
	}

	if !g.IsEmptyOrFall(y+1, x) {
		lastFilledToRight := g.FillToTheRight(y, x)
		lastFilledToLeft := g.FillToTheLeft(y, x)
		if g.IsEmptyOrFall(y+1, lastFilledToLeft) || g.IsEmptyOrFall(y+1, lastFilledToRight) {
			g.Fill(y, lastFilledToLeft)
			g.Fill(y, lastFilledToRight)
		} else if g.IsClay(y, lastFilledToLeft) && g.IsClay(y, lastFilledToRight) {
			for filledX := lastFilledToLeft + 1; filledX < lastFilledToRight; filledX++ {
				g.SetToStill(y, filledX)
			}
		}
	} else if g.IsEmpty(y, x) {
		g.SetToFlow(y, x)
		g.Fill(y+1, x)
		if g.IsStill(y+1, x) {
			g.Fill(y, x)
		}
	}
}

func (g *Grid) FillToTheLeft(y, x int) int {
	cellToTheLeft := x - 1
	for g.IsEmptyOrFall(y, cellToTheLeft) && !g.IsEmptyOrFall(y+1, cellToTheLeft) {
		g.SetToFlow(y, cellToTheLeft)
		cellToTheLeft--
	}
	return cellToTheLeft
}

func (g *Grid) FillToTheRight(y, x int) int {
	cellToTheRight := x
	for g.IsEmptyOrFall(y, cellToTheRight) && !g.IsEmptyOrFall(y+1, cellToTheRight) {
		g.SetToFlow(y, cellToTheRight)
		cellToTheRight++
	}
	return cellToTheRight
}

func (g *Grid) SetToStill(y, x int) {
	g.flowMap[y][x] = "~"
}

func (g *Grid) SetToFlow(y, x int) {
	g.flowMap[y][x] = "|"
}

func (g *Grid) SetToClay(y, x int) {
	g.flowMap[y][x] = "#"
}

func (g *Grid) IsClay(y, x int) bool {
	return g.flowMap[y][x] == "#"
}

func (g *Grid) IsStill(y, x int) bool {
	return g.flowMap[y][x] == "~"
}

func (g *Grid) IsFlow(y, x int) bool {
	return g.flowMap[y][x] == "|"
}

func (g *Grid) IsEmpty(y, x int) bool {
	return g.flowMap[y][x] == ""
}

func (g *Grid) IsEmptyOrFall(y, x int) bool {
	return g.IsEmpty(y, x) || g.IsFlow(y, x)
}

func (g *Grid) Show() {
	for y := g.minY - 1; y <= g.maxY; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			if g.IsEmpty(y, x) {
				fmt.Print(".")
			} else {
				fmt.Print(string(g.flowMap[y][x]))
			}
		}
		fmt.Print("\n")
	}
}

func (g *Grid) CountAllWater() int {
	count := 0
	for y := g.minY - 1; y <= g.maxY; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			if g.IsStill(y, x) || g.IsFlow(y, x) {
				count++
			}
		}
	}
	return count
}

func (g *Grid) CountStillWater() int {
	count := 0
	for y := g.minY - 1; y <= g.maxY; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			if g.IsStill(y, x) {
				count++
			}
		}
	}
	return count
}

func Solution1(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	grid := NewGrid()
	grid.SetMinsAndMaxes(lines)
	grid.SetupMap()
	grid.SetupClay(lines)
	grid.Fill(0, 500)
	//grid.Show()

	return grid.CountAllWater() - 1 // subtracting water input
}

func Solution2(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	grid := NewGrid()
	grid.SetMinsAndMaxes(lines)
	grid.SetupMap()
	grid.SetupClay(lines)
	grid.Fill(0, 500)
	//grid.Show()

	return grid.CountStillWater() // subtracting water input
}
