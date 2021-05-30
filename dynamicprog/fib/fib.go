package dynamicprog

/*
brute force recursive function for fibonacci series

			  f(n)
			 /     \
			/       \
		f(n-1)      f(n-2)
		/   \        /   \
       /     \      /     \
   f(n-2)  f(n-3)  f(n-3) f(n-4)
    .        .        .      .
   . .      . .      .  .   .  .
  .   .    .   .    .    .
f(2) f(1) f(2)  f(1)

   TIME Complexity = O(2^n)
   Space Complexity = O(n)

*/
func fib(n int64) int64 {

	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

/*
memoized recursive function for fibonacci series to optimized the solution

			  f(n)
			 /     \
			/       \
		f(n-1)      f(n-2)  --------- memoizing repeated calls
		/   \        
       /     \           
   f(n-2)  f(n-3)
    .        .        .      .
   . .      . .      .  .   .  .
  .   .    .   .    .    .
f(2) f(1) f(2)  f(1)

   TIME Complexity = O(n)
   Space Complexity = O(n)
*/
func fibOptimized(n int64, memMap map[int64]int64) int64{

	if memMap == nil {
		memMap = map[int64]int64{}
	}

	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	if val,ok := memMap[n]; ok {
		return val
	}

	memMap[n] = fibOptimized(n-1,memMap) + fibOptimized(n-2,memMap)
	return memMap[n]
}

   
