package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timer := time.Now()
	fmt.Println("Part 1:", part1("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
	timer = time.Now()
	fmt.Println("Part 2:", part2("input.txt"))
	fmt.Println("Time: ", time.Since(timer))
}

func part1(filename string) int {
	lines := readFile(filename)

	points := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		opp, usr := split[0], split[1]
		points += evaluateWinner(opp, usr)
	}
	return points
}

func part2(filename string) int {
	lines := readFile(filename)

	points := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		opp, res := split[0], split[1]
		points += findMove(opp, res)
	}
	return points
}

func readFile(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func evaluateWinner(opp, usr string) int {
	var user string
	var outcome string

	switch opp {
	case "A":
		switch usr {
		case "X":
			user = "Rock"
			outcome = "Tie"
		case "Y":
			user = "Paper"
			outcome = "Win"
		case "Z":
			user = "Scissors"
			outcome = "Lose"
		}
	case "B":
		switch usr {
		case "X":
			user = "Rock"
			outcome = "Lose"
		case "Y":
			user = "Paper"
			outcome = "Tie"
		case "Z":
			user = "Scissors"
			outcome = "Win"
		}
	case "C":
		switch usr {
		case "X":
			user = "Rock"
			outcome = "Win"
		case "Y":
			user = "Paper"
			outcome = "Lose"
		case "Z":
			user = "Scissors"
			outcome = "Tie"
		}
	}

	// fmt.Println("Opponent:", opponent)
	// fmt.Println("User:", user)
	// fmt.Println("Outcome:", outcome)

	points := 0
	switch outcome {
	case "Win":
		points += 6
	case "Lose":
		points += 0
	case "Tie":
		points += 3
	}
	switch user {
	case "Rock":
		points += 1
	case "Paper":
		points += 2
	case "Scissors":
		points += 3
	}
	return points
}

func findMove(opp, res string) int {
	points := 0
	switch opp {
	case "A": // Rock
		switch res {
		case "X": // Lose using Scissors
			points += 0 + 3
		case "Y":
			points += 3 + 1
		case "Z":
			points += 6 + 2
		}
	case "B": // Paper
		switch res {
		case "X": // Lose using Rock
			points += 0 + 1
		case "Y":
			points += 3 + 2
		case "Z":
			points += 6 + 3
		}
	case "C": // Scissors
		switch res {
		case "X": // Lose using Paper
			points += 0 + 2
		case "Y":
			points += 3 + 3
		case "Z":
			points += 6 + 1
		}
	}
	return points
}
