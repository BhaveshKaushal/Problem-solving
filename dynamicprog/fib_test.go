package dynamicprog

import (
	"fmt"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

/*
NOTE: Time taken by each test case differs based hardware configuration(you might need to change the timeout to avoid panic while running the test)
Objective of test case is to compare computing time between brute force and optimized approach
Command: go test -timeout 60s -run ^TestFib -v
OUTPUT
=== RUN   TestFib
=== RUN   TestFib/0-->0
    fib_test.go:92: Time Taken By BruteForce: 125ns
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 144ns
=== RUN   TestFib/1-->1
    fib_test.go:92: Time Taken By BruteForce: 81ns
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 153ns
=== RUN   TestFib/2-->1
    fib_test.go:92: Time Taken By BruteForce: 70ns
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 121ns
=== RUN   TestFib/3-->2
    fib_test.go:92: Time Taken By BruteForce: 97ns
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 343ns
=== RUN   TestFib/10-->55
    fib_test.go:92: Time Taken By BruteForce: 661ns
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 1.123µs
=== RUN   TestFib/46-->1836311903
    fib_test.go:92: Time Taken By BruteForce: 6.87384048s
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 36.791µs
=== RUN   TestFib/50-->12586269025
    fib_test.go:92: Time Taken By BruteForce: 47.09437934s
    fib_test.go:100: Time Taken By Memoized(optimized) algorithm: 20.461µs
--- PASS: TestFib (53.97s)


*/
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
		{
			input:          50,
			expectedOutput: 12586269025,
		},
	}

	for _, test := range tests {

		t.Run(fmt.Sprintf("%d-->%d", test.input, test.expectedOutput), func(t *testing.T) {
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
			t.Logf("Time Taken By BruteForce approach: %v", taken)
			
			start = time.Now()

			out = fibOptimized(test.input,nil)
			 end = time.Now()
			taken = end.Sub(start)
			assert.Equal(t, test.expectedOutput, out)
			t.Logf("Time Taken By Memoized(optimized) algorithm: %v", taken)

		})

	}

}
