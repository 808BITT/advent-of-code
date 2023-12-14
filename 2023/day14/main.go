package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const CYCLES = 1000000000

func main() {
	timer := time.Now()
	fmt.Println(part2("input.txt"))
	fmt.Println(time.Since(timer))
}

func showMap(rockMap [][]rune) {
	for _, line := range rockMap {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func part2(file string) int {
	memo := make(map[int]string)
	rockMap := parseFile(file)
	for i := range rockMap {
		rockMap[i] = []rune(rockMap[i])
	}

	start := 0
	end := 0

	for i := 0; i < CYCLES; i++ {
		fmt.Printf("Cycle: %d of %d Current Weight: %d\n", i, CYCLES, sumWeight(rockMap))
		memo[i] = toRockString(rockMap)

		// showMap(rockMap)
		// fmt.Scanln()

		// NORTH TILT
		index := 0
		for {
			if index == len(rockMap)-1 {
				break
			}

			rolling := false
			delta := false
			for j := 0; j < len(rockMap)-1; j++ { // walk each row
				rockMap, delta = tiltNorth(j, rockMap)
				if delta {
					rolling = true
				}
			}

			if !rolling {
				break
			}

			index++
		}
		// fmt.Println("NORTH TILT")
		// showMap(rockMap)
		// fmt.Scanln()

		// WEST TILT
		index = 0
		for {
			if index == len(rockMap[0]) {
				break
			}

			rolling := false
			delta := false
			for j := 0; j < len(rockMap[0])-2; j++ { // walk each column
				rockMap, delta = tiltWest(j, rockMap)
				if delta {
					rolling = true
				}
			}

			if !rolling {
				break
			}

			index++
		}

		// showMap(rockMap)
		// fmt.Scanln()

		// SOUTH TILT
		index = 0
		for {
			if index == len(rockMap)-1 {
				break
			}

			rolling := false
			delta := false
			for j := 0; j < len(rockMap)-1; j++ { // walk each row
				rockMap, delta = tiltSouth(j, rockMap)
				if delta {
					rolling = true
				}
			}

			if !rolling {
				break
			}

			index++
		}

		// showMap(rockMap)
		// fmt.Scanln()

		// EAST TILT
		index = 0
		for {
			if index == len(rockMap[0]) {
				break
			}

			rolling := false
			delta := false
			for j := 0; j < len(rockMap[0])-2; j++ { // walk each column
				rockMap, delta = tiltEast(j, rockMap)
				if delta {
					rolling = true
				}
			}

			if !rolling {
				break
			}

			index++
		}

		// showMap(rockMap)
		// fmt.Scanln()

		memoed := false
		for j := 0; j <= i+1; j++ {
			if memo[j] == toRockString(rockMap) {
				memoed = true
				break
			}
		}

		if memoed {
			start = 0
			for j := 0; j <= i+1; j++ {
				if memo[j] == toRockString(rockMap) {
					start = j
					break
				}
			}
			end = i + 1
			remaining := CYCLES - end
			cycleLength := end - start
			offset := remaining % cycleLength

			cache := fromRockString(memo[start+offset])
			for i := range cache {
				cache[i] = []rune(cache[i])
			}

			return sumWeight(cache)
		}
	}

	return sumWeight(rockMap)
}

func tiltNorth(index int, rockMap [][]rune) ([][]rune, bool) {
	delta := false
	for i := 0; i < len(rockMap[0]); i++ { // walk each column
		if rockMap[index][i] == '.' && rockMap[index+1][i] == 'O' {
			rockMap[index][i] = 'O'
			rockMap[index+1][i] = '.'
			delta = true
		}
	}

	return rockMap, delta
}

func tiltWest(index int, rockMap [][]rune) ([][]rune, bool) {
	delta := false
	for i := 0; i < len(rockMap); i++ { // walk each row
		if rockMap[i][index+1] == 'O' && rockMap[i][index] == '.' {
			rockMap[i][index+1] = '.'
			rockMap[i][index] = 'O'
			delta = true
		}
	}

	return rockMap, delta
}

func tiltSouth(index int, rockMap [][]rune) ([][]rune, bool) {
	delta := false
	for i := 0; i < len(rockMap[0]); i++ { // walk each column
		if rockMap[index][i] == 'O' && rockMap[index+1][i] == '.' {
			rockMap[index][i] = '.'
			rockMap[index+1][i] = 'O'
			delta = true
		}
	}

	return rockMap, delta
}

func tiltEast(index int, rockMap [][]rune) ([][]rune, bool) {
	delta := false
	for i := 0; i < len(rockMap); i++ { // walk each row
		if rockMap[i][index] == 'O' && rockMap[i][index+1] == '.' {
			rockMap[i][index] = '.'
			rockMap[i][index+1] = 'O'
			delta = true
		}
	}

	return rockMap, delta
}

func parseFile(file string) [][]rune {
	data, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(data), "\n")
	rockMap := make([][]rune, len(lines))
	for i, line := range lines {
		rockMap[i] = []rune(line)
	}
	return rockMap
}

func sumWeight(rockMap [][]rune) int {
	sum := 0
	for i, line := range rockMap {
		rocks := countRocks(line)
		sum += rocks * (len(rockMap) - i)
	}

	return sum
}

func countRocks(line []rune) int {
	sum := 0
	for _, char := range line {
		if char == 'O' {
			sum++
		}
	}
	return sum
}

func toRockString(rockMap [][]rune) string {
	var str strings.Builder
	for i, line := range rockMap {
		for _, char := range line {
			str.WriteRune(char)
		}
		if i != len(rockMap)-1 {
			str.WriteRune('\n')
		}
	}
	return str.String()
}

func fromRockString(str string) [][]rune {
	lines := strings.Split(str, "\n")
	rockMap := make([][]rune, len(lines))
	for i, line := range lines {
		rockMap[i] = []rune(line)
	}
	return rockMap
}
