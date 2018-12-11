package task10

import (
	"fmt"
	"math"
	"strconv"
)

func Solution1(gridSerialNumber int) string {
	grid := NewGrid(300, gridSerialNumber)
	grid.SetCellPowers()
	grid.SetCellPartialSums()
	grid.SetCellSquareScores()
	coords := grid.FindBestScoreCoordinates()
	return strconv.FormatInt(int64(coords.x+2), 10) + "," + strconv.FormatInt(int64(coords.y+2), 10)
}

func Solution2(gridSerialNumber int) string {
	grid := NewGrid(300, gridSerialNumber)
	grid.SetCellPowers()
	grid.SetCellPartialSums()
	coords, size := grid.FindBestScoreAnySize()
	return strconv.FormatInt(int64(coords.x+2), 10) + "," + strconv.FormatInt(int64(coords.y+2), 10) + "," + strconv.FormatInt(int64(size), 10)
}

type Grid struct {
	matrix       map[Coordinates]*Powers
	size         int
	serialNumber int
}

func NewGrid(size, serialNumber int) *Grid {
	return &Grid{
		matrix:       make(map[Coordinates]*Powers, size),
		size:         size,
		serialNumber: serialNumber,
	}
}

func (g Grid) SetCellPowers() {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			g.SetCellPower(x, y)
		}
	}
}

func (g Grid) SetCellPower(x, y int) {
	coordinates := Coordinates{x, y}
	if _, exists := g.matrix[coordinates]; !exists {
		g.matrix[coordinates] = &Powers{0, 0, 0}
	}
	rackID := x + 1 + 10
	power := rackID * (y + 1)
	power = power + g.serialNumber
	power = power * rackID
	hundredsDigit := int(math.Abs(float64(power / 100 % 10)))
	power = hundredsDigit - 5
	g.matrix[coordinates].cellPower = power
}

func (g Grid) GetCellPowerStr(x, y int) string {
	cellPower := strconv.FormatInt(int64(g.matrix[Coordinates{x, y}].cellPower), 10)
	if len(cellPower) == 1 {
		return " " + cellPower + " "
	}
	return cellPower + " "
}

func (g Grid) DisplayCellPowers() {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			fmt.Print(g.GetCellPowerStr(x, y))
		}
		fmt.Print("\n")
	}
}

func (g Grid) SetCellPartialSums() {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			g.SetCellPartialSum(x, y)
		}
	}
}

func (g Grid) SetCellPartialSum(x, y int) {
	coordinates := Coordinates{x, y}
	toLeft := Coordinates{x - 1, y}
	toUp := Coordinates{x, y - 1}
	toLeftUp := Coordinates{x - 1, y - 1}
	if x == 0 && y == 0 {
		g.matrix[coordinates].partialSumPower = g.matrix[coordinates].cellPower
		return
	} else if x == 0 {
		g.matrix[coordinates].partialSumPower = g.matrix[coordinates].cellPower +
			g.matrix[toUp].partialSumPower
		return
	} else if y == 0 {
		g.matrix[coordinates].partialSumPower = g.matrix[coordinates].cellPower +
			g.matrix[toLeft].partialSumPower
		return
	} else {
		g.matrix[coordinates].partialSumPower = g.matrix[coordinates].cellPower +
			g.matrix[toLeft].partialSumPower +
			g.matrix[toUp].partialSumPower -
			g.matrix[toLeftUp].partialSumPower
	}
}

func (g Grid) DisplayCellPartialSums() {
	for x := 0; x < g.size; x++ {
		for y := 0; y < g.size; y++ {
			fmt.Print(g.GetCellPartialSumStr(x, y))
		}
		fmt.Print("\n")
	}
}

func (g Grid) GetCellPartialSumStr(x, y int) string {
	cellPower := strconv.FormatInt(int64(g.matrix[Coordinates{x, y}].cellPower), 10)
	cellPartial := strconv.FormatInt(int64(g.matrix[Coordinates{x, y}].partialSumPower), 10)
	return " " + cellPower + "(" + cellPartial + ")"
}

func (g Grid) SetCellSquareScores() {
	for x := 0; x < g.size-3; x++ {
		for y := 0; y < g.size-3; y++ {
			g.SetCellSquareScore(x, y)
		}
	}
}

func (g Grid) SetCellSquareScore(x, y int) {
	g.matrix[Coordinates{x, y}].squareScore = g.matrix[Coordinates{x + 3, y + 3}].partialSumPower +
		g.matrix[Coordinates{x, y}].partialSumPower -
		g.matrix[Coordinates{x, y + 3}].partialSumPower -
		g.matrix[Coordinates{x + 3, y}].partialSumPower
}

func (g Grid) FindBestScoreCoordinates() Coordinates {
	bestCoord := Coordinates{-1, -1}
	bestScore := 0
	for coord, powers := range g.matrix {
		if powers.squareScore > bestScore {
			bestScore = powers.squareScore
			bestCoord = coord
		}
	}
	return bestCoord
}

func (g Grid) FindBestScoreAnySize() (Coordinates, int) {
	bestScore := 0
	bestX := -1
	bestY := -1
	bestSize := 0
	for size := 0; size < 300; size++ {
		for x := size; x < 300; x++ {
			for y := size; y < 300; y++ {
				score := g.matrix[Coordinates{x - size, y - size}].partialSumPower +
					g.matrix[Coordinates{x, y}].partialSumPower -
					g.matrix[Coordinates{x, y - size}].partialSumPower -
					g.matrix[Coordinates{x - size, y}].partialSumPower
				if score > bestScore {
					bestX = x - size
					bestY = y - size
					bestSize = size
					bestScore = score
				}
			}
		}
	}
	return Coordinates{bestX, bestY}, bestSize
}

type Coordinates struct {
	x, y int
}

type Powers struct {
	cellPower, partialSumPower, squareScore int
}
