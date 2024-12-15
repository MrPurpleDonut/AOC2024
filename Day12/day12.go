package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func search(row, col int, val rune, searched *[][]bool, grid [][]rune) (int, int, int) {
	(*searched)[row][col] = true
	area, permiter, side := 1, 0, 0
	rows, cols := len(grid), len(grid[0])

	if (row-1 < 0 || grid[row-1][col] != val) && (col-1 < 0 || grid[row][col-1] != val) {
		side++
	}
	if (row-1 < 0 || grid[row-1][col] != val) && (col+1 >= cols || grid[row][col+1] != val) {
		side++
	}
	if (row+1 >= rows || grid[row+1][col] != val) && (col-1 < 0 || grid[row][col-1] != val) {
		side++
	}
	if (row+1 >= rows || grid[row+1][col] != val) && (col+1 >= cols || grid[row][col+1] != val) {
		side++
	}
	if row-1 >= 0 && col-1 >= 0 &&
		grid[row-1][col] == val && grid[row][col-1] == val && grid[row-1][col-1] != val {
		side++
	}
	if row-1 >= 0 && col+1 < cols &&
		grid[row-1][col] == val && grid[row][col+1] == val && grid[row-1][col+1] != val {
		side++
	}
	if row+1 < rows && col-1 >= 0 &&
		grid[row+1][col] == val && grid[row][col-1] == val && grid[row+1][col-1] != val {
		side++
	}
	if row+1 < rows && col+1 < cols &&
		grid[row+1][col] == val && grid[row][col+1] == val && grid[row+1][col+1] != val {
		side++
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		if newRow >= 0 && newRow < len(grid) &&
			newCol >= 0 && newCol < len(grid[0]) && grid[newRow][newCol] == val {
			if !(*searched)[newRow][newCol] {
				a, p, sides := search(newRow, newCol, val, searched, grid)
				area += a
				permiter += p
				side += sides
			}
		} else {
			permiter++
		}
	}

	return area, permiter, side
}

func main() {
	start := time.Now()

	grid, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)
	searched := aoc.MakeBoolBoard(len(grid), len(grid[0]))
	count := 0
	count2 := 0
	for i, row := range grid {
		for j, val := range row {
			if !searched[i][j] {
				area, permiter, sides := search(i, j, val, &searched, grid)
				count += area * permiter
				count2 += area * sides
			}
		}
	}

	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
