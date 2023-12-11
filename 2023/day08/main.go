package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	input := readInput("input.txt")
	timer := getTime()
	fmt.Println("Part 1 Solution:", part1(input))
	fmt.Println("Part 1 Runtime:", getTime()-timer)
	timer = getTime()
	fmt.Println("Part 2 Solution:", part2V2(input))
	fmt.Println("Part 2 Runtime:", getTime()-timer)
}

func getTime() float64 {
	return float64(time.Now().UnixNano()) / 1000000000
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
	// fmt.Println(moveSequence)

	currentLocation, moveMap := generateMoveMapP1(input)
	// fmt.Println(moveMap)
	// fmt.Println(currentLocation)

	moves := 0
	for {
		// fmt.Println(currentLocation)
		// fmt.Println(moveMap[currentLocation].Left, moveMap[currentLocation].Right)
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

func part2(input []string) int {
	moveSequence := convertMoves(input[0])
	// fmt.Println(moveSequence)

	ghostLocations, endPoints, moveMap := generateMoveMapP2(input)
	fmt.Println(ghostLocations)
	fmt.Println(endPoints)

	endFound := false
	moves := 0
	for {
		// fmt.Println(ghostLocations)
		for i := 0; i < len(moveSequence); i++ {
			for j := 0; j < len(ghostLocations); j++ {
				if moveSequence[i] == 0 {
					ghostLocations[j] = moveMap[ghostLocations[j]].Left
				} else {
					ghostLocations[j] = moveMap[ghostLocations[j]].Right
				}
				moves++
			}
			// each ghost must be at an end point
			count := 0
			for j := 0; j < len(ghostLocations); j++ {
				for k := 0; k < len(endPoints); k++ {
					if ghostLocations[j] == endPoints[k] {
						count++
					}
				}
			}
			if count == len(ghostLocations) {
				endFound = true
				break
			}
		}
		if endFound {
			break
		}
	}
	return moves
}

func part2V2(input []string) int {
	moveSequence := convertMoves(input[0])

	ghostLocations, _, moveMap := generateMoveMapP2(input)

	movesPerGhost := make(MovesPerGhost)
	for i := 0; i < len(ghostLocations); i++ {
		movesPerGhost[ghostLocations[i]] = 0
	}

	// find how many moves it takes each ghost to get back to its end point
	for i := 0; i < len(ghostLocations); i++ {
		found := false
		moves := 0
		var visited []string
		currentLocation := ghostLocations[i]
		for {
			for j := 0; j < len(moveSequence); j++ {
				visited = append(visited, currentLocation)
				if moveSequence[j] == 0 {
					currentLocation = moveMap[currentLocation].Left
				} else {
					currentLocation = moveMap[currentLocation].Right
				}
				moves++
				if currentLocation[2] == 'Z' {
					found = true
					break
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		movesPerGhost[ghostLocations[i]] = moves
	}
	// find the intersection of the moves for each ghost using LCM
	return findIntersection(movesPerGhost)
}

func findIntersection(m MovesPerGhost) int {
	var moves []int
	for _, value := range m {
		moves = append(moves, value)
	}
	lcm := moves[0]
	for i := 1; i < len(m); i++ {
		lcm = findLCM(lcm, moves[i])
	}
	return lcm
}

func findLCM(a, b int) int {
	return a * b / findGCD(a, b)
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

func generateMoveMapP1(input []string) (string, map[string]Mapping) {
	// var startLocations []string
	moveMap := make(map[string]Mapping)
	for i := 2; i < len(input); i++ {
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`) // Updated regex pattern
		matches := re.FindStringSubmatch(input[i])
		// fmt.Println(matches[1], matches[2], matches[3])
		moveMap[matches[1]] = Mapping{matches[2], matches[3]}
		// if i == 2 {
		// 	startLocation = matches[1]
		// }
	}
	return "AAA", moveMap
}

func generateMoveMapP2(input []string) ([]string, []string, map[string]Mapping) {
	var startLocations []string
	var endLocations []string
	moveMap := make(map[string]Mapping)
	for i := 2; i < len(input); i++ {
		re := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`) // Updated regex pattern
		matches := re.FindStringSubmatch(input[i])
		// fmt.Println(matches[1], matches[2], matches[3])
		moveMap[matches[1]] = Mapping{matches[2], matches[3]}
		if matches[1][2] == 'A' {
			startLocations = append(startLocations, matches[1])
		}
		if matches[1][2] == 'Z' {
			endLocations = append(endLocations, matches[1])
		}
	}
	return startLocations, endLocations, moveMap
}

type Mapping struct {
	Left  string
	Right string
}

type MovesPerGhost map[string]int
