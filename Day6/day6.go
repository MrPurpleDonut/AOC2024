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
func testIfReaches2(sX, sY int, walls [][]bool) int {
	rotations, iterations := 0, 0
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for {
		tX, tY := sX+directions[rotations%4][0], sY+directions[rotations%4][1]
		if tX < 0 || tX >= len(walls) ||
			tY < 0 || tY >= len(walls[0]) {
			return 0
		}
		if walls[tX][tY] {
			rotations++
		} else {
			sX, sY = tX, tY
		}
		iterations++
		if iterations > 15000 {
			return 1
		}

	}
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)
	lines := strings.Split(string(data), "\n")
	grid := make([][]rune, len(lines))
	walls := make([][]bool, len(lines))
	visited := make([][]int, len(lines))

	X, Y := 0, 0
	for i := range lines {
		grid[i] = []rune(lines[i])
		walls[i] = make([]bool, len(lines[i]))
		visited[i] = make([]int, len(lines[1]))
	}
	for i := range grid {
		for j, v := range grid[i] {
			switch v {
			case '^':
				X, Y = i, j
			case '#':
				walls[i][j] = true
			}
		}
	}
	sX, sY := X, Y
	count, count2, rotations := 0, 0, 0
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for {
		visited[X][Y] = 1
		tX, tY := X+directions[rotations%4][0], Y+directions[rotations%4][1]
		if tX < 0 || tX >= len(walls) ||
			tY < 0 || tY >= len(walls[0]) {
			break
		}
		if walls[tX][tY] {
			rotations++
		} else {
			X, Y = tX, tY

		}

	}
	for i := range visited {
		for j, v := range visited[i] {
			count += v
			if v > 0 && !(sX == i && sY == j) {
				walls[i][j] = true
				count2 += testIfReaches2(sX, sY, walls)
				walls[i][j] = false
			}
		}
	}

	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
