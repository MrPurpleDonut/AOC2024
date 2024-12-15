package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func handleLine(line string) int {
	width, height, steps := 101, 103, 100
	var x, y, dx, dy int
	fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &dx, &dy)
	x = mod((x + steps*dx), width)
	y = mod((y + steps*dy), height)
	if x == (width-1)/2 || y == (height-1)/2 {
		return -1
	}
	if x > (width-1)/2 && y > (height-1)/2 {
		return 0
	}
	if x < (width-1)/2 && y > (height-1)/2 {
		return 1
	}
	if x < (width-1)/2 {
		return 2
	}
	return 3

}

func printBoard(board [][]bool) {
	for i, row := range board {
		val := ""
		for j := range row {
			if board[i][j] {
				val += "■"
			} else {
				val += " "
			}
		}
		fmt.Println(val)
	}

}

func handleAllLines(cords []string) int {
	robots := make([]robot, len(cords))
	for i, v := range cords {
		var x, y, dx, dy int
		fmt.Sscanf(v, "p=%d,%d v=%d,%d", &x, &y, &dx, &dy)
		robots[i] = robot{x, y, dx, dy}
	}
	count := 0
	for {
		board := aoc.MakeBoolBoard(103, 101)
		allUnique := true
		for _, r := range robots {
			if board[mod(r.y+r.dy*count, 103)][mod(r.x+r.dx*count, 101)] {
				allUnique = false
				break
			} else {
				board[mod(r.y+r.dy*count, 103)][mod(r.x+r.dx*count, 101)] = true
			}
		}
		if allUnique {
			printBoard(board)
			var good string
			fmt.Println("This good? (y/n)")
			fmt.Scan(&good)
			if good == "y" {
				return count
			}
		}
		count++
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	start := time.Now()
	fmt.Println(-1 % 6)
	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)

	var quad [4]int
	for _, v := range lines {
		if r := handleLine(v); r != -1 {
			quad[r]++
		}
	}
	count := 1
	fmt.Println(quad)
	for _, v := range quad {
		count *= v
	}
	fmt.Println(count)
	fmt.Println(handleAllLines(lines))
	fmt.Println(time.Since(start))
}
