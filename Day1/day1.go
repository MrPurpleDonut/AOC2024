package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)

	lines := strings.Split(string(data), "\n")
	list1, list2 := make([]int, len(lines)), make([]int, len(lines))
	for i, v := range lines {
		nums := strings.Split(v, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		list1[i] = num1
		list2[i] = num2
	}
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	for i := range list1 {
		if list1[i] > list2[i] {
			sum += list1[i] - list2[i]
		} else {
			sum += list2[i] - list1[i]
		}
	}

	fmt.Println(sum)

	counts := make(map[int]int)

	for _, v := range list2 {
		counts[v] += 1
	}
	sum2 := 0
	for _, v := range list1 {
		if val, ok := counts[v]; ok {
			sum2 += val * v
		}
	}
	fmt.Println(sum2)
	fmt.Println(time.Since(start))
}
