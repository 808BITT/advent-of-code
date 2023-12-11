package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
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

	v2(lines, seedRanges) // Brute force locations
	// v1(lines, seedRanges) // Brute force seeds .... LOL
}

func v2(lines []string, seedRanges []int) {
	lowest := 999999999
	index := lowest
	for {
		index -= 100000
		if index < 0 {
			index = lowest + 1
			break
		}
		valid := checkIfLocationIsValid(index, lines)
		if valid {
			lowest = index
		}
		fmt.Println("Lowest:", lowest, "Index:", index)
	}
	multiplier := 1
	streak := 0
	for {
		index -= multiplier
		if index < 0 {
			break
		}
		valid := checkIfLocationIsValid(index, lines)
		if valid {
			lowest = index
			streak = 0
			multiplier = 1
		} else {
			streak++
			if streak > 1000 && multiplier < 50000 {
				multiplier += 1
			}
		}
		fmt.Println("Lowest:", lowest, "Index:", index)
	}
}

func checkIfLocationIsValid(location int, lines []string) bool {
	index := len(lines) - 1

	allLocations := false
	locationToHumidity := make(map[int][2]int)
	for !allLocations {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allLocations = true
			index -= 2
			break
		}
		l, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		locationToHumidity[l] = [2]int{r, d - l}
		index--
	}
	humidity := location
	for k, v := range locationToHumidity {
		if humidity >= k && humidity <= k+v[0] {
			humidity += v[1]
			break
		}
	}
	// fmt.Println("-- Humidity:", humidity)

	allHumidities := false
	humidityToTemperature := make(map[int][2]int)
	for !allHumidities {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allHumidities = true
			index -= 2
			break
		}
		h, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		humidityToTemperature[h] = [2]int{r, d - h}
		index--
	}
	temperature := humidity
	for k, v := range humidityToTemperature {
		if humidity >= k && humidity <= k+v[0] {
			temperature += v[1]
			break
		}
	}
	// fmt.Println("-- Temperature:", temperature)

	allTemperatures := false
	temperatureToLight := make(map[int][2]int)
	for !allTemperatures {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allTemperatures = true
			index -= 2
			break
		}
		t, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		temperatureToLight[t] = [2]int{r, d - t}
		index--
	}
	light := temperature
	for k, v := range temperatureToLight {
		if temperature >= k && temperature <= k+v[0] {
			light += v[1]
			break
		}
	}
	// fmt.Println("-- Light:", light)

	allLights := false
	lightToWater := make(map[int][2]int)
	for !allLights {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allLights = true
			index -= 2
			break
		}
		l, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		lightToWater[l] = [2]int{r, d - l}
		index--
	}
	water := light
	for k, v := range lightToWater {
		if light >= k && light <= k+v[0] {
			water += v[1]
			break
		}
	}
	// fmt.Println("-- Water:", water)

	allWaters := false
	waterToFertilizer := make(map[int][2]int)
	for !allWaters {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allWaters = true
			index -= 2
			break
		}
		w, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		waterToFertilizer[w] = [2]int{r, d - w}
		index--
	}
	fertilizer := water
	for k, v := range waterToFertilizer {
		if water >= k && water <= k+v[0] {
			fertilizer += v[1]
			break
		}
	}
	// fmt.Println("-- Fertilizer:", fertilizer)

	allFertilizers := false
	fertilizerToSoil := make(map[int][2]int)
	for !allFertilizers {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allFertilizers = true
			index -= 2
			break
		}
		f, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		fertilizerToSoil[f] = [2]int{r, d - f}
		index--
	}
	soil := fertilizer
	for k, v := range fertilizerToSoil {
		if fertilizer >= k && fertilizer <= k+v[0] {
			soil += v[1]
			break
		}
	}
	// fmt.Println("-- Soil:", soil)

	allSoils := false
	soilToSeed := make(map[int][2]int)
	for !allSoils {
		line := strings.Fields(lines[index])
		if len(line) != 3 {
			allSoils = true
			index -= 2
			break
		}
		s, _ := strconv.Atoi(line[0])
		d, _ := strconv.Atoi(line[1])
		r, _ := strconv.Atoi(line[2])
		soilToSeed[s] = [2]int{r, d - s}
		index--
	}
	seed := soil
	for k, v := range soilToSeed {
		if soil >= k && soil <= k+v[0] {
			seed += v[1]
			break
		}
	}
	// fmt.Println("-- Seed:", seed)

	// check if seed is in one of the seed ranges
	seeds := strings.Split(lines[0], ": ")[1]
	seedLine := strings.Fields(seeds)
	seedRanges := make([]int, 0)
	for _, s := range seedLine {
		if s != "" {
			num, _ := strconv.Atoi(s)
			seedRanges = append(seedRanges, num)
		}
	}

	for i := 0; i < len(seedRanges); i += 2 {
		if seed >= seedRanges[i] && seed <= seedRanges[i+1]+seedRanges[i] {
			// fmt.Println("-- Seed is in range:", seedRanges[i], "-", seedRanges[i+1]+seedRanges[i])
			return true
		}
	}

	// fmt.Println("-- Seed is not in range")
	return false
}

func v1(lines []string, seedRanges []int) {
	splits := 10000000

	splitRanges := make([]int, 0)
	for i := 0; i < len(seedRanges); i += 2 {
		inc := seedRanges[i+1] / splits
		rem := seedRanges[i+1] % splits
		splitRanges = append(splitRanges, seedRanges[i])
		splitRanges = append(splitRanges, inc)
		for j := 0; j < splits-1; j++ {
			splitRanges = append(splitRanges, seedRanges[i]+inc*(j+1))
			splitRanges = append(splitRanges, inc)
		}
		splitRanges = append(splitRanges, seedRanges[i]+inc*splits)
		splitRanges = append(splitRanges, rem)
	}

	// print the number of seeds in each range
	// for i := 0; i < len(splitRanges); i += 2 {
	// 	fmt.Println("Range:", splitRanges[i], "-", splitRanges[i+1]+splitRanges[i], "has", splitRanges[i+1], "seeds")
	// }

	// return

	// lowestList := make([]int, 0)
	var counter int64 = 0
	var globalLowest int64 = 1500948298
	var wg sync.WaitGroup
	fmt.Println("Waiting for workers to finish...")

	// periodically print the number of workers still running
	// go func() {
	// 	for {
	// 		fmt.Println("Workers remaining:", atomic.LoadInt64(&counter))
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	for i := 0; i < len(splitRanges); i += 2 {
		wg.Add(1)
		go func(seedRanges []int, i int, lines []string) {
			atomic.AddInt64(&counter, 1)
			defer wg.Done()
			findLowest(splitRanges, i, lines, &globalLowest)
			atomic.AddInt64(&counter, -1)
		}(splitRanges, i, lines)
	}

	wg.Wait()
	time.Sleep(1 * time.Second)

	fmt.Println("Lowest location:", globalLowest)
}

func findLowest(seedRanges []int, i int, lines []string, globalLowest *int64) {
	// log.Println("Checking seed range:", seedRanges[i], "-", seedRanges[i+1]+seedRanges[i])
	// counter := 0
	for seed := seedRanges[i]; seed <= seedRanges[i]+seedRanges[i+1]; seed++ {
		// log.Println("Checking seed:", seed)
		// counter++
		// if counter%1000 == 0 {
		// 	log.Println("Checked", counter, "seeds")
		// }
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

		if location < int(atomic.LoadInt64(globalLowest)) {
			atomic.StoreInt64(globalLowest, int64(location))
			log.Println("New lowest location:", location)
		}
	}
}
