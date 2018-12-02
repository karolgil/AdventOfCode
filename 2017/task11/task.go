package task11

import "fmt"

func Solution(input string) int {
	runes := []rune(input)
	if len(runes) <= 1 {
		return 0
	}
	if runes[0] == runes[len(runes)-1] {
		runes = append(runes, runes[len(runes)-1])
	}
	result := 0
	for i, char := range runes {
		if i == len(runes)-1 {
			break
		}
		if runes[i+1] == char {
			result += int(char - '0')
		}
	}

	fmt.Println(string(runes))
	return result
}
