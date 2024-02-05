package main

import (
	"fmt"
	"os"
	"strconv"
)

type SudokuGrid [9][9]int

func sudokuSolver(grid *SudokuGrid) bool {
	row, col := isCellEmpty(grid)
	if row == -1 && col == -1 {
		return true
	}
	for num := 1; num <= 9; num++ {
		if isNumValid(grid, row, col, num) {
			grid[row][col] = num
			if sudokuSolver(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}

	return false
}

func isCellEmpty(grid *SudokuGrid) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNumValid(grid *SudokuGrid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func main() {
	var grid SudokuGrid
	if len(os.Args) != 10 {
		fmt.Println("Error")
		fmt.Println()
		return
	}
	for i := 1; i < len(os.Args); i++ {
		input := os.Args[i]

		if len(input) != 9 {
			fmt.Println("Error")
			fmt.Println()
			return
		}

		for j, ch := range input {
			if ch == '.' {
				grid[i-1][j] = 0
			} else {
				num, _ := strconv.Atoi(string(ch))
				if num < 1 || num > 9 {
					fmt.Println("Error")
					return
				}
				grid[i-1][j] = num
			}
		}
	}
	if sudokuSolver(&grid) {
		printGrid(&grid)
	} else {
		fmt.Println("Error")
	}
	fmt.Println()
}

func printGrid(grid *SudokuGrid) {
	for i, row := range grid {
		for j, num := range row {
			fmt.Print(num)
			if j < len(row)-1 {
				fmt.Print(" ")
			}
		}
		if i < len(grid)-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}
