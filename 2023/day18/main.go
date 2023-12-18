package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	curX, curY, minX, maxX, minY, maxY := readData(data)

	CHUNK_SIZE := 100000

	horizontalChunks := (maxX-minX)/CHUNK_SIZE + 1
	verticalChunks := (maxY-minY)/CHUNK_SIZE + 1

	grid := make([][]Chunk, verticalChunks)
	for i := range grid {
		grid[i] = make([]Chunk, horizontalChunks)
	}

	fmt.Println("Starting Chunk: ", curY/CHUNK_SIZE, curX/CHUNK_SIZE, "(", curX%CHUNK_SIZE, ",", curY%CHUNK_SIZE, ")")
	grid[curY/CHUNK_SIZE][curX/CHUNK_SIZE].Visited = true
	grid[curY/CHUNK_SIZE][curX/CHUNK_SIZE].Entry = Coord{curX % CHUNK_SIZE, curY % CHUNK_SIZE}
	fmt.Println("Entry:", grid[curY/CHUNK_SIZE][curX/CHUNK_SIZE].Entry)
	fmt.Println("------------------")

	sum := 0
	for _, line := range data {
		fmt.Println("Current Y, X:", curY, curX)
		grid[curY/CHUNK_SIZE][curX/CHUNK_SIZE].Visited = true

		rawHex := regexp.MustCompile(`\(#([0-9a-f]{6})\)`).FindStringSubmatch(line)[1]
		distance, _ := atoh(rawHex[0:5])
		direction := Direction(rawHex[5])

		switch direction {
		case Right:
			fmt.Println("Right", distance)
			curChunkX := curX / CHUNK_SIZE
			curChunkY := curY / CHUNK_SIZE
			if curX+distance > (curChunkX+1)*CHUNK_SIZE {
				fmt.Println("Exiting chunk to the right")
				grid[curChunkY][curChunkX].Exit = Coord{CHUNK_SIZE - 1, curY % CHUNK_SIZE}

				for i := curChunkX + 1; i < (curX+distance)/CHUNK_SIZE; i++ {
					grid[curChunkY][i].Visited = true
					grid[curChunkY][i].Entry = Coord{0, curY % CHUNK_SIZE}
					grid[curChunkY][i].Exit = Coord{CHUNK_SIZE - 1, curY % CHUNK_SIZE}
					fmt.Println("Passed through chunk", curChunkY, i)
				}

				curX += distance
				newChunkX := curX / CHUNK_SIZE
				grid[curChunkY][newChunkX].Visited = true
				grid[curChunkY][newChunkX].Entry = Coord{0, curY % CHUNK_SIZE}
				grid[curChunkY][newChunkX].Midpoints = append(grid[curChunkY][newChunkX].Midpoints, Coord{curX % CHUNK_SIZE, curY % CHUNK_SIZE})
				fmt.Println("Entry:", grid[curChunkY][newChunkX].Entry)
				fmt.Println("Midpoint:", grid[curChunkY][newChunkX].Midpoints)
			} else {
				curX += distance
				grid[curChunkY][curChunkX].Midpoints = append(grid[curChunkY][curChunkX].Midpoints, Coord{curX % CHUNK_SIZE, curY % CHUNK_SIZE})
				fmt.Println("Stayed in chunk", curChunkY, curChunkX)
				fmt.Println("Midpoint:", grid[curChunkY][curChunkX].Midpoints)
			}

		case Left:
			fmt.Println("Left", distance)
		case Up:
			fmt.Println("Up", distance)
		case Down:
			fmt.Println("Down", distance)
		}
		fmt.Println()
	}

	return sum
}

func readData(data []string) (int, int, int, int, int, int) {
	var curX, curY, minX, maxX, minY, maxY int
	for _, line := range data {
		rawHex := regexp.MustCompile(`\(#([0-9a-f]{6})\)`).FindStringSubmatch(line)[1]
		distance, _ := atoh(rawHex[0:5])
		direction := rawHex[5:6]

		switch direction {
		case "0":
			curX += distance
			if curX > maxX {
				maxX = curX
			}
		case "1":
			curY += distance
			if curY > maxY {
				maxY = curY
			}
		case "2":
			curX -= distance
			if curX < minX {
				minX = curX
			}
		case "3":
			curY -= distance
			if curY < minY {
				minY = curY
			}
		default:
			fmt.Println("Unknown direction:", direction)
			return 0, 0, 0, 0, 0, 0
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

	return curX, curY, minX, maxX, minY, maxY
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

type Chunk struct {
	Visited   bool
	Entry     Coord
	Exit      Coord
	Midpoints []Coord
	Direction Direction
}

type Coord struct {
	X, Y int
}

type Direction string

const (
	Right Direction = "0"
	Down  Direction = "1"
	Left  Direction = "2"
	Up    Direction = "3"
)
