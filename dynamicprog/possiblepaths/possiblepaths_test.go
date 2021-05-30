package dynamicprog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPossiblePaths(t *testing.T) {

	tests := []struct {
		name           string
		ar             [][]int
		m              int
		n              int
		expectedOutput [][]int
	}{
		{
			name:           "Test_Case_1",
			ar:             [][]int{{1, 2, 3}},
			m:              1,
			n:              3,
			expectedOutput: [][]int{{1, 2, 3}},
		},
		{
			name:           "Test_Case_2",
			ar:             [][]int{{1, 2}, {3, 4}},
			m:              2,
			n:              2,
			expectedOutput: [][]int{{1, 3, 4}, {1, 2, 4}},
		},
		{
			name:           "Test_Case_3",
			ar:             [][]int{{1, 2, 3}, {4, 5, 6}},
			m:              2,
			n:              3,
			expectedOutput: [][]int{{1, 4, 5, 6}, {1, 2, 5, 6}, {1, 2, 3, 6}},
		},
		{
			name:           "Test_Case_4",
			ar:             [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			m:              3,
			n:              3,
			expectedOutput: [][]int{{1, 4, 7, 8, 9}, {1, 4, 5, 8, 9}, {1, 4, 5, 6, 9}, {1, 2, 5, 8, 9}, {1, 2, 5, 6, 9}, {1, 2, 3, 6, 9}},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			start := time.Now()

			out := possiblePaths(test.ar, test.m, test.n)
			end := time.Now()
			taken := end.Sub(start)

			assert.Equal(t, test.expectedOutput, out)
			t.Logf("Time Taken By BruteForce approach: %v", taken)
			
			
		})

	}

}
