package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func getPrices(price []int, allPatterns *map[pattern]int) {
	patterns := make(map[pattern]bool)
	for i := 1; i < len(price)-4; i++ {
		val := pattern{price[i] - price[i-1],
			price[i+1] - price[i],
			price[i+2] - price[i+1],
			price[i+3] - price[i+2]}
		if _, ok := patterns[val]; !ok {
			(*allPatterns)[val] += price[i+3]
			patterns[val] = true
		}

	}
}

type pattern struct {
	i int
	j int
	k int
	l int
}

var mod = 16777216 - 1

func opperate(val, iterations int) (int, []int) {
	res := make([]int, iterations+1)
	res[0] = val % 10
	for i := range iterations {
		val ^= val << 6
		val &= mod

		val ^= val >> 5
		val &= mod

		val ^= val << 11
		val &= mod
		res[i+1] = val % 10
	}
	return val, res
}

func findMax(prices [][]int) int {
	max := 0
	priceMap := make(map[pattern]int)
	for _, v := range prices {
		getPrices(v, &priceMap)
	}
	for _, v := range priceMap {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	count := 0
	prices := make([][]int, len(lines))
	for i, l := range lines {
		num, _ := aoc.ParseAllInts(l)
		v, r := opperate(num[0], 2000)
		prices[i] = r
		count += v
	}
	count2 := findMax(prices)
	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
