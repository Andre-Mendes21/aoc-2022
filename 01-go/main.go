package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	sort.SliceStable(foodItems, func(i, j int) bool { return foodItems[i] > foodItems[j] })
	return foodItems
}

func findTopNElves(elves []int, n int) (maxTotal int) {
	newN := int(math.Min(float64(n), float64(len(elves))))
	for i := 0; i < newN; i++ {
		maxTotal += elves[i]
	}
	return maxTotal
}

func solvePartOne(filePath string) (result int) {
	elves := elvesFromFile(filePath)
	return findTopNElves(elves, 1)
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
