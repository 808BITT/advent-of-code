package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1 Solution:", part1("input.txt"))
	fmt.Println("Time taken:", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2 Solution:", part2("input.txt"))
	fmt.Println("Time taken:", time.Since(timer))
}

func part2(filename string) int {
	// open input file
	lines := readFile(filename)
	// log.Println("Input Length:", len(lines))

	answer := 0

	// loop through lines
	for i := 0; i < len(lines); i++ {
		sequence := strings.Split(lines[i], " ")
		var numbers []int
		for _, s := range sequence {
			numbers = append(numbers, atoi(s))
		}
		// log.Println("Numbers:", numbers)

		vals := []int{numbers[0]} // first number is always the same

		var differences []int
		for j := 0; j < len(numbers)-1; j++ {
			differences = append(differences, numbers[j+1]-numbers[j])
		}
		// log.Println("Differences:", differences)
		allZeros := checkAllZeros(differences)

		if allZeros {
			answer += numbers[0]
			break
		}

		vals = append(vals, differences[0])

		for !allZeros {
			var buffer []int
			for j := 0; j < len(differences)-1; j++ {
				buffer = append(buffer, differences[j+1]-differences[j])
			}
			// log.Println("Differences:", buffer)
			vals = append(vals, buffer[0])
			differences = buffer
			allZeros = checkAllZeros(differences)
		}

		// log.Println("Vals:", vals)

		result := vals[len(vals)-1]
		for j := len(vals) - 1; j > 1; j-- {
			result = vals[j-1] - result
		}

		// log.Println("Result:", numbers[0]-result)
		answer += numbers[0] - result
	}

	return answer
}

func part1(filename string) int {
	// open input file
	lines := readFile(filename)
	// log.Println("Input Length:", len(lines))

	answer := 0

	// loop through lines
	for i := 0; i < len(lines); i++ {
		sequence := strings.Split(lines[i], " ")
		var numbers []int
		for _, s := range sequence {
			numbers = append(numbers, atoi(s))
		}
		// log.Println("Numbers:", numbers)

		var differences []int
		for j := 0; j < len(numbers)-1; j++ {
			differences = append(differences, numbers[j+1]-numbers[j])
		}
		allZeros := checkAllZeros(differences)

		if allZeros {
			answer += numbers[len(numbers)-1]
			break
		}

		intermediate := numbers[len(numbers)-1] + differences[len(differences)-1]
		// fmt.Println("Intermediate:", intermediate)
		for !allZeros {
			var buffer []int
			for j := 0; j < len(differences)-1; j++ {
				buffer = append(buffer, differences[j+1]-differences[j])
			}
			// log.Println("Differences:", buffer)
			// log.Println("Last Difference:", buffer[len(buffer)-1])
			intermediate += buffer[len(buffer)-1]
			// fmt.Println("Intermediate:", intermediate)
			differences = buffer
			allZeros = checkAllZeros(differences)
		}

		// fmt.Println("Intermediate:", intermediate)
		answer += intermediate
	}

	return answer
}

func checkAllZeros(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}

func readFile(filename string) []string {
	// open input file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read input file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// return lines
	return lines
}

func atoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
