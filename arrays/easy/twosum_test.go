package easy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Command: go test -timeout 60s -run ^TestTwoSum -v

=== RUN   TestTwoSum
=== RUN   TestTwoSum/Test_empty_array
    twosum_test.go:84: Time Taken By BruteForce approach: 123ns
    twosum_test.go:92: Time Taken By Optimized approach: 168ns
=== RUN   TestTwoSum/Test_Single_Element_No_target
    twosum_test.go:84: Time Taken By BruteForce approach: 76ns
    twosum_test.go:92: Time Taken By Optimized approach: 217ns
=== RUN   TestTwoSum/Test_Single_Element_and_same_target
    twosum_test.go:84: Time Taken By BruteForce approach: 66ns
    twosum_test.go:92: Time Taken By Optimized approach: 96ns
=== RUN   TestTwoSum/Test_NO_Target_found_with_All_positive_integers
    twosum_test.go:84: Time Taken By BruteForce approach: 129ns
    twosum_test.go:92: Time Taken By Optimized approach: 294ns
=== RUN   TestTwoSum/Test_NO_Target_found_with_All_negative_integers
    twosum_test.go:84: Time Taken By BruteForce approach: 76ns
    twosum_test.go:92: Time Taken By Optimized approach: 174ns
=== RUN   TestTwoSum/Test_Target_found_with_All_negative_integers
    twosum_test.go:84: Time Taken By BruteForce approach: 136ns
    twosum_test.go:92: Time Taken By Optimized approach: 442ns
=== RUN   TestTwoSum/Test_Target_found_with_All_positive_integers
    twosum_test.go:84: Time Taken By BruteForce approach: 177ns
    twosum_test.go:92: Time Taken By Optimized approach: 421ns
=== RUN   TestTwoSum/Test_Target_found
    twosum_test.go:84: Time Taken By BruteForce approach: 125ns
    twosum_test.go:92: Time Taken By Optimized approach: 270ns
=== RUN   TestTwoSum/Test_Target_found_with_long_array
    twosum_test.go:84: Time Taken By BruteForce approach: 23.787799ms
    twosum_test.go:92: Time Taken By Optimized approach: 1.215514ms

*/
func TestTwoSum(t *testing.T) {

	tests := []struct {
		name           string
		inputNums      []int
		target         int
		expectedOutput []int
	}{
		{
			name:           "Test_empty_array",
			inputNums:      []int{},
			target:         0,
			expectedOutput: []int{},
		},
		{
			name:           "Test_Single_Element_No_target",
			inputNums:      []int{2},
			target:         10,
			expectedOutput: []int{},
		},
		{
			name:           "Test_Single_Element_and_same_target",
			inputNums:      []int{2},
			target:         2,
			expectedOutput: []int{},
		},
		{
			name:           "Test_NO_Target_found_with_All_positive_integers",
			inputNums:      []int{2, 5, 10},
			target:         2,
			expectedOutput: []int{},
		},
		{
			name:           "Test_NO_Target_found_with_All_negative_integers",
			inputNums:      []int{-2, -5, -10},
			target:         2,
			expectedOutput: []int{},
		},
		{
			name:           "Test_Target_found_with_All_negative_integers",
			inputNums:      []int{-2, -5, -10, -20, -50},
			target:         -70,
			expectedOutput: []int{3, 4},
		},
		{
			name:           "Test_Target_found_with_All_positive_integers",
			inputNums:      []int{2, 5, 10, 20, 50},
			target:         30,
			expectedOutput: []int{2, 3},
		},
		{
			name:           "Test_Target_found",
			inputNums:      []int{2, -5, -10, 5, 50},
			target:         0,
			expectedOutput: []int{1, 3},
		},
		{
			name:           "Test_Target_found_with_long_array",
			inputNums:      generateArray(10000),
			target:         19994,
			expectedOutput: []int{9994, 9998},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			start := time.Now()
			out := twoSumBruteForce(test.inputNums, test.target)
			end := time.Now()
			taken := end.Sub(start)

			assert.Equal(t, test.expectedOutput, out)
			t.Logf("Time Taken By BruteForce approach: %v", taken)

			start = time.Now()
			out = twoSumOptimized(test.inputNums, test.target)
			end = time.Now()
			taken = end.Sub(start)

			assert.Equal(t, test.expectedOutput, out)
			t.Logf("Time Taken By Optimized approach: %v", taken)

		})

	}

}

func generateArray(size int) []int {
	var ar []int = []int{}
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			ar = append(ar, i+1)
		}else{
			ar = append(ar, -i-1)
		}
	}
	return ar
}
