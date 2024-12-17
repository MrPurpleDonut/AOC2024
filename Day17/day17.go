package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

func workBack(index, a int, instructions []int) int {
	for i := range 8 {
		v, _ := aoc.ParseAllInts(run(a*8+i, 0, 0, instructions))
		if reflect.DeepEqual(v, instructions[index:]) {
			if index == 0 {
				return a*8 + i
			}
			val := workBack(index-1, a*8+i, instructions)
			if val != -1 {
				return val
			}
		}

	}
	return -1
}

func run(a, b, c int, instructions []int) string {
	output := ""
	for i := 0; i < len(instructions); i += 2 {
		if i%2 != 0 || i+1 >= len(instructions) {
			continue
		}
		opcode, operand := instructions[i], instructions[i+1]
		switch opcode {
		case 0:
			switch operand {
			case 0, 1, 2, 3:
			case 4:
				operand = a
			case 5:
				operand = b
			case 6:
				operand = c
			}
			a = a >> operand
		case 1:
			b ^= operand
		case 2:
			switch operand {
			case 0, 1, 2, 3:
			case 4:
				operand = a % 8
			case 5:
				operand = b % 8
			case 6:
				operand = c % 8
			}
			b = operand
		case 3:
			if a == 0 {
				continue
			}
			i = operand - 2
		case 4:
			b = b ^ c
		case 5:
			switch operand {
			case 0, 1, 2, 3:
			case 4:
				operand = a
			case 5:
				operand = b
			case 6:
				operand = c
			}
			output += strconv.Itoa(operand % 8)
			output += ","
		case 6:
			switch operand {
			case 0, 1, 2, 3:
			case 4:
				operand = a
			case 5:
				operand = b
			case 6:
				operand = c
			}
			b = a >> operand
		case 7:
			switch operand {
			case 0, 1, 2, 3:
			case 4:
				operand = a
			case 5:
				operand = b
			case 6:
				operand = c
			}
			c = a >> operand
		}
	}
	ret, _ := strings.CutSuffix(output, ",")
	return ret
}

func main() {
	start := time.Now()

	lines, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)

	temp, _ := aoc.ParseAllInts(lines[0])
	a := temp[0]
	b, c := 0, 0

	instructions, _ := aoc.ParseAllInts(lines[4])
	output := run(a, b, c, instructions)

	fmt.Println(output)
	fmt.Println(workBack(len(instructions)-1, 0, instructions))
	fmt.Println(time.Since(start))
}
