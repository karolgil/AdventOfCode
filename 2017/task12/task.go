package task11

import "fmt"

func Solution(input string) int {
	runes := []rune(input)
	if len(runes) <= 1 {
		return 0
	}
	runesAll := append(runes, runes...)
	splitNum := len(runes) / 2
	result := 0
	for i, char := range runes {
		if i == len(runes) {
			break
		}
		if runesAll[i+splitNum] == char {
			result += int(char - '0')
		}
	}

	fmt.Println(string(runes))
	return result
}
