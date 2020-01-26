package algorithms

// https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[row-1][col-1]
}
