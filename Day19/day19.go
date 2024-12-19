package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func isValid(test string, towels []string, memo *map[string]int) int {
	if test == "" {
		return 1
	}
	if val, ok := (*memo)[test]; ok {
		return val
	}
	(*memo)[test] = 0
	for _, p := range towels {
		if len(test) < len(p) {
			continue
		}
		if p != test[:len(p)] {
			continue
		}
		if val := isValid(test[len(p):], towels, memo); val > 0 {
			(*memo)[test] += val

		}
	}

	return (*memo)[test]
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	patterns := strings.Split(lines[0], ", ")
	lines = lines[2:]
	count, valid := 0, 0
	memo := make(map[string]int)

	for _, p := range lines {
		val := isValid(p, patterns, &memo)
		count += val
		if val > 0 {
			valid++
		}
	}
	fmt.Println(valid, count)
	fmt.Println(time.Since(start))
}
