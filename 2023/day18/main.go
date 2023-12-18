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
	data, err := readInput("sample.txt")
	if err {
		return
	}

	timer := time.Now()
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Time taken:", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2:", part2(data))
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
	// calculate min/max x, y
	curX, curY := 0, 0
	var minX, maxX, minY, maxY int
	for _, line := range data {
		split := strings.Split(line, " ")
		dir := split[0]
		dist, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return 0
		}

		switch dir {
		case "R":
			curX += dist
			if curX > maxX {
				maxX = curX
			}
		case "L":
			curX -= dist
			if curX < minX {
				minX = curX
			}
		case "U":
			curY -= dist
			if curY < minY {
				minY = curY
			}
		case "D":
			curY += dist
			if curY > maxY {
				maxY = curY
			}
		default:
			fmt.Println("Unknown direction:", dir)
			return 0
		}
	}

	curX, curY = 0, 0

	if minX < 0 {
		maxX += minX*-1 + 1
		curX += minX*-1 + 1
		minX = 0
	}
	if minY < 0 {
		maxY += minY*-1 + 1
		curY += minY*-1 + 1
		minY = 0
	}

	// create grid
	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	// fmt.Println(minX, maxX, minY, maxY)

	// draw path
	for _, line := range data {
		split := strings.Split(line, " ")
		dir := split[0]
		dist, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return 0
		}

		// fmt.Println(dir, dist, curX, curY)

		switch dir {
		case "R":
			for i := 0; i < dist; i++ {
				curX++
				grid[curY][curX] = 1
			}
		case "L":
			for i := 0; i < dist; i++ {
				curX--
				grid[curY][curX] = 1
			}
		case "U":
			for i := 0; i < dist; i++ {
				curY--
				grid[curY+1][curX] = -1
				grid[curY][curX] = -1
			}
		case "D":
			for i := 0; i < dist; i++ {
				curY++
				grid[curY-1][curX] = 2
				grid[curY][curX] = 2
			}
		default:
			fmt.Println("Unknown direction:", dir)
			return 0
		}
	}
	// show(grid)
	// write grid to file named "grid.txt"
	file, err := os.Create("grid.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 0
	}
	for _, line := range grid {
		for _, val := range line {
			if val == 0 {
				file.WriteString(".")
			} else if val == -1 {
				file.WriteString("^")
			} else if val == 2 {
				file.WriteString("v")
			} else {
				file.WriteString("#")
			}
		}
		file.WriteString("\n")
	}
	file.Close()

	// fill in the "pool"
	for y := 0; y < len(grid)-1; y++ {
		inside := false
		for x := 0; x < len(grid[y])-1; x++ {
			if grid[y][x] == -1 {
				inside = true
			}
			if grid[y][x] == 2 {
				inside = false
				continue
			}
			if inside {
				grid[y][x] = 1
			}
		}
	}

	file, err = os.Create("filled.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 0
	}
	for _, line := range grid {
		for _, val := range line {
			if val == 0 {
				file.WriteString(".")
			} else {
				file.WriteString("#")
			}
		}
		file.WriteString("\n")
	}
	file.Close()

	// show(grid)

	// calculate sum of grid
	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 0 {
				sum++
			}
		}
	}

	// return
	return sum
}

func part2(data []string) int {
	sum := 0
	return sum
}

func atoh(s string) (int, error) {
	var res int
	for _, c := range s {
		res *= 16
		switch {
		case '0' <= c && c <= '9':
			res += int(c - '0')
		case 'a' <= c && c <= 'f':
			res += int(c - 'a' + 10)
		default:
			return 0, fmt.Errorf("invalid hex digit: %q", c)
		}
	}
	return res, nil
}

func show(grid [][]int) {
	for _, line := range grid {
		for _, val := range line {
			if val == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
