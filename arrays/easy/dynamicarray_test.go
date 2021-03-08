package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {

	tests := []struct {
		name           string
		n              int32
		queries        [][]int32
		expectedOutput []int32
	}{
		{
			name:           "test_case_1",
			n:              2,
			queries:        [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}},
			expectedOutput: []int32{7, 3},
		},

		{
			name:           "test_case_2",
			n:              5,
			queries:        [][]int32{{1, 2, 5}, {1, 3, 7}, {1, 1, 3}, {1, 1, 10}, {2, 1, 3}},
			expectedOutput: []int32{10},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			out := dynamicArray(test.n, test.queries)
			assert.Equal(t, test.expectedOutput, out)

		})

	}

}
