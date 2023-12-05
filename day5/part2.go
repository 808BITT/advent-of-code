package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

	seeds := strings.Split(lines[0], ": ")[1]
	seedLine := strings.Fields(seeds)
	seedRanges := make([]int, 0)
	for _, s := range seedLine {
		if s != "" {
			num, _ := strconv.Atoi(s)
			seedRanges = append(seedRanges, num)
		}
	}

	lowestList := make([]int, 0)

	var wg sync.WaitGroup
	for i := 0; i < len(seedRanges); i += 2 {
		wg.Add(1)
		go func(seedRanges []int, i int, lines []string) {
			defer wg.Done()
			lowest := newFunction(seedRanges, i, lines)
			lowestList = append(lowestList, lowest)
		}(seedRanges, i, lines)
	}
	wg.Wait()

	// Print the lowest location
	minLocation := 99999999999999
	for _, l := range lowestList {
		if l < minLocation {
			minLocation = l
		}
	}
}

func newFunction(seedRanges []int, i int, lines []string) int {
	lowest := 99999999999999
	// log.Println("Checking seed range:", seedRanges[i], "-", seedRanges[i+1]+seedRanges[i])
	counter := 0
	for seed := seedRanges[i]; seed <= seedRanges[i+1]+seedRanges[i]; seed++ {
		// log.Println("Checking seed:", seed)
		counter++
		if counter%10000 == 0 {
			log.Println("Checked", counter, "seeds")
		}
		index := 3

		allSoils := false
		seedToSoil := make(map[int][2]int)
		for !allSoils {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allSoils = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			seedToSoil[lookup] = [2]int{r, dest - lookup}
			index++
		}
		soil := seed
		for k, v := range seedToSoil {
			if seed >= k && seed <= k+v[0] {
				soil += v[1]
				break
			}
		}

		allFertilizers := false
		soilToFertilizer := make(map[int][2]int)
		for !allFertilizers {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allFertilizers = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			soilToFertilizer[lookup] = [2]int{r, dest - lookup}
			index++
		}
		fertilizer := soil
		for k, v := range soilToFertilizer {
			if soil >= k && soil <= k+v[0] {
				fertilizer += v[1]
				break
			}
		}

		allWaters := false
		fertilizerToWater := make(map[int][2]int)
		for !allWaters {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allWaters = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			fertilizerToWater[lookup] = [2]int{r, dest - lookup}
			index++
		}
		water := fertilizer
		for k, v := range fertilizerToWater {
			if fertilizer >= k && fertilizer <= k+v[0] {
				water += v[1]
				break
			}
		}

		allLights := false
		waterToLight := make(map[int][2]int)
		for !allLights {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allLights = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			waterToLight[lookup] = [2]int{r, dest - lookup}
			index++
		}
		light := water
		for k, v := range waterToLight {
			if water >= k && water <= k+v[0] {
				light += v[1]
				break
			}
		}

		allTemperatures := false
		lightToTemperature := make(map[int][2]int)
		for !allTemperatures {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allTemperatures = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			lightToTemperature[lookup] = [2]int{r, dest - lookup}
			index++
		}
		temperature := light
		for k, v := range lightToTemperature {
			if light >= k && light <= k+v[0] {
				temperature += v[1]
				break
			}
		}

		allHumidities := false
		temperatureToHumidity := make(map[int][2]int)
		for !allHumidities {
			line := strings.Fields(lines[index])
			if len(line) != 3 {
				allHumidities = true
				index += 2
				break
			}
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			temperatureToHumidity[lookup] = [2]int{r, dest - lookup}
			index++
		}
		humidity := temperature
		for k, v := range temperatureToHumidity {
			if temperature >= k && temperature <= k+v[0] {
				humidity += v[1]
				break
			}
		}

		humidityToLocation := make(map[int][2]int)
		for index < len(lines) {
			line := strings.Fields(lines[index])
			dest, _ := strconv.Atoi(line[0])
			lookup, _ := strconv.Atoi(line[1])
			r, _ := strconv.Atoi(line[2])
			humidityToLocation[lookup] = [2]int{r, dest - lookup}
			index++
		}
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
	return lowest
}
