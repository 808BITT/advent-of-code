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
	fmt.Println("Part 1 Solution:", part1("sample.txt"))
	fmt.Println("Time taken:", time.Since(timer))
	// timer = time.Now()
	// fmt.Println("Part 2 Solution:", part2("input.txt"))
	// fmt.Println("Time taken:", time.Since(timer))
}

type Position struct {
	x int
	y int
}

type Direction struct {
	name string
	dx   int
	dy   int
}

type Pipe struct {
	connections []Direction
}

type PipeMapping map[string]Pipe

func part1(filename string) int {
	north := Direction{"North", 0, 1}
	south := Direction{"South", 0, -1}
	east := Direction{"East", 1, 0}
	west := Direction{"West", -1, 0}

	lookup := make(PipeMapping)
	lookup["|"] = Pipe{[]Direction{north, south}}
	lookup["-"] = Pipe{[]Direction{east, west}}
	lookup["L"] = Pipe{[]Direction{north, east}}
	lookup["J"] = Pipe{[]Direction{north, west}}
	lookup["7"] = Pipe{[]Direction{south, west}}
	lookup["F"] = Pipe{[]Direction{south, east}}
	lookup["S"] = Pipe{[]Direction{north, east, south, west}}

	lines := readFile(filename)
	pipeMap, startX, startY := generateMap(lines)
	// fmt.Println(pipeMap)
	// fmt.Println(startX, startY)

	// start at S
	// check all directions
	// if there is a pipe, move to that pipe and add 1 to the counter
	// if we get back to S, stop

	currentX := startX
	currentY := startY

	backToStart := false
	visited := []Position{{currentX, currentY}}
	steps := 0
	for !backToStart {
		fmt.Scanln()
		current := pipeMap[currentY][currentX]
		fmt.Println("Current Position:", current, currentX, currentY)
		for _, direction := range lookup[current].connections {
			fmt.Println("Checking direction:", direction)

			new := pipeMap[currentY+direction.dy][currentX+direction.dx]
			fmt.Println("-- Checking Position:", new, currentX+direction.dx, currentY+direction.dy)

			pipeFound := false
			for p, _ := range lookup {
				if new == p && !contains(visited, Position{currentX + direction.dx, currentY + direction.dy}) {
					fmt.Println("-- Found new pipe:", Position{currentX + direction.dx, currentY + direction.dy})
					visited = append(visited, Position{currentX + direction.dx, currentY + direction.dy})
					pipeFound = true
					currentX += direction.dx
					currentY += direction.dy
					break
				} else if new == p {
					fmt.Println("-- Already visited pipe:", Position{currentX + direction.dx, currentY + direction.dy})
				}
			}
			if pipeFound {
				fmt.Println("Moving to new pipe:", Position{currentX, currentY})
				steps++
				break
			}
		}
	}
	return steps / 2
}

func contains(list []Position, item Position) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func generateMap(lines []string) (Grid, int, int) {
	var grid Grid
	var x, y int
	for l, line := range lines {
		var row []string
		for c, char := range line {
			row = append(row, string(char))
			if string(char) == "S" {
				x = c
				y = l
			}
		}
		grid = append(grid, row)
	}
	return grid, x, y
}

// need a 2D char array to store the grid
type Grid [][]string
