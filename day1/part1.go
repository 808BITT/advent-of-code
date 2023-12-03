package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	version1()
	version2()
}

func version1() {
	startTime := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sum int = 0
	var left, right int

	for _, line := range lines {
		for i, c := range line {
			if c >= '0' && c <= '9' {
				left = i
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				right = i
				break
			}
		}
		l, r := int(line[left]-'0'), int(line[right]-'0')
		new := l*10 + r
		sum += new
	}
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func version2() {
	startTime := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sum int = 0
	for _, line := range lines {
		var left, right int = 0, 0
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if left == 0 {
					left = i
					right = i
				} else {
					right = i
				}
			}
		}
		l, r := int(line[left]-'0'), int(line[right]-'0')
		sum += l*10 + r
	}
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}
