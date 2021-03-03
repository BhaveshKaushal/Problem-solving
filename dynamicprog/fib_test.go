package dynamicprog

import (
	"fmt"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {

	tests := []struct {
		input          int64
		expectedOutput int64
	}{
		{
			input:          0,
			expectedOutput: 0,
		},
		{
			input:          1,
			expectedOutput: 1,
		},
		{
			input:          2,
			expectedOutput: 1,
		},
		{
			input:          3,
			expectedOutput: 2,
		},
		{
			input:          10,
			expectedOutput: 55,
		},
		{
			input:          46,
			expectedOutput: 1836311903,
		},
	}

	for _, test := range tests {

		t.Run(fmt.Sprintf("%v --> %v", test.input, test.expectedOutput), func(t *testing.T) {
			//t.Logf("RUNNING test CASE: %v -- > %v", test.input, test.expectedOutput)

			/*
			 Below code to handle timeout panic doesn't work because 
			 because test is timing out in different go routine

			defer func() {
				r := recover();
				t.Logf("R=%v",r) 
				if r != nil {
					t.Errorf("Test panicked!!!")
					assert.Contains(t, fmt.Sprintf("%v", r), "test timed out")
				} else {
					t.Log("TEST DID NOT PANICKED!!")
				}
			}()
			*/
			start := time.Now()

			out := fib(test.input)
			end := time.Now()
			taken := end.Sub(start)

			assert.Equal(t, test.expectedOutput, out)
			t.Logf("Time Taken: %v", taken)

		})

	}

}
