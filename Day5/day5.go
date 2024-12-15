package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func fix(pages []int, orderings map[int][]int) []int {
	for i := range pages {
		for j := range len(pages) {
			if j > i {
				if slices.Contains(orderings[pages[j]], pages[i]) {
					temp := pages[i]
					pages[i] = pages[j]
					pages[j] = temp
					return fix(pages, orderings)
				}
			}
		}
	}
	return pages
}

func getBad(line string, orderings map[int][]int) []int {
	temp := strings.Split(line, ",")
	pages := make([]int, len(temp))
	for i, v := range temp {
		num, _ := strconv.Atoi(v)
		pages[i] = num
	}
	for i := range pages {
		for j := range len(pages) {
			if j > i {
				if slices.Contains(orderings[pages[j]], pages[i]) {
					return pages
				}
			}
		}
	}
	return nil
}

func eval(line string, orderings map[int][]int) int {
	temp := strings.Split(line, ",")
	pages := make([]int, len(temp))
	for i, v := range temp {
		num, _ := strconv.Atoi(v)
		pages[i] = num
	}
	for i := range pages {
		for j := range len(pages) {
			if j > i {
				if slices.Contains(orderings[pages[j]], pages[i]) {
					return 0
				}
			}
		}
	}
	return pages[(len(pages)-1)/2]
}

func eval2(pages []int, orderings map[int][]int) int {
	for i := range pages {
		for j := range len(pages) {
			if j > i {
				if slices.Contains(orderings[pages[j]], pages[i]) {
					return 0
				}
			}
		}
	}
	return pages[(len(pages)-1)/2]
}

func makeMap(lines []string) map[int][]int {
	ordering := make(map[int][]int)
	for _, v := range lines {
		sides := strings.Split(v, "|")
		int1, _ := strconv.Atoi(sides[0])
		int2, _ := strconv.Atoi(sides[1])
		if container, ok := ordering[int1]; ok {
			container = append(container, int2)
			ordering[int1] = container
		} else {
			ordering[int1] = []int{int2}
		}
	}
	return ordering
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)

	data1, err := os.ReadFile(os.Args[2])
	handleError(err)
	lines1 := strings.Split(string(data1), "\n")
	ordering := makeMap(lines1)

	lines := strings.Split(string(data), "\n")
	count, count1 := 0, 0
	for _, v := range lines {
		count += eval(v, ordering)
		val := getBad(v, ordering)
		if val != nil {
			count1 += eval2(fix(val, ordering), ordering)
		}
	}
	fmt.Println(count, count1)
	fmt.Println(time.Since(start))
}
