package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -1 * num
}

func excludeIndex(slice []int, index int) []int {
	newSlice := []int{}
	for i, v := range slice {
		if index != i {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)

	lines := strings.Split(string(data), "\n")

	count := 0
	count2 := 0
	for _, v := range lines {
		nums := strings.Split(v, " ")
		good := true

		ints := make([]int, len(nums))
		for i, num := range nums {
			ints[i], _ = strconv.Atoi(num)
		}
		for i, v := range ints {
			if i < len(ints)-1 {
				if i > 0 {
					if v > ints[i-1] && v > ints[i+1] {
						good = false
					}
					if v < ints[i-1] && v < ints[i+1] {
						good = false
					}

				}
				if abs(v-ints[i+1]) > 3 {
					good = false
				}
				if v == ints[i+1] {
					good = false
				}
			}
		}

		if good {
			count++
			count2++
		} else {
			for index := range ints {
				isGood := true
				slice := excludeIndex(ints, index)

				for i, v := range slice {
					if i < len(slice)-1 {
						if i > 0 {
							if v > slice[i-1] && v > slice[i+1] {
								isGood = false
							}
							if v < slice[i-1] && v < slice[i+1] {
								isGood = false
							}
						}
						if abs(v-slice[i+1]) > 3 {
							isGood = false
						}
						if v == slice[i+1] {
							isGood = false
						}
					}
				}

				if isGood {
					count2++
					break
				}
			}
		}

	}
	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
