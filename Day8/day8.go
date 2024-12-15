package main

import (
	"fmt"
	"os"
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
	units := make(map[rune][][]int, 0)

	grid := make([][]rune, len(lines))
	seen := make([][2]int, 0)
	for i, v := range lines {
		grid[i] = []rune(v)
		for j, val := range v {
			if val != '.' {
				seen = append(seen, [2]int{j, i})
				if points, ok := units[val]; ok {
					points = append(points, []int{i, j})
					units[val] = points
				} else {
					units[val] = [][]int{{i, j}}
				}
			}
		}
	}
	size := len(grid)
	for _, points := range units {
		for _, point1 := range points {
			for _, point2 := range points {
				if point1[0] == point2[0] && point1[1] == point2[1] {
					continue
				}
				xDif, yDif := point1[1]-point2[1], point1[0]-point2[0]
				point1Y, point1X := point1[0]+yDif, point1[1]+xDif
				for !(point1Y < 0 || point1Y >= size) && !(point1X < 0 || point1X >= size) {
					contains := false
					for _, v := range seen {
						if v[0] == point1X && v[1] == point1Y {
							contains = true
						}
					}
					if !contains {
						seen = append(seen, [2]int{point1X, point1Y})
					}
					point1X += xDif
					point1Y += yDif
				}
			}
		}
	}

	fmt.Println(len(seen))
	fmt.Println(time.Since(start))
}
