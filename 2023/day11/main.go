package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1 Solution:", part1("sample.txt"))
	fmt.Println("Time taken:", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2 Solution:", part2("input.txt"))
	fmt.Println("Time taken:", time.Since(timer))
}

func part2(filename string) int {
	grid := generateGrid(filename)
	// fmt.Println(len(grid[0]), len(grid), grid)
	r, c := findEmptyRowsAndCols(grid)
	// fmt.Println(len(grid[0]), len(grid), grid)

	// for i, _ := range grid[0] {
	// 	if contains(c, i) {
	// 		fmt.Print("V")
	// 	} else {
	// 		fmt.Print(" ")
	// 	}
	// }
	// fmt.Println()

	// for i, row := range grid {
	// 	for _, pos := range row {
	// 		fmt.Print(pos.t)
	// 	}
	// 	if contains(r, i) {
	// 		fmt.Print("<")
	// 	} else {
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()
	// }

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

	// fmt.Println(galaxyLocations)

	//calculate the distance between each pair of galaxy locations and store in a map
	//find the sum of the distances for each location

	var distances []int
	for i := 0; i < len(galaxyLocations); i++ {
		for j := i + 1; j < len(galaxyLocations); j++ {
			// d := distance(galaxyLocations[i], galaxyLocations[j])
			d := bigDistance(galaxyLocations[i], galaxyLocations[j], r, c, 1000000)

			// fmt.Println(galaxyLocations[i], galaxyLocations[j], d)
			distances = append(distances, d)
		}
	}

	// fmt.Println(distances)
	sum := 0
	for _, d := range distances {
		sum += d
	}

	return sum
}

func bigDistance(a, b Position, r, c []int, factor int) int {
	emptyRowsCrossed := 0
	emptyColsCrossed := 0

	for _, row := range r {
		if (row > a.y && row < b.y) || (row > b.y && row < a.y) {
			emptyRowsCrossed++
		}
	}

	for _, col := range c {
		if (col > a.x && col < b.x) || (col > b.x && col < a.x) {
			emptyColsCrossed++
		}
	}

	d := distance(a, b)

	// fmt.Println(a, b, emptyRowsCrossed, emptyColsCrossed, "d:", d, "factor:", factor)
	if emptyRowsCrossed == 0 && emptyColsCrossed == 0 {
		return d
	}

	baseDistance := distance(a, b) // Regular Manhattan distance
	return baseDistance + (emptyRowsCrossed+emptyColsCrossed)*(factor-1)

}

func part1(filename string) int {
	grid := generateGrid(filename)
	// fmt.Println(len(grid[0]), len(grid), grid)
	grid, _, _ = expandGrid(grid)
	// fmt.Println(len(grid[0]), len(grid), grid)

	// for _, row := range grid {
	// 	for _, pos := range row {
	// 		fmt.Print(pos.t)
	// 	}
	// 	fmt.Println()
	// }

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

	// fmt.Println(galaxyLocations)

	//calculate the distance between each pair of galaxy locations and store in a map
	//find the sum of the distances for each location

	var distances []int
	for i := 0; i < len(galaxyLocations); i++ {
		for j := i + 1; j < len(galaxyLocations); j++ {
			d := distance(galaxyLocations[i], galaxyLocations[j])
			// fmt.Println(galaxyLocations[i], galaxyLocations[j], d)
			distances = append(distances, d)
		}
	}

	// fmt.Println(distances)
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

func findEmptyRowsAndCols(grid Grid) ([]int, []int) {
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

	// fmt.Println(emptyRows, emptyCols)

	return emptyRows, emptyCols
}

func expandGrid(grid Grid) (Grid, []int, []int) {
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

	// fmt.Println(emptyRows, emptyCols)

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

	return new, emptyRows, emptyCols
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
