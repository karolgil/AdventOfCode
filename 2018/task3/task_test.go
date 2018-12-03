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
			expected:  124850,
		},
		{
			inputFile: `input2`,
			expected:  4,
		},
	} {
		solution, err := Solution1(test.inputFile)
		assert.NoError(t, err, "Solution1 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func TestNewFabricPart(t *testing.T) {
	fabric, err := newFabricPart("#12 @ 719,695: 19x10")
	assert.NoError(t, err, "newFabricPart should not raise error on valid input")
	assert.Equal(t, 719, fabric.horizontalPadding)
	assert.Equal(t, 695, fabric.verticalPadding)
	assert.Equal(t, 19, fabric.width)
	assert.Equal(t, 10, fabric.height)
}

func TestBuildFabrics(t *testing.T) {
	fabrics, err := buildFabrics([]string{"#1 @ 1,3: 4x4", "#3 @ 5,5: 2x2"})
	assert.NoError(t, err, "buildFabrics should not raise error on valid input")
	assert.Len(t, fabrics, 2)
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1("input1")
	}
}
