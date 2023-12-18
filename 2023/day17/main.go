package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	data, err := readInput("sample.txt")
	if err {
		return
	}

	timer := time.Now()
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Time taken:", time.Since(timer))
}

func readInput(filename string) ([]string, bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, true
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, true
	}
	return data, false
}

func part1(data []string) int {
	for _, line := range data {
		fmt.Println(line)
	}
	return 0
}
