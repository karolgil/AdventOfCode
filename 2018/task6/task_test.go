package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: `input1`,
			expected:  17,
		},
		{
			inputFile: `input2`,
			expected:  4284,
		},
	} {
		solution, err := Solution1(test.inputFile)
		assert.NoError(t, err, "Solution1 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input2")
	}
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		inputFile   string
		maxDistance int
		expected    int
	}{
		{
			inputFile:   `input1`,
			maxDistance: 32,
			expected:    16,
		},
		{
			inputFile:   `input2`,
			maxDistance: 10000,
			expected:    35490,
		},
	} {
		solution, err := Solution2(test.inputFile, test.maxDistance)
		assert.NoError(t, err, "Solution2 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input2", 10000)
	}
}
