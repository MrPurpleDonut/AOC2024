package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func intializeReg(reg string) map[string]bool {
	registers := strings.Split(reg, "\n")
	vals := make(map[string]bool)
	for _, l := range registers {
		parts := strings.Split(l, ": ")
		vals[parts[0]] = parts[1] == "1"
	}
	return vals
}

type cycle struct {
	lastI int
	cycle int
}

func runInstructions(registers map[string]bool, instructions []string) map[string]bool {
	cycleDetection := make(map[string]cycle)
	for i := 0; i < len(instructions); i++ {
		ins := instructions[i]
		comp := strings.Split(ins, " ")
		if _, ok := registers[comp[0]]; !ok {
			instructions = append(instructions, ins)
			if v, ok := cycleDetection[ins]; !ok {
				cycleDetection[ins] = cycle{i, -1}
			} else {
				if v.cycle == i-v.lastI {
					return map[string]bool{}
				}
				cycleDetection[ins] = cycle{i, i - v.lastI}
			}
			continue
		}
		if _, ok := registers[comp[2]]; !ok {
			instructions = append(instructions, ins)
			if v, ok := cycleDetection[ins]; !ok {
				cycleDetection[ins] = cycle{i, -1}
			} else {
				if v.cycle == i-v.lastI {
					return map[string]bool{}
				}
				cycleDetection[ins] = cycle{i, i - v.lastI}
			}
			continue
		}
		switch comp[1] {
		case "AND":
			registers[comp[4]] = registers[comp[0]] && registers[comp[2]]
		case "OR":
			registers[comp[4]] = registers[comp[0]] || registers[comp[2]]
		case "XOR":
			registers[comp[4]] = (registers[comp[0]] || registers[comp[2]]) && !(registers[comp[0]] && registers[comp[2]])
		}
	}
	return registers
}

func findRegisters(registers map[string]bool, instructions []string) string {
	x := 0
	for k, v := range registers {
		if k[0] == 'x' && v {
			num, _ := aoc.ParseAllInts(k)
			x += 1 << num[0]
		}

	}
	y := 0
	for k, v := range registers {
		if k[0] == 'y' && v {
			num, _ := aoc.ParseAllInts(k)
			x += 1 << num[0]
		}

	}
	cache := make(map[string]bool)
	fmt.Println("Searching for: ", x+y)
	for i, ins1 := range instructions {
		for j, ins2 := range instructions {
			for k, ins3 := range instructions {
				for l, ins4 := range instructions {
					s1 := []string{ins1, ins2, ins3, ins4}
					slices.Sort(s1)
					s2 := strings.Join(s1, "")
					if _, ok := cache[s2]; ok {
						continue
					}
					cache[s2] = true
					if i == j || i == k || i == l || j == k || j == l || k == l {
						continue
					}

					ins := slices.Clone(instructions)
					ins[i] = ins2[:len(ins2)-3] + ins1[len(ins1)-3:]
					ins[j] = ins1[:len(ins1)-3] + ins2[len(ins2)-3:]

					ins[k] = ins4[:len(ins4)-3] + ins3[len(ins3)-3:]
					ins[l] = ins3[:len(ins3)-3] + ins4[len(ins4)-3:]
					reg := runInstructions(maps.Clone(registers), ins)
					z := 0
					for k, v := range reg {
						if k[0] == 'z' && v {
							num, _ := aoc.ParseAllInts(k)
							z += 1 << num[0]
						}

					}
					if z == x+y {
						v1 := strings.Split(ins1, " ")
						v2 := strings.Split(ins2, " ")
						v3 := strings.Split(ins3, " ")
						v4 := strings.Split(ins4, " ")
						vals := []string{v1[0], v1[2], v2[0], v2[2], v3[0], v3[2], v4[0], v4[2]}
						slices.Sort(vals)
						return strings.Join(vals, ",")
					}
				}
			}
		}
	}
	return "ERROR"
}

func main() {
	start := time.Now()

	data, err := os.ReadFile(os.Args[1])
	aoc.HandleError(err)
	components := strings.Split(string(data), "\n\n")
	registers := intializeReg(components[0])
	instructions := strings.Split(components[1], "\n")
	fmt.Println(findRegisters(registers, instructions))

	registers = runInstructions(registers, instructions)
	count := 0
	for k, v := range registers {
		if k[0] == 'z' && v {
			num, _ := aoc.ParseAllInts(k)
			count += 1 << num[0]
		}

	}
	fmt.Println(count)
	fmt.Println(time.Since(start))
}
