package main

import (
	"fmt"
	"os"
	"slices"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type cord struct {
	x int
	y int
}

type point struct {
	x    int
	y    int
	cost int
}

type direction struct {
	dx int
	dy int
}

type size struct {
	x int
	y int
}

func getNeighbors(p point, s size) []point {
	directions := []direction{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	neighbors := make([]point, 0)
	for _, dir := range directions {
		if p.x+dir.dx >= 0 && p.x+dir.dx < s.x && p.y+dir.dy >= 0 && p.y+dir.dy < s.y {
			neighbors = append(neighbors, point{p.x + dir.dx, p.y + dir.dy, p.cost + 1})
		}
	}
	return neighbors
}

func solve(start, end point, grid [][]rune) int {
	visited := aoc.MakeBoolBoard(len(grid), len(grid[0]))
	s := size{len(grid), len(grid[0])}
	queue := make([]point, 1)
	queue[0] = start
	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]
		if curPos.x == end.x && curPos.y == end.y {
			return curPos.cost
		}
		for _, n := range getNeighbors(curPos, s) {
			if visited[n.y][n.x] {
				continue
			}
			if grid[n.y][n.x] != '.' {
				continue
			}
			visited[n.y][n.x] = true
			queue = append(queue, n)
		}

	}
	return -1
}

func populateCache(end point, grid [][]rune) map[cord]int {
	cache := make(map[cord]int)
	visited := aoc.MakeBoolBoard(len(grid), len(grid[0]))
	s := size{len(grid), len(grid[0])}
	queue := make([]point, 1)
	queue[0] = end
	for len(queue) > 0 {
		curPos := queue[0]
		cache[cord{curPos.x, curPos.y}] = curPos.cost
		queue = queue[1:]
		for _, n := range getNeighbors(curPos, s) {
			if visited[n.y][n.x] {
				continue
			}
			if grid[n.y][n.x] != '.' {
				continue
			}
			visited[n.y][n.x] = true
			queue = append(queue, n)
		}

	}
	return cache

}

func solveWithCheat(start, end point, grid [][]rune, dir direction, move int) int {
	visited := aoc.MakeBoolBoard(len(grid), len(grid[0]))
	s := size{len(grid), len(grid[0])}
	queue := make([]point, 1)
	queue[0] = start
	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]
		if curPos.x == end.x && curPos.y == end.y {
			return curPos.cost
		}
		if curPos.cost == move {
			n := point{curPos.x + 2*dir.dx, curPos.y + 2*dir.dy, curPos.cost + 2}
			if n.x >= 0 && n.x < s.x && n.y >= 0 && n.y < s.y {
				if visited[n.y][n.x] {
					continue
				}
				if grid[n.y][n.x] != '.' {
					continue
				}
				queue = append(queue, n)
			}
		} else {
			for _, n := range getNeighbors(curPos, s) {
				if visited[n.y][n.x] {
					continue
				}
				if grid[n.y][n.x] != '.' {
					continue
				}
				visited[n.y][n.x] = true
				queue = append(queue, n)
			}
		}
		slices.SortFunc(queue, func(a, b point) int { return a.cost - b.cost })

	}
	return -1
}

func main() {
	s := time.Now()

	grid, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)
	count := 0
	var start, end point
	for i, row := range grid {
		for j, v := range row {
			if v == 'S' {
				start = point{j, i, 0}
				grid[i][j] = '.'
			}
			if v == 'E' {
				end = point{j, i, 0}
				grid[i][j] = '.'
			}

		}
	}
	val := solve(start, end, grid)
	cache := populateCache(end, grid)

	directions := []direction{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for i := range val {
		for _, d := range directions {
			if v := solveWithCheat(start, end, grid, d, i); v > 0 && v+100 <= val {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Println(time.Since(s))
}
