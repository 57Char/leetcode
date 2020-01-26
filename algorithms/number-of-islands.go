package algorithms

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	lenRow, lenCol := len(grid), len(grid[0])
	res := 0
	for r := 0; r < lenRow; r++ {
		for c := 0; c < lenCol; c++ {
			if grid[r][c] == '1' {
				res++
				dfs(&grid, r, c) // pointer
				//dfsV2(&grid, r, c)
			}
		}
	}
	return res
}

func dfs(grid *[][]byte, r, c int) {
	lenRow, lenCol := len(*grid), len((*grid)[0])
	if r < 0 || r >= lenRow || c < 0 || c >= lenCol || (*grid)[r][c] == '0' {
		return
	}
	(*grid)[r][c] = '0'
	dfs(grid, r - 1, c)
	dfs(grid, r + 1, c)
	dfs(grid, r, c - 1)
	dfs(grid, r, c + 1)
}

func dfsV2(grid [][]byte, r, c int) {
	lenRow, lenCol := len(grid), len(grid[0])
	if r < 0 || r >= lenRow || c < 0 || c >= lenCol || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfsV2(grid, r - 1, c)
	dfsV2(grid, r + 1, c)
	dfsV2(grid, r, c - 1)
	dfsV2(grid, r, c + 1)
}

