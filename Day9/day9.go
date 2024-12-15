package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type gap struct {
	start int
	end   int
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func s(vals []int) [1000000]int {
	var nums [1000000]int
	index := 0

	for i, v := range vals {
		if i%2 == 0 {
			for range v {
				nums[index] = i / 2
				index++
			}
		} else {
			index += v
		}
	}
	s := vals[0]
	gaps := make([]gap, 0)
	for s < index {
		if nums[s] == 0 {
			newGap := gap{
				start: s,
			}
			for nums[s] == 0 {
				s++
			}
			newGap.end = s - 1
			gaps = append(gaps, newGap)
		}
		s++
	}
	s = index - 1
	for s > vals[0] {
		if nums[s] != 0 {
			val := nums[s]
			end := s
			for nums[s] == val {
				s--
			}
			start := s + 1
			found := false
			for j, gap := range gaps {
				if gap.end-gap.start >= end-start && gap.end < end {
					for i := range end - start + 1 {
						nums[gap.start+i] = val
					}
					if gap.end-gap.start == end-start {
						gaps = append(gaps[:j], gaps[j+1:]...)
					} else {
						gaps[j].start += end - start + 1
					}
					found = true
					break
				}
			}
			if found {
				for i := range end - start + 1 {
					nums[s+i+1] = 0
				}
			}
		} else {
			s--
		}
	}
	return nums
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)
	vals := make([]int, 0)
	var nums [1000000]int
	for _, v := range string(data) {
		val, _ := strconv.Atoi(string(v))
		vals = append(vals, val)
	}
	vals2 := make([]int, len(vals))
	copy(vals2, vals)
	nums2 := s(vals2)
	i, j, index := 0, len(vals)-1, 0
	for i < j {
		if i%2 == 0 {
			for range vals[i] {
				nums[index] = i / 2
				index++
			}
			i++
		} else {
			openSpots := vals[i]
			for openSpots > 0 {
				if vals[j] == 0 {
					j -= 2
					if j <= i {
						break
					}
				} else {
					nums[index] = j / 2
					index++
					vals[j]--
					openSpots--
				}
			}
			i++
		}
	}

	count, count2 := 0, 0
	for i, v := range nums {
		count += i * v
	}
	for i, v := range nums2 {
		count2 += i * v
	}
	fmt.Println(count, count2)
	fmt.Println(time.Since(start))
}
