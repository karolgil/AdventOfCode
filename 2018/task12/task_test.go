package task10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile      string
		expectedResult int
	}{
		{
			inputFile:      "input1",
			expectedResult: 2281,
		},
		{
			inputFile:      "input2",
			expectedResult: 325,
		},
	} {
		result := Solution1(test.inputFile)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input1")
	}
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		inputFile      string
		expectedResult int
	}{
		{
			inputFile:      "input1",
			expectedResult: 2250000000120,
		},
		{
			inputFile:      "input2",
			expectedResult: 999999999374,
		},
	} {
		result := Solution2(test.inputFile)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input1")
	}
}
