package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	lo int
	hi int
}

func newSection(lo, hi string) Section {
	low, _ := strconv.Atoi(lo)
	high, _ := strconv.Atoi(hi)
	return Section{
		lo: low,
		hi: high,
	}
}

func parseRangePair(str []string) (left, right []string) {
	left = strings.Split(str[0], "-")
	right = strings.Split(str[1], "-")
	return left, right
}

func sectionsFromFile(filePath string) (sections1, sections2 []Section) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitSection := strings.Split(line, ",")
		left, right := parseRangePair(splitSection)
		sections1 = append(sections1, newSection(left[0], left[1]))
		sections2 = append(sections2, newSection(right[0], right[1]))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return sections1, sections2
}

func isContained(sects1, sects2 Section) bool {
	return sects1.lo <= sects2.lo && sects2.hi <= sects1.hi
}

func isOverlap(sects1, sects2 Section) bool {
	return sects1.lo <= sects2.lo && sects2.lo <= sects1.hi
}

func solveBoth(filePath string, cmp func(sect1, sects2 Section) bool) (count int) {
	sects1, sects2 := sectionsFromFile(filePath)
	for i := range sects1 {
		if cmp(sects1[i], sects2[i]) || cmp(sects2[i], sects1[i]) {
			count++
		}
	}
	return count
}

func main() {
	for _, filepath := range os.Args[1:] {
		fmt.Printf("\nInput file: %s\n", filepath)
		fmt.Println("Part One:", solveBoth(filepath, isContained))
		fmt.Println("Part Two:", solveBoth(filepath, isOverlap))
	}
}
