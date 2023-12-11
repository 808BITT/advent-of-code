package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	cards := make(map[string]int)
	for _, line := range lines {
		cardString := strings.Split(line, ": ")[0]
		cardId := strings.Fields(cardString)[1]
		cards[cardId] = 1
	}
	// fmt.Println(cards)

	for _, line := range lines {
		split := strings.Split(line, ": ")
		cardId := strings.Fields(split[0])[1]

		nums := split[1]
		splitNums := strings.Split(nums, " | ")
		winning := strings.Fields(splitNums[0])
		possible := strings.Fields(splitNums[1])
		// fmt.Println("Card:", cardId, "Winning:", winning, "Possible:", possible)

		matches := 0
		for _, win := range winning {
			for _, pos := range possible {
				if win == pos {
					matches++
				}
			}
		}
		// fmt.Println("Matches:", matches)

		if matches > 0 {
			for i := 1; i < matches+1; i++ {
				winId, _ := strconv.Atoi(cardId)
				winId = winId + i
				// fmt.Println("Adding", cards[cardId], "to", winId)
				cards[strconv.Itoa(winId)] += cards[cardId]
			}

		}
	}

	sum := 0
	for _, v := range cards {
		sum += v
	}
	fmt.Println(sum)
	fmt.Printf("%v\n", time.Since(startTime))
}
