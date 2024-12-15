package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type button struct {
	dx int
	dy int
}

func getButton(line string) button {
	pattern := `X\+(\d+),\s*Y\+(\d+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)

	x, y := matches[1], matches[2]
	xint, _ := strconv.Atoi(x)
	yint, _ := strconv.Atoi(y)

	return button{xint, yint}

}

func getResult(line string) button {
	pattern := `X=(\d+),\s*Y=(\d+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)

	x, y := matches[1], matches[2]
	xint, _ := strconv.Atoi(x)
	yint, _ := strconv.Atoi(y)
	offset := 10000000000000
	return button{xint + offset, yint + offset}

}

func handleGame(game string) int {
	lines := strings.Split(game, "\n")
	a := getButton(lines[0])
	b := getButton(lines[1])
	c := getResult(lines[2])
	bVal := (c.dy*a.dx - c.dx*a.dy) / (b.dy*a.dx - b.dx*a.dy)
	aVal := (c.dx - bVal*b.dx) / a.dx
	if a.dx*aVal+b.dx*bVal != c.dx || a.dy*aVal+b.dy*bVal != c.dy {
		return 0
	}
	return 3*aVal + bVal
}

func main() {
	start := time.Now()

	data, err := os.ReadFile(os.Args[1])
	aoc.HandleError(err)

	games := strings.Split(string(data), "\n\n")
	count := 0

	for _, game := range games {
		count += handleGame(game)
	}
	fmt.Println(count)
	fmt.Println(time.Since(start))
}
