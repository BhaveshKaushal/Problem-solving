package dynamicprog

/*
Print all possible paths from top left to bottom right of a mXn matrix
The problem is to print all the possible paths from top left to bottom right of a mXn matrix with the constraints that from each cell you can either move only to right or down.

Examples :

Input : 1 2 3
        4 5 6
Output : 1 4 5 6
         1 2 5 6
         1 2 3 6

Input : 1 2
        3 4
Output : 1 2 4
         1 3 4
*/
type grid [][]int

var (
	out grid
)

func possiblePaths(ar [][]int, m int, n int) [][]int {
	//fmt.Printf("\n0 out=%v",out)
	out = grid{}
	possiblePathsRecursive(ar, 0, 0, m, n, []int{})
	return out
}

func possiblePathsRecursive(ar [][]int, startM int, startN int, m int, n int, path []int) {

	if startM == m-1 && startN == n-1 {

		path = append(path, ar[startM][startN])
		out = append(out, path)

		return
	}

	path = append(path, ar[startM][startN])
	if startM+1 < m {
		possiblePathsRecursive(ar, startM+1, startN, m, n, path)
	}

	if startN+1 < n {
		possiblePathsRecursive(ar, startM, startN+1, m, n, path)
	}

}
