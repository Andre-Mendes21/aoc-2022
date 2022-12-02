package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func foodItemsToMap(foodItems []int) map[int]int {
	elfs := map[int]int{}
	for i, j := 0, 0; i < len(foodItems); i++ {
		eachElfFood := 0
		for foodItems[i] != -1 {
			eachElfFood += foodItems[i]
			i++
		}
		elfs[j] = eachElfFood
		j++
	}
	return elfs
}

func elvesFromFile(filePath string) map[int]int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	foodItems := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "":
			foodItems = append(foodItems, -1)
		default:
			food, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			foodItems = append(foodItems, food)
		}
	}
	foodItems = append(foodItems, -1)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return foodItemsToMap(foodItems)
}

func findMaxElf(elves map[int]int) (elf, maxTotal int) {
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

func findTopNElves(elves map[int]int, n int) (maxTotal int) {
	maxTotal = 0
	newN := int(math.Min(float64(n), float64(len(elves))))
	for i := 0; i < newN; i++ {
		elf, total := findMaxElf(elves)
		delete(elves, elf)
		maxTotal += total
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
		// First Part
		fmt.Println(solvePartOne(filePath))
		// Second Part
		fmt.Println(solvePartTwo(filePath))
	}
}
