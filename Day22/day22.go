package main

import (
	"fmt"
	"os"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

var mod = 16777216

func getPrices(price []int) map[pattern]int {
	patterns := make(map[pattern]int)
	for i := 1; i < len(price)-4; i++ {
		val := pattern{price[i] - price[i-1],
			price[i+1] - price[i],
			price[i+2] - price[i+1],
			price[i+3] - price[i+2]}
		if _, ok := patterns[val]; !ok {
			patterns[val] = price[i+3]
		}

	}
	return patterns
}

type pattern struct {
	i int
	j int
	k int
	l int
}

func opperate(val, iterations int) (int, []int) {
	res := make([]int, iterations+1)
	res[0] = val % 10
	for i := range iterations {
		temp := val * 64
		val ^= temp
		val %= mod
		temp = val / 32
		val ^= temp
		val %= mod

		val ^= (val * 2048)
		val %= mod
		res[i+1] = val % 10
	}
	return val, res
}

func findMax(prices [][]int) int {
	max := 0
	priceMaps := make([]map[pattern]int, len(prices))
	for i, v := range prices {
		priceMaps[i] = getPrices(v)
	}
	for i := -9; i <= 9; i++ {
		for j := -9; j <= 9; j++ {
			for k := -9; k <= 9; k++ {
				for l := -9; l <= 9; l++ {
					sum := 0
					for _, p := range priceMaps {
						if v, ok := p[pattern{i, j, k, l}]; ok {
							sum += v
						}
					}
					if sum > max {
						max = sum
					}
				}
			}
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
