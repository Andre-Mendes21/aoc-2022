package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func itemToPriority(item rune) int {
	var priority int
	switch {
	case item >= 'A' && item <= 'Z':
		priority = int(item) - 38
	case item >= 'a' && item <= 'z':
		priority = int(item) - 96
	}

	return priority
}

func removeRepeatRunes(str string) string {
	var newString string
	for i, rn := range str {
		cutoff := strings.Index(str[i+1:], string(rn))
		if cutoff == -1 {
			newString += string(rn)
		}
	}

	return newString
}

func findCommonItem(bag string) (commonItems int) {
	compart1 := bag[:len(bag)/2]
	compart2 := bag[len(bag)/2:]
	for _, vComp1 := range compart1 {
		if strings.Contains(compart2, string(vComp1)) {
			priority := itemToPriority(vComp1)
			commonItems += priority
			break
		}
	}

	return commonItems
}

func findBadge(bags []string) (acc int) {
	noDupBags := []string{}
	for _, bag := range bags {
		noDupBag := removeRepeatRunes(bag)
		noDupBags = append(noDupBags, noDupBag)
	}

	for _, rn := range noDupBags[0] {
		if strings.Contains(noDupBags[1], string(rn)) && strings.Contains(noDupBags[2], string(rn)) {
			priority := itemToPriority(rn)
			acc += priority
			break
		}
	}

	return acc
}

func bagsFromFile(filePath string) (bags []string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bags = append(bags, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return bags
}

func solvePartOne(filePath string) (acc int) {
	bags := bagsFromFile(filePath)
	for _, str := range bags {
		acc += findCommonItem(str)
	}

	return acc
}

func solvePartTwo(filePath string) (acc int) {
	bags := bagsFromFile(filePath)
	for i := 0; i < len(bags); i += 3 {
		acc += findBadge(bags[i : i+3])
	}

	return acc
}

func main() {
	for _, filePath := range os.Args[1:] {
		fmt.Printf("Input file: %s\n", filePath)
		fmt.Printf("Part One\n")
		fmt.Println(solvePartOne(filePath))

		fmt.Printf("\nPart Two\n")
		fmt.Println(solvePartTwo(filePath))
	}
}
