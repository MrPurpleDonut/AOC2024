package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func makeLockKeys(board []string) ([][]int, [][]int) {
	locks := make([][]int, 0)
	keys := make([][]int, 0)

	for _, b := range board {
		v := MakeMatrix(b)
		counts := make([]int, 5)
		for _, row := range v {
			for j, val := range row {
				if val == '#' {
					counts[j]++
				}
			}
		}
		if v[0][0] == '#' {
			locks = append(locks, counts)
		} else {
			keys = append(keys, counts)
		}
	}

	return locks, keys
}

func MakeMatrix(data string) [][]rune {

	vals := strings.Split(data, "\n")

	rows := make([][]rune, len(vals))
	for i, row := range vals {
		rows[i] = []rune(row)
	}
	return rows
}

func findValid(locks, keys [][]int) int {
	count := 0

	for _, lock := range locks {
		for _, key := range keys {
			valid := true
			for i := range 5 {
				if lock[i]+key[i] > 7 {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}

	return count
}

func main() {
	start := time.Now()

	data, err := os.ReadFile(os.Args[1])
	aoc.HandleError(err)

	board := strings.Split(string(data), "\n\n")
	locks, keys := makeLockKeys(board)
	fmt.Println(findValid(locks, keys))
	fmt.Println(time.Since(start))
}
