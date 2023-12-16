package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var directions = []string{"up", "right", "down", "left"}

func main() {
	data, err := readInput()
	if err {
		return
	}
	start := time.Now()
	fmt.Println("Part 1:", part1(data, directions))
	fmt.Println("Time taken:", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2:", part2(data, directions))
	fmt.Println("Time taken for Part 2:", time.Since(start))
}

func readInput() ([]string, bool) {
	file, err := os.Open("input.txt")
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

func reflect(direction, mirror string) string {
	switch mirror {
	case "/":
		switch direction {
		case "up":
			return "right"
		case "right":
			return "up"
		case "down":
			return "left"
		case "left":
			return "down"
		}
	case "\\":
		switch direction {
		case "up":
			return "left"
		case "left":
			return "up"
		case "down":
			return "right"
		case "right":
			return "down"
		}
	}
	return ""
}

func dirIndex(direction string) int {
	switch direction {
	case "up":
		return 0
	case "right":
		return 1
	case "down":
		return 2
	case "left":
		return 3
	}
	return -1
}

func trace(startX, startY int, startDirect string, data []string, directions []string, grid map[[3]int]int) {
	visited := [][3]int{{startX, startY, dirIndex(startDirect)}}
	for len(visited) > 0 {
		coord := visited[0]
		visited = visited[1:]
		x, y, direct := coord[0], coord[1], coord[2]
		if x < 0 || y < 0 || x >= len(data) || y >= len(data[0]) || grid[[3]int{x, y, direct}] > 0 {
			continue
		}
		current := string(data[x][y])
		switch current {
		case ".":
			grid[[3]int{x, y, direct}]++
		case "/", "\\":
			grid[[3]int{x, y, direct}]++
			newDirection := reflect(directions[direct], current)
			direct = dirIndex(newDirection)
		case "-":
			if direct == 1 || direct == 3 { // right or left
				grid[[3]int{x, y, direct}]++
			} else {
				grid[[3]int{x, y, direct}]++
				visited = append(visited, [3]int{x, y + 1, 1}, [3]int{x, y - 1, 3})
				continue
			}
		case "|":
			if direct == 0 || direct == 2 { // up or down
				grid[[3]int{x, y, direct}]++
			} else {
				grid[[3]int{x, y, direct}]++
				visited = append(visited, [3]int{x + 1, y, 2}, [3]int{x - 1, y, 0})
				continue
			}
		}
		switch direct {
		case 1: // right
			visited = append(visited, [3]int{x, y + 1, direct})
		case 3: // left
			visited = append(visited, [3]int{x, y - 1, direct})
		case 0: // up
			visited = append(visited, [3]int{x - 1, y, direct})
		case 2: // down
			visited = append(visited, [3]int{x + 1, y, direct})
		}
	}
}

func energize(x, y int, direct string, data []string, directions []string) int {
	grid := make(map[[3]int]int)
	for i := range data {
		for j := range data[i] {
			for k := range directions {
				grid[[3]int{i, j, k}] = 0
			}
		}
	}
	trace(x, y, direct, data, directions, grid)

	ctr := 0
	for i := range data {
		for j := range data[i] {
			if grid[[3]int{i, j, 0}] > 0 || grid[[3]int{i, j, 1}] > 0 || grid[[3]int{i, j, 2}] > 0 || grid[[3]int{i, j, 3}] > 0 {
				ctr++
			}
		}
	}
	return ctr
}

func part1(data []string, directions []string) int {
	return energize(0, 0, "right", data, directions)
}

func part2(data []string, directions []string) int {
	energies := []int{}
	for i := range data {
		energies = append(energies, energize(i, 0, "right", data, directions))
		energies = append(energies, energize(i, len(data[0])-1, "left", data, directions))
	}

	for i := range data[0] {
		energies = append(energies, energize(0, i, "down", data, directions))
		energies = append(energies, energize(len(data)-1, i, "up", data, directions))
	}

	maxEnergy := energies[0]
	for _, energy := range energies {
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	return maxEnergy
}
