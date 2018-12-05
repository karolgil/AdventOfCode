package task1

import (
	"math"
	"strings"

	"github.com/karolgil/AdventOfCode/2018/utils"
)

func Solution1(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	reactionChain := []rune(lines[0])
	reactionChain = getChainPostReactions(reactionChain)
	return len(reactionChain), nil
}

func getChainPostReactions(reactionChain []rune) []rune {
	for {
		reacted := false
		for i, char := range reactionChain[:len(reactionChain)-1] {
			if isReacting(char, rune(reactionChain[i+1])) {
				reactionChain = append(reactionChain[:i], reactionChain[i+2:]...)
				reacted = true
				break
			}
		}
		if !reacted {
			break
		}
	}
	return reactionChain
}

func isReacting(a, b rune) bool {
	return math.Abs(float64(a-b)) == 32
}

func Solution2(inputFile string) (int, error) {
	lines, err := utils.ReadLinesFrom(inputFile)
	if err != nil {
		return 0, err
	}
	reactionChain := []rune(lines[0])

	currentBest := len(reactionChain)
	for _, char := range utils.Alphabet() {
		chainWithoutChar := getChainWithoutChar(reactionChain, char)
		chainWithoutChar = getChainPostReactions(chainWithoutChar)
		if len(chainWithoutChar) < currentBest {
			currentBest = len(chainWithoutChar)
		}
	}
	return currentBest, nil
}

func getChainWithoutChar(reactionChain []rune, charToRemove rune) []rune {
	reactionString := string(reactionChain)
	reactionString = strings.Replace(reactionString, strings.ToLower(string(charToRemove)), "", -1)
	reactionString = strings.Replace(reactionString, strings.ToUpper(string(charToRemove)), "", -1)
	return []rune(reactionString)
}
