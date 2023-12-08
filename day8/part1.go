package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	input := readInput("input.txt")
	fmt.Println("Part 1 Solution:", part1(input))
}

func readInput(filename string) []string {
	currDir, _ := os.Getwd()
	os.Chdir(currDir)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

func part1(input []string) int {
	moveSequence := convertMoves(input[0])
	fmt.Println(moveSequence)

	currentLocation, moveMap := generateMoveMap(input)
	fmt.Println(moveMap)
	fmt.Println(currentLocation)

	moves := 0
	for {
		fmt.Println(currentLocation)
		fmt.Println(moveMap[currentLocation].Left, moveMap[currentLocation].Right)
		fmt.Scanln()
		for i := 0; i < len(moveSequence); i++ {
			if moveSequence[i] == 0 {
				currentLocation = moveMap[currentLocation].Left
			} else {
				currentLocation = moveMap[currentLocation].Right
			}
			moves++
			if currentLocation == "ZZZ" {
				break
			}
		}
		if currentLocation == "ZZZ" {
			break
		}
	}
	return moves
}

func convertMoves(input string) []int {
	var moves []int
	for i := 0; i < len(input); i++ {
		if input[i] == 'L' {
			moves = append(moves, 0)
		} else if input[i] == 'R' {
			moves = append(moves, 1)
		}
	}
	return moves
}

func generateMoveMap(input []string) (string, map[string]Mapping) {
	var startLocation string
	moveMap := make(map[string]Mapping)
	for i := 2; i < len(input); i++ {
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`) // Updated regex pattern
		matches := re.FindStringSubmatch(input[i])
		fmt.Println(matches[1], matches[2], matches[3])
		moveMap[matches[1]] = Mapping{matches[2], matches[3]}
		if i == 2 {
			startLocation = matches[1]
		}
	}
	return startLocation, moveMap
}

type Mapping struct {
	Left  string
	Right string
}
