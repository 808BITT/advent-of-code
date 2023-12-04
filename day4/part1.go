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
	startTime := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// var lines []string = []string{
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	// }

	sum := 0
	for _, line := range lines {
		matches := 0
		split := strings.Split(line, ": ")
		nums := split[1]
		splitNums := strings.Split(nums, " | ")
		winning := strings.Fields(splitNums[0])
		possible := strings.Fields(splitNums[1])

		for _, win := range winning {
			for _, pos := range possible {
				if win == pos {
					matches++
				}
			}
		}

		if matches > 0 {
			points := int(math.Pow(2, float64(matches-1)))
			// fmt.Println(card, matches, points)
			sum += points
		}
	}

	fmt.Println(sum)
	fmt.Printf("%v\n", time.Since(startTime))
}
