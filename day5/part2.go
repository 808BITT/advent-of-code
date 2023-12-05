package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	currDir, _ := os.Getwd()
	os.Chdir(currDir)

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	lowest := 99999999999999
	seedLine := strings.Split(strings.TrimSuffix(strings.Split(lines[0], ": ")[1], "\n"), " ")

	seedRanges := make([]int, 0)
	for _, s := range seedLine {
		if s != "" {
			num, _ := strconv.Atoi(s)
			seedRanges = append(seedRanges, num)
		}
	}

	for i := 0; i < len(seedRanges); i += 2 {
		for seed := seedRanges[i]; seed <= seedRanges[i+1]+seedRanges[i]; seed++ {
			index := 3

			// Read in soil map
			allSoils := false
			seedToSoil := make(map[int][2]int)
			for !allSoils {
				if lines[index] == "" {
					allSoils = true
					index += 2
					break
				}

				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				seedToSoil[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply soil map to seed
			soil := seed
			for k, v := range seedToSoil {
				if seed >= k && seed <= k+v[0] {
					soil += v[1]
					break
				}
			}

			// Read in fertilizer map
			allFertilizers := false
			soilToFertilizer := make(map[int][2]int)
			for !allFertilizers {
				if lines[index] == "" {
					allFertilizers = true
					index += 2
					break
				}
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				soilToFertilizer[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply fertilizer map to soil
			fertilizer := soil
			for k, v := range soilToFertilizer {
				if soil >= k && soil <= k+v[0] {
					fertilizer += v[1]
					break
				}
			}

			// Read in water map
			allWaters := false
			fertilizerToWater := make(map[int][2]int)
			for !allWaters {
				if lines[index] == "" {
					allWaters = true
					index += 2
					break
				}
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				fertilizerToWater[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply water map to fertilizers
			water := fertilizer
			for k, v := range fertilizerToWater {
				if fertilizer >= k && fertilizer <= k+v[0] {
					water += v[1]
					break
				}
			}

			// Read in light map
			allLights := false
			waterToLight := make(map[int][2]int)
			for !allLights {
				if lines[index] == "" {
					allLights = true
					index += 2
					break
				}
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				waterToLight[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply light map to waters
			light := water
			for k, v := range waterToLight {
				if water >= k && water <= k+v[0] {
					light += v[1]
					break
				}
			}

			// Read in temperature map
			allTemperatures := false
			lightToTemperature := make(map[int][2]int)
			for !allTemperatures {
				if lines[index] == "" {
					allTemperatures = true
					index += 2
					break
				}
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				lightToTemperature[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply temperature map to lights
			temperature := light
			for k, v := range lightToTemperature {
				if light >= k && light <= k+v[0] {
					temperature += v[1]
					break
				}
			}

			// Read in humidity map
			allHumidities := false
			temperatureToHumidity := make(map[int][2]int)
			for !allHumidities {
				if lines[index] == "" {
					allHumidities = true
					index += 2
					break
				}
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				temperatureToHumidity[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply humidity map to temperatures
			humidity := temperature
			for k, v := range temperatureToHumidity {
				if temperature >= k && temperature <= k+v[0] {
					humidity += v[1]
					break
				}
			}

			// Read in location map
			humidityToLocation := make(map[int][2]int)
			for index < len(lines) {
				line := strings.Split(strings.TrimSuffix(lines[index], "\n"), " ")
				dest, _ := strconv.Atoi(line[0])
				lookup, _ := strconv.Atoi(line[1])
				r, _ := strconv.Atoi(line[2])
				humidityToLocation[lookup] = [2]int{r, dest - lookup}
				index++
			}

			// Apply location map to humidities
			location := humidity
			for k, v := range humidityToLocation {
				if humidity >= k && humidity <= k+v[0] {
					location += v[1]
					break
				}
			}

			if location < lowest {
				lowest = location
			}
		}
	}

	// Print the lowest location
	fmt.Println("Lowest Location:", lowest)
}
