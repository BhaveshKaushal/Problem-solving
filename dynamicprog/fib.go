package dynamicprog

/*
recursive function for fibonacci series

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

	if n <= 0  {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}