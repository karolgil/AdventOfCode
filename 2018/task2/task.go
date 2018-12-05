package task1

import (
	"errors"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
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
	for _, l := range utils.Alphabet() {
		charMap[l] = 0
	}
	return charMap
}

func Solution2(inputFile string) (string, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				break
			}
			str1 := lines[i]
			str2 := lines[j]
			diffCount := 0
			for k := 0; k < len(str1); k++ {
				if str1[k] != str2[k] {
					diffCount += 1
				}
				if diffCount > 1 {
					break
				}
			}
			if diffCount == 1 {
				var result string
				for k := 0; k < len(str1); k++ {
					if str1[k] == str2[k] {
						result += string(str1[k])
					}
				}
				return result, nil
			}
		}
	}

	return "", errors.New("could not find result")
}
