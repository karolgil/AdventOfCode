package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		numOfPlayers, highestNumber, expectedHighestScore int
	}{
		{
			numOfPlayers:         9,
			highestNumber:        25,
			expectedHighestScore: 32,
		},
		{
			numOfPlayers:         10,
			highestNumber:        1618,
			expectedHighestScore: 8317,
		},
		{
			numOfPlayers:         13,
			highestNumber:        7999,
			expectedHighestScore: 146373,
		},
		{
			numOfPlayers:         17,
			highestNumber:        1104,
			expectedHighestScore: 2764,
		},
		{
			numOfPlayers:         21,
			highestNumber:        6111,
			expectedHighestScore: 54718,
		},
		{
			numOfPlayers:         30,
			highestNumber:        5807,
			expectedHighestScore: 37305,
		},
		{
			numOfPlayers:         429,
			highestNumber:        70901,
			expectedHighestScore: 399645,
		},
		{
			numOfPlayers:         429,
			highestNumber:        70901 * 100,
			expectedHighestScore: 3352507536,
		},
	} {
		solution := Solution(test.numOfPlayers, test.highestNumber)
		assert.Equal(t, test.expectedHighestScore, solution, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution(429, 70901*100)
	}
}
