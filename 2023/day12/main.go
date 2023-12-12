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

func part2(filename string) int {
	lines := readFile(filename)
	sum := 0
	for _, line := range lines {
		val := evaluatePart1(line)
		additional := part2Expand(line)
		sum += val + additional
		fmt.Println(val + additional)
		fmt.Scanln()
	}
	return sum
}

func part2Expand(line string) int {
	modded := "?" + line
	newValue := evaluatePart1(modded)
	fmt.Println(modded, newValue)

	value := math.Pow(float64(newValue), 4)
	// fmt.Println(value)

	return int(value)
}

func part1(filename string) int {
	lines := readFile(filename)
	sum := 0
	for _, line := range lines {
		val := evaluatePart1(line)
		sum += val
		// fmt.Println(val)
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
	// fmt.Println(combinations)

	// return int(combinations)

	count := 0
	//for each combination, replace ? with . and # and evaluate
	for i := 0; i < int(combinations); i++ {
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
	// values corresponds to the number of contiguous #s in arrangement
	// check that the number of contiguous #s in arrangement matches the values
	// if so, return true
	// else, return false

	split := strings.Split(arrangement, ".")
	// fmt.Println(split)

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
