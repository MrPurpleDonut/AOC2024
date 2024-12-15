package main

import (
	"fmt"
	"math"
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

func handLine(line string) int {
	line = strings.Replace(line, ":", "", 1)
	numStrings := strings.Split(line, " ")
	nums := make([]int, len(numStrings)-1)
	counter := make([]int, len(numStrings)-2)
	sum := 0
	for i, v := range numStrings {
		num, _ := strconv.Atoi(v)
		if i == 0 {
			sum = num
		} else {
			nums[i-1] = num
		}
	}
	for {
		total := nums[0]
		for i, v := range counter {
			if v == 0 {
				total += nums[i+1]
			} else if v == 1 {
				total *= nums[i+1]
			} else {
				power := int(math.Log10(float64(nums[i+1]))) + 1
				total = total*int(math.Pow10(power)) + nums[i+1]
			}
		}
		if total == sum {
			return sum
		}
		rollOver := true
		for i, v := range counter {
			if v == 2 && rollOver {
				counter[i] = 0
			} else if rollOver {
				counter[i]++
				rollOver = false
			}
		}
		if rollOver {
			break
		}
	}

	return 0
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)
	lines := strings.Split(string(data), "\n")

	count := 0
	for _, v := range lines {
		count += handLine(v)
	}

	fmt.Println(count)
	fmt.Println(time.Since(start))
}
