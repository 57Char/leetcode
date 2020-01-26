package algorithms

import "strconv"

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var rowSet, colSet, subBoxSet []map[byte]bool
	for i := 0; i < 9; i++ {
		rowSet = append(rowSet, map[byte]bool{})
		colSet = append(colSet, map[byte]bool{})
		subBoxSet = append(subBoxSet, map[byte]bool{})
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val != '.' {
				if rowSet[i][val] {
					return false
				}
				rowSet[i][val] = true
				if colSet[j][val] {
					return false
				}
				colSet[j][val] = true
				k := (i / 3) * 3 + j / 3
				if subBoxSet[k][val] {
					return false
				}
				subBoxSet[k][val] = true
			}
		}
	}
	return true
}

func isValidSudokuV2(board [][]byte) bool {
	set := map[string]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val != '.' {
				v := "(" + string(val) + ")"
				// row
				r := v + strconv.Itoa(i)
				// col
				c := strconv.Itoa(j) + v
				// sub box
				s := strconv.Itoa(i/3) + v + strconv.Itoa(j/3)
				if set[r] || set[c] || set[s] {
					return false
				}
				set[r], set[c], set[s] = true, true, true
			}
		}
	}
	return true
}
