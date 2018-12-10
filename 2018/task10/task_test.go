package task10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		inputFile    string
		expectedText string
		expectedTime int
	}{
		{
			inputFile:    `input1`,
			expectedText: `ZZCBGGCJ`,
			expectedTime: 10886,
		},
		{
			inputFile:    `input2`,
			expectedText: `HI`,
			expectedTime: 3,
		},
	} {
		text, time, err := Solution(test.inputFile)
		assert.NoError(t, err, "Solution should not return any errors")
		assert.Equal(t, test.expectedText, text, "should be equal")
		assert.Equal(t, test.expectedTime, time, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution("input1")
	}
}
