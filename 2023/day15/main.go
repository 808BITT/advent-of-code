package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2: ", part2("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
}

func part1(filename string) int {
	input := readInput(filename)
	// fmt.Println(input)
	sequence := strings.Split(input, ",")
	// fmt.Println(sequence)
	sum := 0
	for _, command := range sequence {
		// fmt.Println(hash(command))
		sum += hash(command)
	}
	return sum
}

func readInput(filename string) string {
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	val := ""
	for _, line := range lines {
		val += line
	}
	return val
}

func hash(input string) int {
	val := 0
	for _, char := range input {
		ascii := int(char)
		val += ascii
		val *= 17
		val %= 256
	}
	return val
}

func part2(filename string) int {
	lenses := make(LensMap)

	input := readInput(filename)
	// fmt.Println(input)
	sequence := strings.Split(input, ",")
	// fmt.Println(sequence)

	for _, s := range sequence {
		command := ""
		split := strings.Split(s, "=")
		if len(split) == 2 {
			command = "ASSIGN"
		} else {
			command = "REMOVE"
		}
		// fmt.Println(command, s)
		lenses = applyCommand(lenses, command, s)
		// fmt.Println(lenses)

	}
	totalPower := 0
	for b, box := range lenses {
		for l, lens := range box.Lenses {
			// fmt.Println(b, l, lens)
			focalPower := (b + 1) * (l + 1) * lens.Length
			// fmt.Println(focalPower)
			totalPower += focalPower
			// fmt.Scanln()
		}
	}

	return totalPower
}

func applyCommand(lenses LensMap, command string, s string) LensMap {
	if command == "ASSIGN" {
		lenses = assignLens(lenses, s)
	} else {
		lenses = removeLens(lenses, s)
	}
	return lenses
}

func assignLens(lenses LensMap, s string) LensMap {
	split := strings.Split(s, "=")
	hash := hash(split[0])
	lens := Lens{Sequence: split[0], Length: atoi(split[1])}
	if _, ok := lenses[hash]; !ok {
		lenses[hash] = LensBox{Lenses: []Lens{lens}}
	} else {
		// swap lens if the sequence is already in the box
		for _, l := range lenses[hash].Lenses {
			if l.Sequence == lens.Sequence {
				box := lenses[hash]
				for i, l := range box.Lenses {
					if l.Sequence == lens.Sequence {
						box.Lenses[i] = lens
						lenses[hash] = box
						return lenses
					}
				}
			}
		}
		// if not, add it
		lenses[hash] = LensBox{Lenses: append(lenses[hash].Lenses, lens)}
	}
	return lenses
}

func removeLens(lenses LensMap, s string) LensMap {
	split := strings.Split(s, "-")
	hash := hash(split[0])
	lens := Lens{Sequence: split[0]}
	if _, ok := lenses[hash]; !ok {
		return lenses
	} else {
		// check if the Lens is in the box
		for i, l := range lenses[hash].Lenses {
			if l.Sequence == lens.Sequence {
				box := lenses[hash]
				box.Lenses = append(box.Lenses[:i], box.Lenses[i+1:]...)
				if len(box.Lenses) == 0 {
					delete(lenses, hash)
					return lenses
				}
				lenses[hash] = box
				return lenses
			}
		}
	}
	return lenses
}

type LensMap map[int]LensBox

type LensBox struct {
	Lenses []Lens
}

type Lens struct {
	Sequence string
	Length   int
}

func atoi(s string) int {
	val := 0
	for _, char := range s {
		val *= 10
		val += int(char) - 48
	}
	return val
}
