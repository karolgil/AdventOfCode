package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile string
		expected  string
	}{
		{
			inputFile: `input1`,
			expected:  "EPWCFXKISTZVJHDGNABLQYMORU",
		},
		{
			inputFile: `input2`,
			expected:  "CABDFE",
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
		inputFile    string
		numOfWorkers int
		expected     int
	}{
		{
			inputFile:    `input1`,
			numOfWorkers: 15,
			expected:     0,
		},
		{
			inputFile:    `input2`,
			numOfWorkers: 2,
			expected:     256,
		},
	} {
		solution, err := Solution2(test.inputFile, test.numOfWorkers)
		assert.NoError(t, err, "Solution2 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input2", 15)
	}
}
