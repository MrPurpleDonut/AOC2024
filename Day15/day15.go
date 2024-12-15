package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type point struct {
	y int
	x int
}

func findAbove(grid [][]rune, y, x int) []point {
	points := make([]point, 0)
	if grid[y-1][x] != '[' && grid[y-1][x] != ']' {
		return points
	}
	if grid[y-1][x] == '[' {
		points = append(points, point{y - 1, x})
		points1 := findAbove(grid, y-1, x)
		points2 := findAbove(grid, y-1, x+1)
		points = append(points, points1...)
		points = append(points, points2...)
	} else {
		points = append(points, point{y - 1, x - 1})
		points1 := findAbove(grid, y-1, x-1)
		points2 := findAbove(grid, y-1, x)
		points = append(points, points1...)
		points = append(points, points2...)
	}

	return points
}

func findBelow(grid [][]rune, y, x int) []point {
	points := make([]point, 0)
	if grid[y+1][x] != '[' && grid[y+1][x] != ']' {
		return points
	}
	if grid[y+1][x] == '[' {
		points = append(points, point{y + 1, x})
		points1 := findBelow(grid, y+1, x)
		points2 := findBelow(grid, y+1, x+1)
		points = append(points, points1...)
		points = append(points, points2...)
	} else {
		points = append(points, point{y + 1, x - 1})
		points1 := findBelow(grid, y+1, x-1)
		points2 := findBelow(grid, y+1, x)
		points = append(points, points1...)
		points = append(points, points2...)
	}

	return points
}

func countBoxes(grid [][]rune) int {
	sum := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 'O' || val == '[' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func expandGrid(grid [][]rune) ([][]rune, int, int) {
	var y, x int
	newgrid := make([][]rune, len(grid))
	for i, row := range grid {
		newgrid[i] = make([]rune, len(row)*2)
		for j, val := range row {
			if val == '#' {
				newgrid[i][2*j] = '#'
				newgrid[i][2*j+1] = '#'
			} else if val == 'O' {
				newgrid[i][2*j] = '['
				newgrid[i][2*j+1] = ']'
			} else if val == '.' {
				newgrid[i][2*j] = '.'
				newgrid[i][2*j+1] = '.'
			} else if val == '@' {
				newgrid[i][2*j] = '@'
				newgrid[i][2*j+1] = '.'
				y, x = i, 2*j
			}
		}
	}
	return newgrid, y, x
}

func main() {
	start := time.Now()

	grid, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)
	newgrid, newY, newX := expandGrid(grid)
	var y, x int
	for i, row := range grid {
		for j, val := range row {
			if val == '@' {
				y, x = i, j
			}
		}
	}

	data, err := os.ReadFile(os.Args[2])
	aoc.HandleError(err)
	moves := string(data)
	for _, move := range moves {
		switch move {
		case '\n':
			continue
		case '<':
			if grid[y][x-1] == '#' {
				continue
			} else if grid[y][x-1] == '.' {
				grid[y][x-1] = '@'
				grid[y][x] = '.'
				x -= 1
			} else {
				curX, curY := x-1, y
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curX -= 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					grid[curY][curX] = 'O'
					grid[y][x] = '.'
					grid[y][x-1] = '@'
					x -= 1
				}
			}
		case '^':
			if grid[y-1][x] == '#' {
				continue
			} else if grid[y-1][x] == '.' {
				grid[y-1][x] = '@'
				grid[y][x] = '.'
				y -= 1
			} else {
				curX, curY := x, y-1
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curY -= 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					grid[curY][curX] = 'O'
					grid[y][x] = '.'
					grid[y-1][x] = '@'
					y -= 1
				}
			}
		case '>':
			if grid[y][x+1] == '#' {
				continue
			} else if grid[y][x+1] == '.' {
				grid[y][x+1] = '@'
				grid[y][x] = '.'
				x += 1
			} else {
				curX, curY := x+1, y
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curX += 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					grid[curY][curX] = 'O'
					grid[y][x] = '.'
					grid[y][x+1] = '@'
					x += 1
				}
			}
		case 'v':
			if grid[y+1][x] == '#' {
				continue
			} else if grid[y+1][x] == '.' {
				grid[y+1][x] = '@'
				grid[y][x] = '.'
				y += 1
			} else {
				curX, curY := x, y+1
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curY += 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					grid[curY][curX] = 'O'
					grid[y][x] = '.'
					grid[y+1][x] = '@'
					y += 1
				}
			}
		}
	}

	y, x, grid = newY, newX, newgrid
	for _, move := range moves {
		switch move {
		case '\n':
			continue
		case '<':
			if grid[y][x-1] == '#' {
				continue
			} else if grid[y][x-1] == '.' {
				grid[y][x-1] = '@'
				grid[y][x] = '.'
				x -= 1
			} else {
				curX, curY := x-1, y
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curX -= 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					for i := range x - curX - 1 {
						if i%2 == 0 {
							grid[curY][curX+i] = '['
						} else {
							grid[curY][curX+i] = ']'
						}
					}
					grid[y][x-1] = '@'
					grid[y][x] = '.'
					x -= 1
				}
			}
		case '^':
			if grid[y-1][x] == '#' {
				continue
			} else if grid[y-1][x] == '.' {
				grid[y-1][x] = '@'
				grid[y][x] = '.'
				y -= 1
			} else {
				points := findAbove(grid, y, x)
				legal := true
				for _, point := range points {
					if grid[point.y-1][point.x] == '#' || grid[point.y-1][point.x+1] == '#' {
						legal = false
					}
				}
				if !legal {
					continue
				}
				for _, point := range points {
					grid[point.y][point.x] = '.'
					grid[point.y][point.x+1] = '.'
				}
				for _, point := range points {
					grid[point.y-1][point.x] = '['
					grid[point.y-1][point.x+1] = ']'
				}
				grid[y][x] = '.'
				grid[y-1][x] = '@'
				y -= 1
			}
		case '>':
			if grid[y][x+1] == '#' {
				continue
			} else if grid[y][x+1] == '.' {
				grid[y][x+1] = '@'
				grid[y][x] = '.'
				x += 1
			} else {
				curX, curY := x+1, y
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curX += 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					for i := range curX - x - 1 {
						if i%2 == 0 {
							grid[curY][curX-i] = ']'
						} else {
							grid[curY][curX-i] = '['
						}
					}
					grid[y][x+1] = '@'
					grid[y][x] = '.'
					x += 1
				}
			}
		case 'v':
			if grid[y+1][x] == '#' {
				continue
			} else if grid[y+1][x] == '.' {
				grid[y+1][x] = '@'
				grid[y][x] = '.'
				y += 1
			} else {
				curX, curY := x, y+1
				for grid[curY][curX] != '#' && grid[curY][curX] != '.' {
					curY += 1
				}
				if grid[curY][curX] == '#' {
					continue
				} else {
					points := findBelow(grid, y, x)
					legal := true
					for _, point := range points {
						if grid[point.y+1][point.x] == '#' || grid[point.y+1][point.x+1] == '#' {
							legal = false
						}
					}
					if !legal {
						continue
					}
					for _, point := range points {
						grid[point.y][point.x] = '.'
						grid[point.y][point.x+1] = '.'
					}
					for _, point := range points {
						grid[point.y+1][point.x] = '['
						grid[point.y+1][point.x+1] = ']'
					}
					grid[y][x] = '.'
					grid[y+1][x] = '@'
					y += 1
				}
			}
		}
	}
	fmt.Println(countBoxes(grid))
	fmt.Println(time.Since(start))
}
