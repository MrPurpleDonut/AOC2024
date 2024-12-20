package main

import (
	"fmt"
	"math"
	"os"
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
	visited[end.y][end.x] = true
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

func mattDistance(a, b cord) int {
	val1 := a.x - b.x
	val2 := a.y - b.y
	if val1 < 0 {
		val1 *= -1
	}
	if val2 < 0 {
		val2 *= -1
	}
	return val1 + val2
}

func calculateCost(p1, p2, extra, total int) int {
	if p1 <= p2 {
		return math.MaxInt16
	}
	return total - p1 + p2 + extra
}

func getVal(dist, val, threshold int, cache map[cord]int) int {
	count := 0
	for k := range cache {
		for j := range cache {
			d := mattDistance(k, j)
			if d == 1 || d > dist {
				continue
			}
			if calculateCost(cache[k], cache[j], d, val)+threshold <= val {
				count++
			}

		}
	}
	return count
}

func main() {
	s := time.Now()

	grid, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)
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
	fmt.Println(getVal(2, val, 100, cache))
	fmt.Println(getVal(20, val, 100, cache))

	fmt.Println(time.Since(s))
}
