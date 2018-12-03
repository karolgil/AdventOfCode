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
	for _, test := range []struct {
		input             string
		id                int
		horizontalPadding int
		verticalPadding   int
		width             int
		height            int
	}{
		{
			input:             "#12 @ 719,695: 19x10",
			id:                12,
			horizontalPadding: 719,
			verticalPadding:   695,
			width:             19,
			height:            10,
		},
		{
			input:             "#1097 @ 382,427: 14x15",
			id:                1097,
			horizontalPadding: 382,
			verticalPadding:   427,
			width:             14,
			height:            15,
		},
	} {
		fabric, err := newFabricPart(test.input)
		assert.NoError(t, err, "newFabricPart should not raise error on valid input")
		assert.Equal(t, test.id, fabric.id)
		assert.Equal(t, test.horizontalPadding, fabric.horizontalPadding)
		assert.Equal(t, test.verticalPadding, fabric.verticalPadding)
		assert.Equal(t, test.width, fabric.width)
		assert.Equal(t, test.height, fabric.height)
	}
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

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: `input1`,
			expected:  1097,
		},
		{
			inputFile: `input2`,
			expected:  3,
		},
	} {
		solution, err := Solution2(test.inputFile)
		assert.NoError(t, err, "Solution2 should not return any errors")
		assert.Equal(t, test.expected, solution, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2("input1")
	}
}
