package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func main() {
	start := time.Now()

	data, err := os.ReadFile(os.Args[1])
	aoc.HandleError(err)
	s := strings.Split(string(data), " ")
	counts := make(map[int]int)

	for _, v := range s {
		num, err := strconv.Atoi(v)
		aoc.HandleError(err)
		counts[num] = 1
	}
	for range 75 {
		newCounts := make(map[int]int)
		for key, val := range counts {
			if key == 0 {
				newCounts[1] += val
			} else if length := len(strconv.Itoa(key)); length%2 == 0 {
				val1 := key % int(math.Pow10(length/2))
				val2 := key / int(math.Pow10(length/2))
				newCounts[val1] += val
				newCounts[val2] += val
			} else {
				newCounts[key*2024] += val
			}
		}
		counts = newCounts
	}

	count := 0
	for _, v := range counts {
		count += v
	}
	fmt.Println(count, len(counts))
	fmt.Println(time.Since(start))
}
