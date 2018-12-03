package task1

import (
	"strconv"
	"strings"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}

	fabricParts, err := buildFabrics(lines)
	if err != nil {
		return 0, err
	}

	fabricMap := populateFabricMap(fabricParts)
	return countOverlappingParts(fabricMap), nil
}

type fabricPart struct {
	horizontalPadding int
	verticalPadding   int
	width             int
	height            int
}

func newFabricPart(inputLine string) (*fabricPart, error) {
	split := strings.Split(inputLine, " ")
	padding := strings.Split(split[2][0:len(split[2])-1], ",")
	size := strings.Split(split[3], "x")
	horizontalPadding, err := strconv.Atoi(padding[0])
	if err != nil {
		return nil, err
	}
	verticalPadding, err := strconv.Atoi(padding[1])
	if err != nil {
		return nil, err
	}
	width, err := strconv.Atoi(size[0])
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(size[1])
	if err != nil {
		return nil, err
	}
	return &fabricPart{
		horizontalPadding: horizontalPadding,
		verticalPadding:   verticalPadding,
		width:             width,
		height:            height,
	}, nil
}

func buildFabrics(input []string) ([]*fabricPart, error) {
	var fabricParts []*fabricPart
	for _, line := range input {
		f, err := newFabricPart(line)
		if err != nil {
			return []*fabricPart{}, err
		}
		fabricParts = append(fabricParts, f)
	}
	return fabricParts, nil
}

func populateFabricMap(fabricParts []*fabricPart) map[int]map[int]int {
	fabric := getEmptyFabric()
	for _, part := range fabricParts {
		for i := part.horizontalPadding; i < part.horizontalPadding+part.width; i++ {
			for j := part.verticalPadding; j < part.verticalPadding+part.height; j++ {
				fabric[i][j] += 1
			}
		}
	}
	return fabric
}

// getEmptyResultMap returns two-dimensional map of size 1000x1000 filled with 0s
func getEmptyFabric() map[int]map[int]int {
	outerMap := make(map[int]map[int]int)
	for i := 0; i < 1000; i++ {
		outerMap[i] = make(map[int]int)
		for j := 0; j < 1000; j++ {
			outerMap[i][j] = 0
		}
	}
	return outerMap
}

func countOverlappingParts(fabricMap map[int]map[int]int) int {
	var result int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabricMap[i][j] > 1 {
				result += 1
			}
		}
	}
	return result
}
