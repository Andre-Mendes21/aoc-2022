package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
* A, X = Rock = 1
* B, Y = Paper = 2
* C, Z = Scissors = 3
 */

const (
	WON      = 6
	DREW     = 3
	LOST     = 0
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func roundToObject(round string) int {
	switch round {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISSORS
	default:
		return -1
	}
}

func roundToScore(round string) int {
	switch round {
	case "X":
		return LOST
	case "Y":
		return DREW
	case "Z":
		return WON
	default:
		return -1
	}
}

func roundsFromFile(filePath string) map[int][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rounds = map[int][]string{}
	X, Y, Z := make([]string, 0), make([]string, 0), make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Split(line, " ")
		switch round[0] {
		case "A":
			X = append(X, round[1])
		case "B":
			Y = append(Y, round[1])
		case "C":
			Z = append(Z, round[1])
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	rounds[1] = X
	rounds[2] = Y
	rounds[3] = Z

	return rounds
}

func objectToWin(object int) int {
	var opposite int
	switch object {
	case ROCK:
		opposite = PAPER
	case PAPER:
		opposite = SCISSORS
	case SCISSORS:
		opposite = ROCK
	}
	return opposite
}

func objectToLose(object int) int {
	return objectToWin(objectToWin(object))
}

func totalScore(rounds map[int][]string) int {
	var totalScore int
	for key, val := range rounds {
		for _, v := range val {
			object := roundToObject(v)
			totalScore += object
			switch object {
			case key:
				totalScore += DREW
			case objectToWin(key):
				totalScore += WON
			default:
				totalScore += LOST
			}
		}
	}

	return totalScore
}

func solvePartOne(filepath string) int {
	rounds := roundsFromFile(filepath)
	return totalScore(rounds)
}

func totalScore2(rounds map[int][]string) int {
	var totalScore int
	for key, val := range rounds {
		for _, v := range val {
			score := roundToScore(v)
			totalScore += score
			switch score {
			case WON:
				totalScore += objectToWin(key)
			case DREW:
				totalScore += key
			case LOST:
				totalScore += objectToLose(key)
			}
		}
	}
	return totalScore
}

func solvePartTwo(filepath string) int {
	rounds := roundsFromFile(filepath)
	return totalScore2(rounds)
}

func main() {
	for _, filepath := range os.Args[1:] {
		fmt.Printf("Input File: %s\n", filepath)
		fmt.Printf("Part One:\n")
		fmt.Println(solvePartOne(filepath))

		fmt.Printf("\nPart Two:\n")
		fmt.Println(solvePartTwo(filepath))
	}
}
