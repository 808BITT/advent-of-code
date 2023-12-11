package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1:", part1())
	fmt.Println("Time: ", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2:", part2())
	fmt.Println("Time: ", time.Since(timer))
}

func part1() int {
	lines := readFile("input.txt")
	elves := []int{}
	sum := 0
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			sum += atoi(line)
			elves = append(elves, sum)
			sum = 0
		} else {
			sum += atoi(line)
		}
	}
	sort.Ints(elves)
	return elves[len(elves)-1]
}

func part2() int {
	lines := readFile("input.txt")
	elves := []int{}
	sum := 0
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			sum += atoi(line)
			elves = append(elves, sum)
			sum = 0
		} else {
			sum += atoi(line)
		}
	}
	sort.Ints(elves)
	return elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
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
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	return lines
}

func atoi(s string) int {
	i := 0
	for _, v := range s {
		i *= 10
		i += int(v - '0')
	}
	return i
}
