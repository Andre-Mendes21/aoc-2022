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

func sectionsFromFile(filePath string) (sections1, sections2 []Section) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitSection := strings.Split(line, ",")
		ids1 := strings.Split(splitSection[0], "-")
		ids2 := strings.Split(splitSection[1], "-")

		section1 := newSection(ids1[0], ids1[1])
		section2 := newSection(ids2[0], ids2[1])

		sections1 = append(sections1, section1)
		sections2 = append(sections2, section2)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return sections1, sections2
}

func isContained(sects1, sects2 Section) bool {
	sect1IsContained := sects1.lo >= sects2.lo && sects1.hi <= sects2.hi
	sect2IsContained := sects2.lo >= sects1.lo && sects2.hi <= sects1.hi
	return sect1IsContained || sect2IsContained
}

func isOverlap(sects1, sects2 Section) bool {
	sect1IsContained := sects1.lo <= sects2.hi && sects1.hi >= sects2.lo
	sect2IsContained := sects2.lo <= sects1.hi && sects2.hi >= sects1.lo
	return sect1IsContained || sect2IsContained
}

func solveBoth(filePath string, cmp func(sect1, sects2 Section) bool) (count int) {
	sects1, sects2 := sectionsFromFile(filePath)
	for i := range sects1 {
		if cmp(sects1[i], sects2[i]) {
			count++
		}
	}
	return count
}

func main() {
	for _, filepath := range os.Args[1:] {
		fmt.Printf("Input file: %s\n", filepath)
		fmt.Println("Part One")
		fmt.Println(solveBoth(filepath, isContained))

		fmt.Println("\nPart Two")
		fmt.Println(solveBoth(filepath, isOverlap))
	}
}
