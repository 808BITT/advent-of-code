package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2:", part2("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
}

func part2(filename string) int {
	lines := readFile(filename)

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

		found := false
		for j := 0; j < len(first); j++ {
			for k := 0; k < len(second); k++ {
				for l := 0; l < len(third); l++ {
					if first[j] == second[k] && second[k] == third[l] {
						sum += convertToInt(string(first[j]))
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
	}
	return sum
}

func part1(filename string) int {
	lines := readFile(filename)
	sum := 0
	for _, line := range lines {
		subWidth := len(line) / 2
		leftChars := line[:subWidth]
		rightChars := line[subWidth:]

		var match string
		for i := 0; i < subWidth; i++ {
			for j := 0; j < subWidth; j++ {
				if leftChars[i] == rightChars[j] {
					match = string(leftChars[i])
					break
				}
			}
			if match != "" {
				break
			}
		}
		intMatch := convertToInt(match)
		sum += intMatch
		// fmt.Println(intMatch)
	}
	return sum
}

func readFile(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func convertToInt(s string) int {
	char := s[0]
	if char >= 'a' && char <= 'z' {
		// a-z -> 1-26
		return int(char) - 96
	} else if char >= 'A' && char <= 'Z' {
		// A-Z -> 27-52
		return int(char) - 38
	}
	// Handle non-alphabetic characters or return a default value
	return 0
}
