package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourglassSum(t *testing.T) {

	tests := []struct {
		name           string
		input        [][]int32
		expectedOutput int32
	}{
		{
			name:           "test_case_1",
			input:        [][]int32{{1,1,1,0,0,0}},
			expectedOutput: -2147483648,
		},

		{
			name:           "test_case_2",
			input:        [][]int32{{1,1,1,0,0,0},{0,1,0,0,0,0},{1,1,1,0,0,0},{0,0,2,4,4,0},{0,0,0,2,0,0},{0,0,1,2,4,0}},
			expectedOutput: 19,
		},
		{
			name:           "test_case_3",
			input:        [][]int32{{0,-3,0,0,0},{0,-1,0,0,0},{-1,-1,-1,0,0},{0,0,-2,-4,-4},{0,0,0,-2,-4}},
			expectedOutput: -1,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			out := hourglassSum(test.input)
			assert.Equal(t, test.expectedOutput, out)

		})

	}

}
