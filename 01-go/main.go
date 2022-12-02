package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func elvesFromFile(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	foodItems := make([]int, 0)
	totalFood := 0
	food := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			foodItems = append(foodItems, totalFood)
			totalFood = 0
		} else {
			food, err = strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			totalFood += food
		}
	}
	foodItems = append(foodItems, food)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return foodItems
}

func findMaxElf(elves []int) (elf, maxTotal int) {
	elf = 0
	maxTotal = math.MinInt
	for key, val := range elves {
		if val > maxTotal {
			elf = key
			maxTotal = val
		}
	}
	return elf, maxTotal
}

func solvePartOne(filePath string) (elf, result int) {
	elves := elvesFromFile(filePath)
	return findMaxElf(elves)
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func findTopNElves(elves []int, n int) (maxTotal int) {
	newN := int(math.Min(float64(n), float64(len(elves))))
	for i := 0; i < newN; i++ {
		elf, total := findMaxElf(elves)
		maxTotal += total
		elves = remove(elves, elf)
	}
	return maxTotal
}

func solvePartTwo(filePath string) (result int) {
	elves := elvesFromFile(filePath)
	return findTopNElves(elves, 3)
}

func main() {
	for _, filePath := range os.Args[1:] {
		fmt.Printf("Input file: %s\n", filePath)

		fmt.Printf("Part One:\n")
		fmt.Println(solvePartOne(filePath))

		fmt.Printf("\nPart Two:\n")
		fmt.Println(solvePartTwo(filePath))
	}
}
