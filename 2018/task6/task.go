package task1

import (
	"math"
	"strconv"
	"strings"

	"golang.org/x/tools/container/intsets"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	alphabet := utils.AlphabetLowerAndUpper()
	maxX, maxY, points := getPoints(lines, alphabet)
	matrix := NewEmptyMatrix(maxX+1, maxY+1)
	matrix.SetNearestPoints(maxX, maxY, points)
	resultCount := NewResultsCounter(points)
	resultCount.RemoveSidesOfMatrix(maxX, maxY, matrix)

	return resultCount.FindBest(matrix), nil
}

func getPoints(lines []string, alphabet string) (int, int, Points) {
	maxX := 0
	maxY := 0
	var points Points
	for i, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Trim(strings.Split(line, ",")[1], " "))
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		points = append(points, Point{id: string(alphabet[i]), x: x, y: y})
	}
	return maxX, maxY, points
}

type ResultsCounter map[string]int

func (rc ResultsCounter) FindBest(matrix Matrix) int {
	for _, row := range matrix {
		for _, col := range row {
			if _, exists := rc[col]; exists {
				rc[col] += 1
			}
		}
	}
	best := 0
	for _, count := range rc {
		if count > best {
			best = count
		}
	}
	return best
}

func (rc ResultsCounter) RemoveSidesOfMatrix(maxX, maxY int, matrix Matrix) {
	for x := 1; x <= maxX; x = x + maxX - 1 {
		for y := 1; y <= maxY; y++ {
			delete(rc, matrix[x][y])
		}
	}
	for y := 1; y <= maxY; y = y + maxY - 1 {
		for x := 1; x <= maxX; x++ {
			delete(rc, matrix[x][y])
		}
	}
}

func NewResultsCounter(points Points) ResultsCounter {
	resultCount := make(map[string]int)
	for _, point := range points {
		resultCount[point.id] = 0
	}
	return resultCount
}

type Matrix [][]string

func NewEmptyMatrix(x, y int) Matrix {
	matrix := make([][]string, x)
	for i := range matrix {
		matrix[i] = make([]string, y)
	}
	return matrix
}

func (m Matrix) SetNearestPoints(maxX, maxY int, ps Points) {
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			m[x][y] = ps.findNearestTo(Point{"", x, y})
		}
	}
}

func (m Matrix) GetSumOfFieldWithDistLowerThan(maxX, maxY, maxDist int, points Points) int {
	sum := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if points.isCloserThenDistFromAll(Point{"", x, y}, maxDist) {
				sum += 1
			}
		}
	}
	return sum
}

type Points []Point

func (ps Points) findNearestTo(p Point) string {
	shortestDist := intsets.MaxInt
	nearest := ""
	for _, point := range ps {
		dist := point.distanceTo(p)
		if dist == 0 {
			return point.id
		}
		if dist == shortestDist {
			nearest = "."
			shortestDist = dist
		}
		if dist < shortestDist {
			nearest = point.id
			shortestDist = dist
		}
	}
	return nearest
}

func (ps Points) isCloserThenDistFromAll(p Point, maxDist int) bool {
	var distSum int
	for _, point := range ps {
		distSum += point.distanceTo(p)
	}
	if distSum < maxDist {
		return true
	}
	return false
}

type Point struct {
	id   string
	x, y int
}

func (p Point) distanceTo(p2 Point) int {
	xDist := math.Abs(float64(p.x - p2.x))
	yDist := math.Abs(float64(p.y - p2.y))
	return int(xDist + yDist)
}

func Solution2(inputFile string, maxDist int) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	alphabet := utils.AlphabetLowerAndUpper()
	maxX, maxY, points := getPoints(lines, alphabet)
	matrix := NewEmptyMatrix(maxX+1, maxY+1)

	return matrix.GetSumOfFieldWithDistLowerThan(maxX, maxY, maxDist, points), nil
}
