package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dataFromFile(filePath string) (dataStr []string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dataStr = append(dataStr, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return dataStr
}

func hasRepeats(marker string) bool {
	for i, rn := range marker {
		if strings.Contains(marker[i+1:], string(rn)) {
			return true
		}
	}
	return false
}

func solveBoth(filePath string, n int) (result []int) {
	dataStr := dataFromFile(filePath)
	for _, data := range dataStr {
		for i := 0; i < len(data); i++ {
			marker := data[i : i+n]
			if !hasRepeats(marker) {
				numChar := i + n
				result = append(result, numChar)
				break
			}
		}
	}
	return result
}

func main() {
	for _, filePath := range os.Args[1:] {
		fmt.Printf("Input: %s\n", filePath)
		fmt.Println("Part One:", solveBoth(filePath, 4))
		fmt.Println("Part Two:", solveBoth(filePath, 14))
	}
}
