package task10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile      string
		expectedResult string
	}{
		{
			inputFile:      "input1",
			expectedResult: "71,121",
		},
		{
			inputFile:      "input2",
			expectedResult: "7,3",
		},
		{
			inputFile:      "input3",
			expectedResult: "0,3",
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
		expectedResult string
	}{
		{
			inputFile:      "input1",
			expectedResult: "71,76",
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
