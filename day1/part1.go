package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// set lines to the sample input for testing
	// var lines = []string{
	// 	"pqr3stu8vwx",
	// 	"1234",
	// 	"91212129",
	// 	"1111",
	// 	"a1b2c3d4e5f",
	// }

	// left and right most digits are combined to a new number
	// pqr3stu8vwx -> 38
	// 1234 -> 14
	// 91212129 -> 9
	// 1111 -> 1
	// a1b2c3d4e5f = 15

	var sum int = 0
	var left, right int

	for _, line := range lines {
		// find left most digit
		for i, c := range line {
			if c >= '0' && c <= '9' {
				left = i
				break
			}
		}
		// find right most digit
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				right = i
				break
			}
		}
		// convert to int
		l, r := int(line[left]-'0'), int(line[right]-'0')
		new := l*10 + r
		sum += new

		// debug
		// fmt.Printf("%s -> %d\n", line, l*10+r)
	}

	fmt.Printf("Sum: %d\n", sum)
}
