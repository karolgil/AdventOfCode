package task1

import (
	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, nil
	}

	var twos int
	var threes int

	for _, line := range lines {
		charMap := getEmptyCharMap()
		for _, char := range line {
			charMap[char] = charMap[char] + 1
			// this could probably break faster after reaching 2 & 3 in the same line, but strings are not so long, so w/e
		}
		if containsValue(charMap, 2) {
			twos += 1
		}
		if containsValue(charMap, 3) {
			threes += 1
		}
	}

	return twos * threes, nil
}

func containsValue(m map[rune]int, value int) bool {
	for _, val := range m {
		if val == value {
			return true
		}
	}
	return false
}

func getEmptyCharMap() map[rune]int {
	charMap := make(map[rune]int)
	for _, l := range alphabet() {
		charMap[l] = 0
	}
	return charMap
}

func alphabet() string {
	b := make([]byte, 26)
	for i := range b {
		b[i] = 'a' + byte(i)
	}
	return string(b)
}
