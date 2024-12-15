package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func sumLine(text string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	vals := re.FindAllString(text, -1)

	sum := 0
	for _, v := range vals {
		v, _ = strings.CutPrefix(v, "mul(")
		v, _ = strings.CutSuffix(v, ")")
		vals := strings.Split(v, ",")

		num1, err := strconv.Atoi(vals[0])
		handleError(err)

		num2, err := strconv.Atoi(vals[1])
		handleError(err)
		sum += num1 * num2
	}
	return sum
}

func sumLine2(text string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	vals := re.FindAllString(text, -1)

	sum, enabled := 0, true
	for _, v := range vals {
		if v == "do()" {
			enabled = true
		} else if v == "don't()" {
			enabled = false
		} else if enabled {
			v, _ = strings.CutPrefix(v, "mul(")
			v, _ = strings.CutSuffix(v, ")")
			vals := strings.Split(v, ",")

			num1, err := strconv.Atoi(vals[0])
			handleError(err)

			num2, err := strconv.Atoi(vals[1])
			handleError(err)
			sum += num1 * num2
		}

	}
	return sum
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	handleError(err)

	text := string(data)

	fmt.Println(sumLine(text), sumLine2(text))
	fmt.Println(time.Since(start))
}
