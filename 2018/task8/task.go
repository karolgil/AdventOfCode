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
	_, metadataSum := calculateMetadataSum(splitStringToIntSlice(lines[0]))
	return metadataSum, nil
}

func Solution2(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	_, treeScore := calculateScore(splitStringToIntSlice(lines[0]))
	return treeScore, nil
}

func calculateScore(input []int) ([]int, int) {
	childNum := input[0]
	metadataNum := input[1]
	input = input[2:]
	var childrenScores []int

	for range make([]int, childNum) {
		var childScore int
		input, childScore = calculateScore(input)
		childrenScores = append(childrenScores, childScore)
	}

	if childNum == 0 {
		ownScore := sum(input[:metadataNum]...)
		return input[metadataNum:], ownScore
	} else {
		sumScore := 0
		for _, metadata := range input[:metadataNum] {
			if metadata < len(childrenScores)+1 {
				sumScore += childrenScores[metadata-1]
			}
		}
		return input[metadataNum:], sumScore
	}
}

func calculateMetadataSum(input []int) ([]int, int) {
	childNum := input[0]
	metadataNum := input[1]
	input = input[2:]
	childrenMetadata := 0

	for range make([]int, childNum) {
		var childMetadata int
		input, childMetadata = calculateMetadataSum(input)
		childrenMetadata += childMetadata
	}

	ownMetadata := sum(input[:metadataNum]...)
	return input[metadataNum:], childrenMetadata + ownMetadata
}

func sum(input ...int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}

func splitStringToIntSlice(inputLine string) []int {
	var intSlice []int
	for _, char := range strings.Split(inputLine, " ") {
		num, _ := strconv.Atoi(char)
		intSlice = append(intSlice, num)
	}
	return intSlice
}
