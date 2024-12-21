package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type point struct {
	x int
	y int
}

type state struct {
	val   string
	depth int
}

func calculateNumMoves(output string) string {
	numericalMap := make(map[string]point)
	numericalMap["A"] = point{2, 3}
	numericalMap["0"] = point{1, 3}
	numericalMap["1"] = point{0, 2}
	numericalMap["2"] = point{1, 2}
	numericalMap["3"] = point{2, 2}
	numericalMap["4"] = point{0, 1}
	numericalMap["5"] = point{1, 1}
	numericalMap["6"] = point{2, 1}
	numericalMap["7"] = point{0, 0}
	numericalMap["8"] = point{1, 0}
	numericalMap["9"] = point{2, 0}

	loc := numericalMap["A"]
	result := ""
	for _, char := range output {
		dest := numericalMap[string(char)]
		xdif, ydif := dest.x-loc.x, dest.y-loc.y
		h, v := "", ""
		for range abs(xdif) {
			if xdif > 0 {
				h += ">"
			} else {
				h += "<"
			}
		}
		for range abs(ydif) {
			if ydif > 0 {
				v += "v"
			} else {
				v += "^"
			}
		}
		if loc.y == 3 && dest.x == 0 {
			result += v
			result += h
		} else if loc.x == 0 && dest.y == 3 {
			result += h
			result += v
		} else if xdif < 0 {
			result += h
			result += v
		} else {
			result += v
			result += h
		}

		loc = dest
		result += "A"
	}
	return result

}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

func calculateMoveMoves(output string) string {
	directionalMap := make(map[string]point)
	directionalMap["A"] = point{2, 0}
	directionalMap["^"] = point{1, 0}
	directionalMap["<"] = point{0, 1}
	directionalMap["v"] = point{1, 1}
	directionalMap[">"] = point{2, 1}

	loc := directionalMap["A"]
	result := ""
	for _, char := range output {
		dest := directionalMap[string(char)]
		xdif, ydif := dest.x-loc.x, dest.y-loc.y
		h, v := "", ""
		for range abs(xdif) {
			if xdif > 0 {
				h += ">"
			} else {
				h += "<"
			}
		}
		for range abs(ydif) {
			if ydif > 0 {
				v += "v"
			} else {
				v += "^"
			}
		}
		if loc.x == 0 && dest.y == 0 {
			result += h
			result += v
		} else if dest.x == 0 && loc.y == 0 {
			result += v
			result += h
		} else if xdif < 0 {
			result += h
			result += v
		} else {
			result += v
			result += h
		}

		loc = dest
		result += "A"
	}
	return result
}

func solve(depth int, cache *map[state]int, val string) int {
	if v, ok := (*cache)[state{val, depth}]; ok {
		return v
	}
	if depth == 0 {
		return len(val)
	}
	pieces := strings.Split(val, "A")
	pieces = pieces[:len(pieces)-1]
	count := 0
	for _, p := range pieces {
		v := calculateMoveMoves(p + "A")
		count += solve(depth-1, cache, v)
	}
	(*cache)[state{val, depth}] = count
	return count
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	count, count2 := 0, 0
	cache := make(map[state]int)
	for _, l := range lines {

		num, _ := aoc.ParseAllInts(l)
		val := solve(2, &cache, calculateNumMoves(l))
		val2 := solve(25, &cache, calculateNumMoves(l))
		count += num[0] * val
		count2 += num[0] * val2
	}
	//126384
	//154115708116294

	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
