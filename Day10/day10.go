package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	row int
	col int
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func search(row, col, val int, grid [][]int, found *[]point) []point {
	if val == 9 {

		*found = append(*found, point{row: row, col: col})

	}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		if newRow >= 0 && newRow < len(grid) &&
			newCol >= 0 && newCol < len(grid) {
			if grid[newRow][newCol] == val+1 {
				search(newRow, newCol, val+1, grid, found)
			}
		}
	}

	return *found
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)
	lines := strings.Split(string(data), "\n")
	count := 0
	grid := make([][]int, len(lines))
	starts := make([]point, 0)
	ends := make([]point, 0)

	for i, line := range lines {
		grid[i] = make([]int, len(lines))
		for j, v := range []rune(line) {
			val, _ := strconv.Atoi(string(v))
			if val == 0 {
				starts = append(starts, point{i, j})
			}
			if val == 9 {
				ends = append(ends, point{i, j})
			}
			grid[i][j] = val
		}
	}
	for _, p := range starts {
		found := make([]point, 0)
		count += len(search(p.row, p.col, 0, grid, &found))
	}

	fmt.Println(count)
	fmt.Println(time.Since(start))
}
