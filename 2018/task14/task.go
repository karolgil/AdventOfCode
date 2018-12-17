package task10

import (
	"strconv"
	"strings"
)

func Solution1(input int) string {
	scores := calculateRecipeScores([]byte{'3', '7'})
	return string(scores[input : input+10])
}

func Solution2(input int) int {
	scores := calculateRecipeScores([]byte{'3', '7'})
	return strings.Index(string(scores), strconv.Itoa(input))
}

func calculateRecipeScores(scores []byte) []byte {
	a, b := 0, 1
	for len(scores) < 50000000 {
		score := []byte(strconv.Itoa(int(scores[a] - '0' + scores[b] - '0')))
		scores = append(scores, score...)

		a = (a + 1 + int(scores[a]-'0')) % len(scores)
		b = (b + 1 + int(scores[b]-'0')) % len(scores)
	}
	return scores
}
