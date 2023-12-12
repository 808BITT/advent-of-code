package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	// fmt.Println("Part 1:", part1("input.txt"))
	// fmt.Println("Time: ", time.Since(timer))
	fmt.Println("Part 2:", part2("sample.txt"))
	fmt.Println("Time: ", time.Since(timer))
}

func part2(filename string) float64 {
	lines := readFile(filename)
	var sum float64
	for _, line := range lines {
		val := evaluatePart2(line)
		sum += val
		fmt.Println(val)
		// fmt.Scanln()
	}
	return sum
}

func evaluatePart2(line string) float64 {
	split := strings.Split(line, " ")
	arrangement := split[0]
	var values []int
	for _, val := range strings.Split(split[1], ",") {
		values = append(values, atoi(val))
	}

	// var combosPerValue []float64
	for i, val := range values {
		leftValues := values[:i]
		rightValues := values[i+1:]
		fmt.Println(arrangement, values, val, "Left:", leftValues, "Right:", rightValues)
		subArrangement := arrangement[sum(leftValues) : sum(leftValues)+val+1]
		fmt.Println(subArrangement)
		fmt.Scanln()
	}

	return float64(0)
}

func sum(values []int) int {
	out := 0
	for _, val := range values {
		out += val
	}
	return out
}

func reverseArrangement(arrangement string) string {
	out := ""
	for i := len(arrangement) - 1; i >= 0; i-- {
		out += string(arrangement[i])
	}
	return out
}

func reverseValues(values []string) string {
	out := ""
	for i := len(values) - 1; i >= 0; i-- {
		out += values[i] + ","
	}
	return out[:len(out)-1]
}

func atoi(str string) int {
	out := 0
	for _, char := range str {
		out = out*10 + int(char-'0')
	}
	return out
}

func part2Expand(line string) float64 {
	original := evaluatePart1(line)
	split := strings.Split(line, " ")
	arrangement := split[0]

	var builder strings.Builder
	builder.WriteString(arrangement)
	builder.WriteString("?")
	builder.WriteString(arrangement)
	builder.WriteString(" ")
	builder.WriteString(split[1])
	builder.WriteString(",")
	builder.WriteString(split[1])

	modded := builder.String()
	doubled := evaluatePart1(modded)
	factor := float64(doubled) / float64(original)
	out := float64(original) * math.Pow(factor, 4)

	return out
}

func part1(filename string) int {
	lines := readFile(filename)
	sum := 0
	for _, line := range lines {
		val := evaluatePart1(line)
		sum += val
		fmt.Println(val)
	}
	return sum
}

func evaluatePart1(line string) int {
	split := strings.Split(line, " ")
	arrangement := split[0]
	values := strings.Split(split[1], ",")
	// fmt.Println(arrangement, values)

	//find all indices of ? in arrangement
	// for each index, replace with . and # and evaluate
	unknowns := []int{}
	for i, char := range arrangement {
		if char == '?' {
			unknowns = append(unknowns, i)
		}
	}
	// fmt.Println(unknowns, len(unknowns))

	//find all combinations of . and # for each ? in arrangement
	combinations := math.Pow(2, float64(len(unknowns)))
	fmt.Println("Combinations:", combinations)

	// return int(combinations)

	count := 0
	//for each combination, replace ? with . and # and evaluate
	for i := 0; i < int(combinations); i++ {
		if i%1000000 == 0 && i != 0 {
			fmt.Println(i)
		}
		//convert i to binary
		binary := fmt.Sprintf("%b", i)
		// fmt.Println(binary)

		//pad binary with 0s to match length of unknowns
		for len(binary) < len(unknowns) {
			binary = "0" + binary
		}
		// fmt.Println(binary)

		for j, index := range unknowns {
			if binary[j] == '0' {
				arrangement = arrangement[:index] + "." + arrangement[index+1:]
			} else {
				arrangement = arrangement[:index] + "#" + arrangement[index+1:]
			}
		}
		// fmt.Println(arrangement)

		if isValid(arrangement, values) {
			count++
		}
	}

	return int(count)
}

func isValid(arrangement string, values []string) bool {
	split := strings.Split(arrangement, ".")

	contiguousValues := []int{}
	for _, str := range split {
		if str != "" {
			contiguousValues = append(contiguousValues, len(str))
		}
	}
	// fmt.Println(contiguousValues, values)

	if len(contiguousValues) != len(values) {
		return false
	}

	for i, val := range values {
		if val != fmt.Sprintf("%d", contiguousValues[i]) {
			return false
		}
	}

	return true
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
