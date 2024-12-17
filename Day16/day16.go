package main

import (
	"fmt"
	"math"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type direction struct {
	dx int
	dy int
}

type point struct {
	x int
	y int
}

type node struct {
	p   point
	dir direction
}

type state struct {
	curNode node
	path    []point
	score   int
}

func getNeighbors(n node) []node {
	directions := []direction{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	neighbors := make([]node, 0)
	for _, dir := range directions {
		if n.dir.dx == -dir.dx && n.dir.dy == -dir.dy {
			continue
		}
		neighbors = append(neighbors, node{point{n.p.x + dir.dx, n.p.y + dir.dy}, dir})
	}
	return neighbors
}

func search(startX, startY, endX, endY int, grid [][]rune) (int, int) {

	endPoint := point{endX, endY}
	minScore := math.MaxInt
	curNode := node{point{startX, startY}, direction{1, 0}}
	queue := []state{
		{curNode, []point{{startX, startY}}, 0},
	}
	visited := make(map[node]int)

	costToPoint := make(map[int][]point)

	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]
		if curState.score > minScore {
			continue
		}
		if curState.curNode.p == endPoint {
			minScore = curState.score
			costToPoint[minScore] = append(costToPoint[minScore], curState.path...)
			continue
		}
		for _, neighbor := range getNeighbors(curState.curNode) {
			if grid[neighbor.p.y][neighbor.p.x] == '#' {
				continue
			}
			newScore := curState.score + 1
			if curState.curNode.dir != neighbor.dir {
				newScore += 1000
			}
			if prevScore, ok := visited[neighbor]; ok {
				if prevScore < newScore {
					continue
				}
			}
			visited[neighbor] = newScore
			newPath := make([]point, len(curState.path))
			copy(newPath, curState.path)

			queue = append(queue, state{
				neighbor,
				append(newPath, neighbor.p),
				newScore,
			})

		}

	}
	countMap := make(map[point]bool)
	for _, index := range costToPoint[minScore] {
		countMap[index] = true
	}
	return minScore, len(countMap)
}

func main() {
	start := time.Now()

	grid, err := aoc.MakeMatrix(os.Args[1])
	aoc.HandleError(err)

	startX, startY, endX, endY := 0, 0, 0, 0
	for i, row := range grid {
		for j, v := range row {
			if v == 'S' {
				startX = j
				startY = i
				grid[i][j] = '.'
			}
			if v == 'E' {
				endX = j
				endY = i
				grid[i][j] = '.'
			}
		}
	}

	fmt.Println(search(startX, startY, endX, endY, grid))
	fmt.Println(time.Since(start))
}
