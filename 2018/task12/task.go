package task10

import (
	"fmt"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

type PotPatterns map[string]string

func Solution1(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	patterns := make(PotPatterns, len(lines)-2)
	for i := 2; i < len(lines); i++ {
		patterns[string(lines[i][0:5])] = string(lines[i][9])
	}

	currentPots := lines[0][15:]
	for i := 0; i < 20; i++ {
		currentPots = "...." + currentPots + "...."
		nextPots := ""
		for x := 2; x < len(currentPots)-2; x++ {
			if _, matches := patterns[currentPots[x-2:x+3]]; matches {
				nextPots += patterns[currentPots[x-2:x+3]]
			} else {
				nextPots += "."
			}

		}
		currentPots = nextPots
	}

	return CalculateSum(currentPots, 40)
}

func Solution2(inputFile string) int {
	lines, _ := utils.ReadLinesFrom(inputFile)

	patterns := make(PotPatterns, len(lines)-2)
	for i := 2; i < len(lines); i++ {
		patterns[string(lines[i][0:5])] = string(lines[i][9])
	}

	currentPots := lines[0][15:]
	previousSum := CalculateSum(currentPots, 0)
	var diffs []int
	for i := 0; i < 400; i++ {
		currentPots = fmt.Sprintf("....%s....", currentPots)
		nextPots := ""
		for x := 2; x < len(currentPots)-2; x++ {
			if _, matches := patterns[currentPots[x-2:x+3]]; matches {
				nextPots += patterns[currentPots[x-2:x+3]]
			} else {
				nextPots += "."
			}

		}
		currentPots = nextPots
		currSum := CalculateSum(currentPots, i*2)
		diffs = append(diffs, currSum-previousSum)
		if len(diffs) > 100 {
			diffs = diffs[1:]
		}
		previousSum = currSum
	}

	sumDiffs := 0
	for _, d := range diffs {
		sumDiffs += d
	}
	sumAverage := sumDiffs / len(diffs)

	total := (50000000000-400)*sumAverage + CalculateSum(currentPots, 400*2)

	return total
}

func CalculateSum(pots string, offset int) int {
	sum := 0
	for i, char := range pots {
		if string(char) == "#" {
			sum += i - offset
		}
	}
	return sum
}
