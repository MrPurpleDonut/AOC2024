package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type point struct {
	x    int
	y    int
	cost int
}

type direction struct {
	dx int
	dy int
}

func getNeighbors(p point) []point {
	directions := []direction{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	neighbors := make([]point, 0)
	for _, dir := range directions {
		if p.x+dir.dx >= 0 && p.x+dir.dx <= 70 && p.y+dir.dy >= 0 && p.y+dir.dy <= 70 {
			neighbors = append(neighbors, point{p.x + dir.dx, p.y + dir.dy, p.cost + 1})
		}
	}
	return neighbors
}

func search(num int, lines []string) bool {
	unsafe := aoc.MakeBoolBoard(71, 71)
	for i := range num {
		v := lines[i]
		nums, _ := aoc.ParseAllInts(v)
		unsafe[nums[1]][nums[0]] = true
	}

	visited := aoc.MakeBoolBoard(71, 71)
	visited[0][0] = true
	queue := make([]point, 1)
	queue[0] = point{0, 0, 0}
	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]
		if curPos.x == 70 && curPos.y == 70 {
			return true
		}
		for _, n := range getNeighbors(curPos) {
			if unsafe[n.y][n.x] {
				continue
			}
			if visited[n.y][n.x] {
				continue
			}
			visited[n.y][n.x] = true
			queue = append(queue, n)
		}

	}
	return false
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	unsafe := aoc.MakeBoolBoard(71, 71)
	for i := range 1024 {
		v := lines[i]
		nums, _ := aoc.ParseAllInts(v)
		unsafe[nums[1]][nums[0]] = true
	}

	visited := aoc.MakeBoolBoard(71, 71)
	visited[0][0] = true
	queue := make([]point, 1)
	queue[0] = point{0, 0, 0}
	for len(queue) > 0 {
		curPos := queue[0]
		queue = queue[1:]
		if curPos.x == 70 && curPos.y == 70 {
			fmt.Println(curPos.cost)
			break
		}
		for _, n := range getNeighbors(curPos) {
			if unsafe[n.y][n.x] {
				continue
			}
			if visited[n.y][n.x] {
				continue
			}
			visited[n.y][n.x] = true
			queue = append(queue, n)
		}

	}

	count := 1500
	for {
		if !search(count, lines) {
			break
		}
		count++
	}
	fmt.Println(count)
	fmt.Println(time.Since(start))
}
