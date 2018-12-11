package task10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	for _, test := range []struct {
		gridSerialNumber int
		expectedResult   string
	}{
		{
			gridSerialNumber: 18,
			expectedResult:   "33,45",
		},
		{
			gridSerialNumber: 5177,
			expectedResult:   "235,22",
		},
	} {
		result := Solution1(test.gridSerialNumber)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1(5177)
	}
}

func TestSolution2(t *testing.T) {
	for _, test := range []struct {
		gridSerialNumber int
		expectedResult   string
	}{
		{
			gridSerialNumber: 18,
			expectedResult:   "90,269,16",
		},
		{
			gridSerialNumber: 5177,
			expectedResult:   "231,135,8",
		},
	} {
		result := Solution2(test.gridSerialNumber)
		assert.Equal(t, test.expectedResult, result, "should be equal")
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2(5177)
	}
}
