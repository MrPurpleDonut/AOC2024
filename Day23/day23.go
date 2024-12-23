package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func makeGraph(lines []string) map[string][]string {
	connections := make(map[string][]string)
	for _, l := range lines {
		vals := strings.Split(l, "-")
		if v, ok := connections[vals[0]]; ok {
			v = append(v, vals[1])
			connections[vals[0]] = v
		} else {
			connections[vals[0]] = []string{vals[1]}
		}
		if v, ok := connections[vals[1]]; ok {
			v = append(v, vals[0])
			connections[vals[1]] = v
		} else {
			connections[vals[1]] = []string{vals[0]}
		}
	}
	for k, v := range connections {
		slices.Sort(v)
		connections[k] = v
	}
	return connections
}

func findLargest(graph map[string][]string) string {
	vertices := make([]string, 0)
	for k := range graph {
		vertices = append(vertices, k)
	}
	val := bronKerbosch1([]string{}, vertices, []string{}, graph)
	slices.Sort(val)
	return strings.Join(val, ",")
}

func union(v1, v2 []string) []string {
	val := slices.Clone(v1)
	for _, v := range v2 {
		if !slices.Contains(val, v) {
			val = append(val, v)
		}
	}
	slices.Sort(val)
	return val

}

func intersect(v1, v2 []string) []string {
	val := make([]string, 0)
	for _, s1 := range v1 {
		if slices.Contains(v2, s1) {
			val = append(val, s1)
		}
	}

	slices.Sort(val)
	return val
}

func remove(v1 []string, val string) []string {
	for i, v := range v1 {
		if v == val {
			return append(v1[:i], v1[i+1:]...)
		}
	}
	return v1
}

func bronKerbosch1(r, p, x []string, graph map[string][]string) []string {
	if len(p) == 0 && len(x) == 0 {
		return r
	}
	possibilities := make([][]string, 0)
	for _, vert := range slices.Clone(p) {
		possibilities = append(possibilities,
			bronKerbosch1(union(r, []string{vert}), intersect(p, graph[vert]), intersect(x, graph[vert]), graph))
		x = union(x, []string{vert})
		p = remove(p, vert)
	}
	slices.SortFunc(possibilities, func(a, b []string) int {
		return len(b) - len(a)
	})
	if len(possibilities) > 0 {
		return possibilities[0]
	}
	return []string{}
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	graph := makeGraph(lines)
	set := make(map[string]bool)
	for key, value := range graph {
		if key[0] != 't' {
			continue
		}
		for _, v2 := range value {
			for _, v3 := range value {
				if v2 == v3 {
					continue
				}
				if !slices.Contains(graph[v2], v3) {
					continue
				}
				v := []string{key, v2, v3}
				slices.Sort(v)
				set[strings.Join(v, ",")] = true
			}

		}

	}
	fmt.Println(len(set))
	fmt.Println(findLargest(graph))
	fmt.Println(time.Since(start))
}
