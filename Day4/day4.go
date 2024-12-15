package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	letters := []rune{
		'M', 'A', 'S',
	}
	offsets := [][][]int{
		{{1, 0}, {2, 0}, {3, 0}},
		{{-1, 0}, {-2, 0}, {-3, 0}},
		{{0, 1}, {0, 2}, {0, 3}},
		{{0, -1}, {0, -2}, {0, -3}},
		{{1, 1}, {2, 2}, {3, 3}},
		{{-1, -1}, {-2, -2}, {-3, -3}},
		{{1, -1}, {2, -2}, {3, -3}},
		{{-1, 1}, {-2, 2}, {-3, 3}},
	}

	data, err := os.ReadFile(os.Args[1])
	handleError(err)
	lines := strings.Split(string(data), "\n")
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'X' {
				for _, offset := range offsets {
					good := true
					for index := range 3 {
						if i+offset[index][0] < 0 ||
							j+offset[index][1] < 0 ||
							i+offset[index][0] >= len(grid) ||
							j+offset[index][1] >= len(grid[0]) ||
							grid[i+offset[index][0]][j+offset[index][1]] != letters[index] {
							good = false
						}
					}
					if good {
						count++
					}
				}
			}
		}
	}

	count2 := 0
	for i := range len(grid) - 2 {
		for j := range len(grid[0]) - 2 {
			if grid[i+1][j+1] == 'A' {
				if grid[i][j] == 'M' && grid[i+2][j] == 'M' &&
					grid[i][j+2] == 'S' && grid[i+2][j+2] == 'S' ||
					grid[i][j] == 'M' && grid[i+2][j] == 'S' &&
						grid[i][j+2] == 'M' && grid[i+2][j+2] == 'S' ||
					grid[i][j] == 'S' && grid[i+2][j] == 'M' &&
						grid[i][j+2] == 'S' && grid[i+2][j+2] == 'M' ||
					grid[i][j] == 'S' && grid[i+2][j] == 'S' &&
						grid[i][j+2] == 'M' && grid[i+2][j+2] == 'M' {
					count2++
				}
			}
		}
	}
	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
