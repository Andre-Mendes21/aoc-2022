package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	DISKSPACE = 70000000
	UPDATE    = 30000000
)

func moveToDir(dir string, currentDir *Dir) *Dir {
	if dir == currentDir.name {
		return currentDir
	}
	switch dir {
	case "..":
		return currentDir.goToParent()
	default:
		return currentDir.goToChild(dir)
	}
}

func getDirContents(dirContent []string, currentDir *Dir) {
	for _, content := range dirContent {
		split := strings.Split(content, " ")
		switch split[0] {
		case "dir":
			newDir := NewDir(split[1], currentDir, nil, nil)
			currentDir.AddNewDir(newDir)
		default:
			size, _ := strconv.Atoi(split[0])
			newFile := NewFile(size, split[1])
			currentDir.AddNewFile(&newFile)
		}
	}
}

func terminalFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var terminalInput = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		terminalInput = append(terminalInput, line)
	}
	return terminalInput
}

func findNext(strs []string, s string) int {
	for i := range strs {
		if strings.Contains(strs[i], s) {
			return i + 1
		}
	}
	return -1
}

func createFileSystem(filePath string) *Dir {
	terminalInput := terminalFromFile(filePath)

	var root *Dir = NewDir("/", nil, nil, nil)
	currentDir := root
	for i, line := range terminalInput {
		command := line[2:4]
		switch command {
		case "cd":
			dir := line[5:]
			currentDir = moveToDir(dir, currentDir)
		case "ls":
			if idx := findNext(terminalInput[i+1:], "$"); idx != -1 {
				getDirContents(terminalInput[i+1:idx+i], currentDir)
			} else {
				getDirContents(terminalInput[i+1:], currentDir)
			}
		}
	}
	return root
}

func findDirsSizeCmpN(root *Dir, found *[]int, n int, cmp func(int, int) bool) {
	if cmp(root.getTotalSize(), n) {
		*found = append(*found, root.getTotalSize())
	}
	for _, child := range root.children {
		findDirsSizeCmpN(child, found, n, cmp)
	}
}

func findDirsSizeLessThanEqN(root *Dir, found *[]int, n int) {
	findDirsSizeCmpN(root, found, n, func(i, j int) bool { return i <= j })
}

func findDirsSizeGreaterThanEqN(root *Dir, found *[]int, n int) {
	findDirsSizeCmpN(root, found, n, func(i, j int) bool { return i >= j })
}

func solvePartOne(filePath string) (sum int) {
	root := createFileSystem(filePath)
	var smallDirs = []int{}
	findDirsSizeLessThanEqN(root, &smallDirs, 100000)

	for _, totalSize := range smallDirs {
		sum += totalSize
	}
	return sum
}

func solvePartTwo(filePath string) (size int) {
	root := createFileSystem(filePath)
	unusedSpace := DISKSPACE - root.getTotalSize()
	spaceNeeded := UPDATE - unusedSpace
	var smallestDirs = []int{}
	findDirsSizeGreaterThanEqN(root, &smallestDirs, spaceNeeded)
	currentUnusedSpace := math.MinInt
	for _, dirSize := range smallestDirs {
		spaceIfRemoved := unusedSpace + (root.getTotalSize() - dirSize)
		if spaceIfRemoved > currentUnusedSpace {
			currentUnusedSpace = spaceIfRemoved
			size = dirSize
		}
	}
	return size
}

func main() {
	for _, filePath := range os.Args[1:] {
		fmt.Printf("Input file: %s\n", filePath)
		fmt.Println("Part One:", solvePartOne(filePath))
		fmt.Println("Part Two:", solvePartTwo(filePath))
	}
}
