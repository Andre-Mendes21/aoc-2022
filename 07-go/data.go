package main

import (
	"fmt"
	"strings"
)

type File struct {
	size int
	name string
}

type Dir struct {
	name      string
	parent    *Dir
	files     map[string]int
	children  map[string]*Dir
	totalSize int
}

func NewFile(size int, name string) File {
	return File{
		size: size,
		name: name,
	}
}

func NewDir(dirName string, parent *Dir, files []File, children []*Dir) *Dir {
	var allFiles = make(map[string]int)
	for _, file := range files {
		allFiles[file.name] = file.size
	}

	var allDirs = make(map[string]*Dir)
	for _, dir := range children {
		allDirs[dir.name] = dir
	}
	return &Dir{
		name:      dirName,
		parent:    parent,
		files:     allFiles,
		children:  allDirs,
		totalSize: 0,
	}
}

func tree(dir *Dir, str *string, indentLevel int) {
	indentLevel++
	indent := strings.Repeat("\x20\x20\x20\x20", indentLevel)
	parent := ""
	if dir.parent == nil {
		parent = "nil"
	} else {
		parent = dir.parent.name
	}
	*str += "- " + dir.name + " (dir, parent=" + parent + ")\n"
	for _, child := range dir.children {
		*str += indent
		tree(child, str, indentLevel)
	}
	for name, size := range dir.files {
		*str += indent
		*str += "- " + name + "(file, size=" + fmt.Sprint(size) + " parent=" + dir.name + ")\n"
	}
}

func Tree(dir *Dir) string {
	str := ""
	tree(dir, &str, 0)
	return str
}

func (dir *Dir) AddNewDir(newDir *Dir) {
	if _, ok := dir.children[newDir.name]; ok {
		fmt.Printf("Directory \"%s\" already exists in \"%s\"\n", newDir.name, dir.name)
		return
	}
	dir.children[newDir.name] = newDir
}

func (dir *Dir) backPropagateSize(size int) {
	if dir.parent != nil {
		dir.parent.totalSize += size
		dir.parent.backPropagateSize(size)
	}
}

func (dir *Dir) AddNewFile(newFile *File) {
	if _, ok := dir.files[newFile.name]; ok {
		fmt.Printf("File \"%s\" already exists in \"%s\"\n", newFile.name, dir.name)
		return
	}
	dir.totalSize += newFile.size
	dir.backPropagateSize(newFile.size)
	dir.files[newFile.name] = newFile.size
}

func (dir *Dir) goToParent() *Dir {
	if dir.parent == nil {
		return dir
	}
	return dir.parent
}

func (dir *Dir) goToChild(dirName string) *Dir {
	child, ok := dir.children[dirName]
	if !ok {
		fmt.Printf("Directory \"%s\" does not exist in \"%s\"\n", dirName, dir.name)
		return dir
	}
	return child
}

func (dir *Dir) getTotalSize() int {
	return dir.totalSize
}
