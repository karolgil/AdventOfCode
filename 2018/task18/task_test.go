package task18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile      string
		expectedResult int
	}{
		//{
		//	inputFile:      "input1",
		//	expectedResult: 1147,
		//},
		{
			inputFile:      "input2",
			expectedResult: 621205,
		},
	} {
		result := Solution1(test.inputFile)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input2")
	}
	// withoutParallel:     0.0137
	// with naive sync.Map: 0.0210
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		inputFile      string
		expectedResult int
	}{
		{
			inputFile:      "input2",
			expectedResult: 0,
		},
	} {
		result := Solution2(test.inputFile)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input2")
	}
}
