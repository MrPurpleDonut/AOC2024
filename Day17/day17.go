package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

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
			a = a / int(math.Pow(2, float64(operand)))
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
			b = a / int(math.Pow(2, float64(operand)))
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
			c = a / int(math.Pow(2, float64(operand)))
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

	count, good := 38045431114282, 0
	fmt.Println(run(count, 0, 0, instructions))
	for {
		val := run(count, 0, 0, instructions)
		if val[:26] == lines[4][:26] {
			fmt.Println(count, count-good, val)
			good = count
		}
		if val == lines[4] {
			break
		}
		count += 68719476736
	}

	fmt.Println(count)
	fmt.Println(time.Since(start))
}
