package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := readInput("input.txt")
	if err {
		return
	}
	timer := time.Now()
	fmt.Println("Part 1:", part1(data))
	fmt.Println("Time taken:", time.Since(timer))

	data, err = readInput("sample.txt")
	if err {
		return
	}
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

func parseInput(data []string) ([]Workflow, []Part) {
	workflows := []Workflow{}
	parts := []Part{}

	flag := false
	for _, line := range data {
		if line == "" {
			flag = true
			continue
		}
		if !flag {
			wf := Workflow{}
			// parse Workflow
			// px{a<2006:qkq,m>2090:A,rfg}
			split := strings.Split(line, "{")
			wf.Name = split[0]
			log.Println("Name:", wf.Name)
			rules := strings.Split(split[1], ",")
			for i := 0; i < len(rules); i++ {
				rules[i] = strings.TrimSuffix(rules[i], "}")
				rule := Rule{}
				// a(<|>)2006:qkq
				re := regexp.MustCompile(`([a-zA-Z]+)(<|>)([0-9]+):([a-zA-Z]+)`)
				mats := re.FindStringSubmatch(rules[i])
				if len(mats) == 5 {
					rule.Letter = mats[1]
					rule.Symbol = mats[2]
					rule.Value, _ = strconv.Atoi(mats[3])
					rule.True = mats[4]
				} else {
					rule.True = rules[i]
				}
				log.Println("Rule:", rule)

				wf.Rules = append(wf.Rules, rule)
			}
			workflows = append(workflows, wf)
		} else {
			part := Part{}
			// parse Part
			trimmed := strings.TrimSuffix(line, "}")
			trimmed = strings.TrimPrefix(trimmed, "{")
			split := strings.Split(trimmed, ",")
			for _, s := range split {
				split2 := strings.Split(s, "=")
				switch split2[0] {
				case "x":
					part.x, _ = strconv.Atoi(split2[1])
				case "m":
					part.m, _ = strconv.Atoi(split2[1])
				case "a":
					part.a, _ = strconv.Atoi(split2[1])
				case "s":
					part.s, _ = strconv.Atoi(split2[1])
				}
			}
			log.Println("Part:", part)
			parts = append(parts, part)
		}
	}

	return workflows, parts
}

func findStartingWorkflow(workflows []Workflow) Workflow {
	for _, wf := range workflows {
		if wf.Name == "in" {
			return wf
		}
	}
	return Workflow{}
}

func findWorkflow(name string, workflows []Workflow) Workflow {
	for _, wf := range workflows {
		if wf.Name == name {
			return wf
		}
	}
	return Workflow{}
}

func checkPart(part Part, wf Workflow, workflows []Workflow) bool {
	log.Println("Checking part:", part, "in workflow:", wf)
	currentWorkflow := wf
	// loop through rules of wf
	for _, rule := range wf.Rules {
		log.Println("Rule:", rule)
		if rule.Value == 0 && rule.True == "A" {
			return true
		} else if rule.Value == 0 && rule.True == "R" {
			return false
		} else if rule.Value == 0 {
			currentWorkflow = findWorkflow(rule.True, workflows)
			log.Println("New workflow:", currentWorkflow)
			return checkPart(part, currentWorkflow, workflows)
		} else {
			switch rule.Letter {
			case "x":
				if rule.Symbol == "<" && part.x < rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				} else if rule.Symbol == ">" && part.x > rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				}
			case "m":
				if rule.Symbol == "<" && part.m < rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				} else if rule.Symbol == ">" && part.m > rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				}
			case "a":
				if rule.Symbol == "<" && part.a < rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				} else if rule.Symbol == ">" && part.a > rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				}
			case "s":
				if rule.Symbol == "<" && part.s < rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				} else if rule.Symbol == ">" && part.s > rule.Value {
					if rule.True == "A" {
						return true
					} else if rule.True == "R" {
						return false
					}
					currentWorkflow = findWorkflow(rule.True, workflows)
					log.Println("New workflow:", currentWorkflow)
					return checkPart(part, currentWorkflow, workflows)
				}
			}
		}
	}
	return false
}

func part1(data []string) int {
	workflows, parts := parseInput(data)
	log.Println("Workflows:", workflows)
	log.Println("Parts:", parts)

	inWorkflow := findStartingWorkflow(workflows)
	log.Println("Starting workflow:", inWorkflow)

	acceptedParts := []Part{}
	for _, part := range parts {
		accepted := checkPart(part, inWorkflow, workflows)
		if accepted {
			acceptedParts = append(acceptedParts, part)
		}
	}
	log.Println("Accepted parts:", acceptedParts)
	sum := 0
	for _, part := range acceptedParts {
		sum += part.x
		sum += part.m
		sum += part.a
		sum += part.s
	}

	return sum
}

func checkGroup(group MappingGroup, wf Workflow, workflows []Workflow, buffer []MappingGroup) []MappingGroup {
	currentWorkflow := wf
	log.Println("Checking group:", group, "in workflow:", currentWorkflow)

	// loop through rules of wf
	for _, rule := range wf.Rules {
		log.Println("Rule:", rule)
		if rule.Value == 0 && rule.True == "A" {
			log.Println("Accepted group:", group)
			buffer = append(buffer, MappingGroup{
				Accepted: true,
				MinX:     group.MinX,
				MaxX:     group.MaxX,
				MinM:     group.MinM,
				MaxM:     group.MaxM,
				MinA:     group.MinA,
				MaxA:     group.MaxA,
				MinS:     group.MinS,
				MaxS:     group.MaxS,
			})
		} else if rule.Value == 0 && rule.True == "R" {
			log.Println("Rejected group:", group)
			continue
		} else if rule.Value == 0 {
			currentWorkflow = findWorkflow(rule.True, workflows)
			log.Println("New workflow:", currentWorkflow)
			buffer = checkGroup(group, currentWorkflow, workflows, buffer)
		} else if !group.Accepted {
			log.Println("Spliting:", group)
			// split group into 4 groups based on rule
			switch rule.Letter {
			case "x":
				if rule.Symbol == "<" {
					// split group into 2 groups
					// group 1: minx to rule.Value
					// group 2: rule.Value+1 to maxx
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     rule.Value,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     rule.Value + 1,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else if rule.Symbol == ">" {
					// split group into 2 groups
					// group 1: minx to rule.Value-1
					// group 2: rule.Value to maxx
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     rule.Value - 1,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     rule.Value,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else {
					// split group into 2 groups
					// group 1: minx to rule.Value
					// group 2: rule.Value to maxx
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     rule.Value,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     rule.Value,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				}
			case "m":
				if rule.Symbol == "<" {
					// split group into 2 groups
					// group 1: minm to rule.Value
					// group 2: rule.Value+1 to maxm
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     rule.Value,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     rule.Value + 1,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else if rule.Symbol == ">" {
					// split group into 2 groups
					// group 1: minm to rule.Value-1
					// group 2: rule.Value to maxm
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     rule.Value - 1,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     rule.Value,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else {
					// split group into 2 groups
					// group 1: minm to rule.Value
					// group 2: rule.Value to maxm
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     rule.Value,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     rule.Value,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				}
			case "a":
				if rule.Symbol == "<" {
					// split group into 2 groups
					// group 1: mina to rule.Value
					// group 2: rule.Value+1 to maxa
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     rule.Value,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     rule.Value + 1,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else if rule.Symbol == ">" {
					// split group into 2 groups
					// group 1: mina to rule.Value-1
					// group 2: rule.Value to maxa
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     rule.Value - 1,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     rule.Value,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else {
					// split group into 2 groups
					// group 1: mina to rule.Value
					// group 2: rule.Value to maxa
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     rule.Value,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     rule.Value,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				}
			case "s":
				if rule.Symbol == "<" {
					// split group into 2 groups
					// group 1: mins to rule.Value
					// group 2: rule.Value+1 to maxs
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     rule.Value,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     rule.Value + 1,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else if rule.Symbol == ">" {
					// split group into 2 groups
					// group 1: mins to rule.Value-1
					// group 2: rule.Value to maxs
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     rule.Value - 1,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     rule.Value,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				} else {
					// split group into 2 groups
					// group 1: mins to rule.Value
					// group 2: rule.Value to maxs
					group1 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     group.MinS,
						MaxS:     rule.Value,
					}
					group2 := MappingGroup{
						Accepted: false,
						MinX:     group.MinX,
						MaxX:     group.MaxX,
						MinM:     group.MinM,
						MaxM:     group.MaxM,
						MinA:     group.MinA,
						MaxA:     group.MaxA,
						MinS:     rule.Value,
						MaxS:     group.MaxS,
					}
					buffer = checkGroup(group1, currentWorkflow, workflows, buffer)
					buffer = checkGroup(group2, currentWorkflow, workflows, buffer)
				}
			}
		}

	}

	return buffer
}

func part2(data []string) int {
	workflows, _ := parseInput(data)
	log.Println("Workflows:", workflows)

	inWorkflow := findStartingWorkflow(workflows)
	log.Println("Starting workflow:", inWorkflow)

	initial := MappingGroup{
		MinX: 1,
		MaxX: 4000,
		MinM: 1,
		MaxM: 4000,
		MinA: 1,
		MaxA: 4000,
		MinS: 1,
		MaxS: 4000,
	}

	groups := checkGroup(initial, inWorkflow, workflows, []MappingGroup{})
	log.Println("Groups:", groups)
	return 0
}

type Workflow struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	Letter string
	Symbol string
	Value  int
	True   string // points to a Workflow name or A or R
}

type Part struct {
	x int
	m int
	a int
	s int
}

type MappingGroup struct {
	Accepted bool
	MinX     int
	MaxX     int
	MinM     int
	MaxM     int
	MinA     int
	MaxA     int
	MinS     int
	MaxS     int
}
