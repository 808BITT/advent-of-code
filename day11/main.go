package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1 Solution:", part1("sample.txt"))
	// fmt.Println("Part 2 Solution:", part2("sample.txt"))
}

func part1(filename string) int {
	grid := generateGrid(filename)
	fmt.Println(len(grid[0]), len(grid), grid)
	grid = expandGrid(grid)
	fmt.Println(len(grid[0]), len(grid), grid)

	for _, row := range grid {
		for _, pos := range row {
			fmt.Print(pos.t)
		}
		fmt.Println()
	}

	var galaxyLocations []Position
	galaxyNum := 0
	for i, row := range grid {
		for j, pos := range row {
			if pos.t == "#" {
				galaxyNum++
				galaxyLocations = append(galaxyLocations, Position{itoa(galaxyNum), j, i})
			}
		}
	}

	fmt.Println(galaxyLocations)

	//calculate the distance between each pair of galaxy locations and store in a map
	//find the sum of the distances for each location

	var distances []int
	for i := 0; i < len(galaxyLocations); i++ {
		for j := i + 1; j < len(galaxyLocations); j++ {
			d := distance(galaxyLocations[i], galaxyLocations[j])
			fmt.Println(galaxyLocations[i], galaxyLocations[j], d)
			distances = append(distances, d)
		}
	}

	fmt.Println(distances)
	sum := 0
	for _, d := range distances {
		sum += d
	}

	return sum
}

func distance(a, b Position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func generateGrid(filename string) Grid {
	lines := readFile(filename)
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = make([]Position, len(line))
		for j, char := range line {
			grid[i][j] = Position{string(char), j, i}
		}
	}
	return grid
}

func expandGrid(grid Grid) Grid {
	emptyRows := []int{}
	emptyCols := []int{}

	for i, row := range grid {
		allEmpty := true
		for _, pos := range row {
			if pos.t != "." {
				allEmpty = false
			}
		}
		if allEmpty {
			emptyRows = append(emptyRows, i)
		}
	}

	for i, _ := range grid[0] {
		allEmpty := true
		for _, pos := range grid {
			if pos[i].t != "." {
				allEmpty = false
			}
		}
		if allEmpty {
			emptyCols = append(emptyCols, i)
		}
	}

	fmt.Println(emptyRows, emptyCols)

	for i := 0; i < len(grid); i++ {
		// insert a "." into each row at the emptyCols indices
		inserted := 0
		for _, col := range emptyCols {
			grid[i] = append(grid[i][:col+inserted], append([]Position{{".", col + inserted, i}}, grid[i][col+inserted:]...)...)
			inserted++
		}
	}

	// insert a row of "." at the emptyRows indices
	emptyRow := make([]Position, len(grid[0]))
	for i := range grid[0] {
		emptyRow[i] = Position{".", i, 0}
	}

	inserted := 0
	for _, row := range emptyRows {
		grid = append(grid[:row+inserted], append([][]Position{emptyRow}, grid[row+inserted:]...)...)
		inserted++
	}

	new := make(Grid, len(grid))
	for i, row := range grid {
		new[i] = make([]Position, len(row))
		for j, pos := range row {
			new[i][j] = Position{pos.t, j, i}
		}
	}

	return new
}

func contains(list []int, item int) bool {
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
		// fmt.Println(scanner.Text())
	}
	return lines
}

type Grid [][]Position

type Position struct {
	t string
	x int
	y int
}
