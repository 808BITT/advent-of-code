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
	var sum int
	for _, line := range lines {
		val := evaluatePart2(line)
		sum += val
		fmt.Scanln()
	}
	return sum
}

func evaluatePart2(line string) int {
	// Parse the line
	split := strings.Split(line, " ")
	original := split[0]
	groups := []int{}
	for _, val := range strings.Split(split[1], ",") {
		groups = append(groups, atoi(val))
	}

	out := 0

	valid := make(ArrangementMemo, len(groups))

	index := 0
	count := 0
	for {
		fmt.Println("index:", index, "count:", count)
		if index == len(original)-1 {
			break
		}
		if original[index] == '?' || original[index] == '#' {
			count++
		}
		if count == groups[0] && original[index+1] != '#' && original[index+1] != '?' {
			break
		}
		index++
	}

	valid[0] = findArrangements(original, groups[0], index+1)
	fmt.Println(valid[0])

	fmt.Println(original, groups, "->", out)
	return out
}

func findArrangements(arrangement string, groupSize int, startIndex int) ValidArrangements {
	fmt.Println("findArrangements", arrangement, groupSize, startIndex)

	subString := arrangement[:startIndex]
	fmt.Println(subString)

	indexesOfUnknowns := []int{}
	for i, char := range subString {
		if char == '?' {
			indexesOfUnknowns = append(indexesOfUnknowns, i)
		}
	}
	// fmt.Println(indexesOfUnknowns)

	var arrangements ValidArrangements
	for i := 0; i < int(math.Pow(2, float64(len(indexesOfUnknowns)))); i++ {
		binary := fmt.Sprintf("%b", i)
		for len(binary) < len(indexesOfUnknowns) {
			binary = "0" + binary
		}
		// fmt.Println(binary)

		arrangement := subString
		for j, index := range indexesOfUnknowns {
			if binary[j] == '0' {
				arrangement = arrangement[:index] + "." + arrangement[index+1:]
			} else {
				arrangement = arrangement[:index] + "#" + arrangement[index+1:]
			}
		}
		fmt.Println(arrangement)

		if isValid(arrangement, []string{fmt.Sprintf("%d", groupSize)}) {
			// trim trailing .'s from arrangement
			for arrangement[len(arrangement)-1] == '.' {
				arrangement = arrangement[:len(arrangement)-1]
			}
			arrangements = append(arrangements, arrangement+".")
		}
	}

	return arrangements
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

func atoi(s string) int {
	var num int
	for _, digit := range s {
		num = num*10 + int(digit-'0')
	}
	return num
}

type ArrangementMemo map[int]ValidArrangements

type ValidArrangements []string
