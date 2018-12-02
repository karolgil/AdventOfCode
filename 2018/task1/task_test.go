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
			expected:  442,
		},
		{
			inputFile: `input2`,
			expected:  3,
		},
		{
			inputFile: `input3`,
			expected:  2,
		},
	} {
		solution, err := Solution1(test.inputFile)
		assert.NoError(t, err, "Solution1 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input1")
	}
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: `input1`,
			expected:  59908,
		},
		{
			inputFile: `input2`,
			expected:  2,
		},
		{
			inputFile: `input3`,
			expected:  -2,
		},
	} {
		solution, err := Solution2(test.inputFile)
		assert.NoError(t, err, "Solution1 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal for file %s", test.inputFile)
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input1")
	}
}
