package task10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		input          int
		expectedResult string
	}{
		{
			input:          509671,
			expectedResult: "2810862211",
		},
	} {
		result := Solution1(test.input)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1(509671)
	}
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		input          int
		expectedResult int
	}{
		{
			input:          509671,
			expectedResult: 20227889,
		},
	} {
		result := Solution2(test.input)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2(509671)
	}
}
