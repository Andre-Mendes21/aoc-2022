package main

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) peek() string {
	return (*s)[len(*s)-1]
}

func SampleState() []string {
	return []string{"ZN", "MCD", "P"}
}

func SampleMoves() [][3]int {
	return [][3]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}
}

func TestState() []string {
	return []string{
		"BLDTWCFM",
		"NBL",
		"JCHTLV",
		"SPJW",
		"ZSCFTLR",
		"WDGBHNZ",
		"FMSPVGCN",
		"WQRJFVCZ",
		"RPMLH",
	}
}

func TestMoves() [][3]int {
	return [][3]int{
		{5, 3, 6},
		{2, 2, 5},
		{1, 9, 1},
		{1, 3, 1},
		{5, 7, 5},
		{2, 9, 8},
		{1, 2, 8},
		{1, 4, 2},
		{8, 1, 6},
		{4, 6, 9},
		{1, 2, 1},
		{2, 4, 8},
		{2, 8, 4},
		{3, 7, 5},
		{6, 5, 3},
		{1, 1, 8},
		{1, 5, 7},
		{5, 6, 9},
		{3, 5, 8},
		{2, 4, 3},
		{1, 7, 8},
		{2, 8, 6},
		{2, 1, 8},
		{8, 3, 8},
		{11, 6, 3},
		{1, 4, 7},
		{1, 3, 7},
		{2, 6, 1},
		{7, 9, 7},
		{10, 3, 5},
		{1, 9, 3},
		{2, 9, 5},
		{5, 5, 2},
		{19, 8, 6},
		{1, 9, 6},
		{1, 3, 8},
		{4, 2, 6},
		{1, 1, 4},
		{5, 8, 9},
		{1, 2, 1},
		{6, 7, 2},
		{3, 5, 8},
		{3, 8, 1},
		{2, 9, 6},
		{1, 7, 8},
		{6, 2, 7},
		{1, 4, 8},
		{3, 8, 4},
		{2, 1, 5},
		{7, 7, 6},
		{1, 7, 2},
		{3, 4, 6},
		{2, 9, 2},
		{1, 1, 8},
		{2, 1, 3},
		{1, 8, 7},
		{3, 2, 5},
		{5, 5, 8},
		{4, 5, 3},
		{1, 7, 8},
		{2, 8, 1},
		{1, 8, 5},
		{5, 3, 5},
		{13, 5, 1},
		{1, 3, 4},
		{2, 8, 3},
		{3, 1, 4},
		{1, 3, 1},
		{1, 8, 1},
		{5, 1, 9},
		{1, 3, 7},
		{2, 9, 6},
		{2, 1, 7},
		{3, 1, 5},
		{3, 1, 5},
		{1, 6, 1},
		{4, 4, 3},
		{3, 9, 1},
		{5, 1, 7},
		{7, 7, 8},
		{1, 3, 9},
		{28, 6, 8},
		{5, 5, 9},
		{6, 6, 1},
		{4, 1, 8},
		{5, 9, 1},
		{12, 8, 7},
		{1, 3, 8},
		{6, 1, 4},
		{5, 4, 1},
		{3, 6, 4},
		{2, 3, 4},
		{3, 1, 5},
		{6, 7, 1},
		{2, 4, 9},
		{2, 5, 4},
		{19, 8, 1},
		{4, 9, 5},
		{5, 4, 3},
		{4, 1, 4},
		{5, 5, 1},
		{3, 8, 5},
		{7, 7, 3},
		{14, 1, 8},
		{5, 4, 2},
		{12, 8, 7},
		{1, 3, 6},
		{3, 5, 9},
		{1, 7, 8},
		{8, 1, 2},
		{5, 1, 2},
		{9, 3, 4},
		{8, 4, 6},
		{2, 1, 9},
		{3, 6, 1},
		{5, 6, 7},
		{14, 7, 1},
		{1, 4, 7},
		{6, 8, 2},
		{14, 1, 4},
		{13, 4, 9},
		{2, 3, 5},
		{3, 1, 7},
		{1, 8, 4},
		{1, 4, 1},
		{1, 1, 3},
		{1, 3, 4},
		{1, 4, 1},
		{1, 6, 9},
		{1, 7, 6},
		{1, 4, 5},
		{11, 9, 3},
		{6, 3, 8},
		{5, 3, 1},
		{2, 8, 4},
		{1, 6, 2},
		{7, 9, 2},
		{1, 7, 2},
		{1, 9, 8},
		{2, 8, 6},
		{30, 2, 3},
		{2, 7, 2},
		{2, 8, 2},
		{3, 8, 7},
		{6, 2, 5},
		{1, 2, 5},
		{3, 1, 8},
		{2, 6, 7},
		{1, 1, 9},
		{1, 9, 3},
		{7, 3, 1},
		{6, 7, 8},
		{8, 3, 9},
		{7, 9, 1},
		{1, 5, 8},
		{7, 5, 9},
		{2, 4, 2},
		{11, 3, 6},
		{2, 2, 7},
		{11, 1, 8},
		{2, 5, 4},
		{11, 6, 4},
		{12, 4, 9},
		{4, 1, 5},
		{3, 7, 9},
		{12, 8, 4},
		{1, 1, 7},
		{6, 8, 3},
		{2, 3, 5},
		{3, 8, 4},
		{3, 3, 7},
		{9, 9, 7},
		{5, 3, 9},
		{1, 3, 2},
		{13, 7, 5},
		{1, 2, 6},
		{1, 6, 1},
		{1, 1, 6},
		{16, 4, 5},
		{1, 5, 6},
		{16, 5, 4},
		{13, 4, 5},
		{3, 4, 2},
		{1, 6, 7},
		{3, 2, 1},
		{8, 5, 2},
		{3, 1, 4},
		{1, 7, 9},
		{14, 5, 1},
		{10, 1, 5},
		{1, 2, 8},
		{19, 9, 1},
		{1, 9, 1},
		{6, 2, 7},
		{4, 1, 7},
		{1, 8, 6},
		{16, 5, 3},
		{1, 5, 4},
		{2, 5, 2},
		{1, 5, 6},
		{1, 6, 5},
		{1, 2, 4},
		{7, 7, 2},
		{4, 4, 7},
		{2, 6, 2},
		{8, 2, 9},
		{4, 9, 2},
		{16, 3, 7},
		{4, 9, 7},
		{14, 1, 3},
		{26, 7, 8},
		{1, 5, 4},
		{20, 8, 4},
		{5, 1, 8},
		{2, 4, 6},
		{4, 3, 2},
		{1, 6, 5},
		{8, 2, 4},
		{1, 6, 5},
		{1, 7, 8},
		{8, 3, 1},
		{6, 1, 9},
		{1, 3, 6},
		{14, 4, 1},
		{1, 3, 8},
		{2, 2, 1},
		{1, 6, 8},
		{1, 2, 8},
		{5, 8, 1},
		{2, 1, 6},
		{2, 5, 9},
		{1, 6, 3},
		{1, 6, 1},
		{5, 9, 2},
		{5, 4, 1},
		{5, 4, 2},
		{16, 1, 8},
		{9, 1, 4},
		{24, 8, 6},
		{1, 8, 7},
		{7, 6, 5},
		{1, 3, 4},
		{3, 1, 8},
		{3, 5, 8},
		{10, 4, 8},
		{3, 4, 6},
		{1, 7, 4},
		{20, 6, 7},
		{1, 4, 9},
		{1, 4, 9},
		{7, 2, 3},
		{13, 8, 9},
		{4, 5, 9},
		{4, 8, 5},
		{18, 9, 2},
		{14, 7, 5},
		{6, 3, 8},
		{1, 3, 2},
		{1, 8, 6},
		{4, 8, 2},
		{1, 2, 3},
		{17, 5, 3},
		{18, 3, 5},
		{6, 7, 2},
		{3, 9, 7},
		{1, 8, 6},
		{5, 2, 5},
		{26, 2, 7},
		{1, 6, 9},
		{29, 7, 9},
		{15, 5, 2},
		{1, 6, 7},
		{8, 9, 2},
		{14, 2, 6},
		{16, 9, 1},
		{6, 9, 1},
		{1, 7, 1},
		{3, 2, 1},
		{5, 2, 6},
		{15, 1, 4},
		{1, 2, 8},
		{1, 9, 7},
		{1, 8, 6},
		{19, 6, 7},
		{10, 1, 8},
		{4, 8, 3},
		{1, 7, 5},
		{3, 5, 3},
		{13, 7, 6},
		{2, 8, 9},
		{7, 3, 6},
		{5, 5, 3},
		{1, 1, 6},
		{2, 5, 1},
		{4, 4, 8},
		{7, 8, 7},
		{8, 7, 3},
		{1, 8, 4},
		{2, 9, 2},
		{8, 6, 5},
		{1, 4, 5},
		{4, 5, 4},
		{2, 2, 8},
		{9, 4, 5},
		{2, 1, 9},
		{2, 8, 9},
		{14, 6, 4},
		{5, 3, 4},
		{3, 9, 7},
		{3, 5, 3},
		{2, 4, 8},
		{2, 4, 7},
		{2, 8, 9},
		{4, 5, 8},
		{16, 4, 6},
		{1, 9, 6},
		{3, 7, 5},
		{7, 7, 5},
		{10, 5, 1},
		{6, 3, 8},
		{2, 9, 3},
		{3, 6, 9},
		{3, 3, 6},
		{2, 1, 7},
		{13, 6, 2},
		{2, 4, 5},
		{2, 7, 6},
		{2, 6, 7},
		{2, 4, 1},
		{3, 9, 5},
		{1, 1, 4},
		{3, 2, 5},
		{2, 4, 1},
		{2, 3, 2},
		{5, 8, 5},
		{1, 7, 2},
		{1, 7, 1},
		{1, 3, 5},
		{1, 8, 7},
		{1, 6, 7},
		{1, 3, 5},
		{12, 5, 6},
		{6, 6, 2},
		{1, 7, 4},
		{1, 5, 7},
		{2, 8, 9},
		{1, 9, 6},
		{1, 8, 9},
		{5, 6, 9},
		{1, 8, 1},
		{14, 2, 4},
		{1, 7, 1},
		{1, 7, 2},
		{3, 2, 3},
		{2, 3, 4},
		{1, 2, 4},
		{4, 6, 2},
		{8, 5, 8},
		{15, 4, 8},
		{3, 4, 8},
		{7, 8, 4},
		{6, 1, 3},
		{1, 6, 1},
		{5, 4, 8},
		{7, 9, 1},
		{1, 5, 6},
		{4, 2, 6},
		{10, 1, 8},
		{29, 8, 3},
		{1, 4, 5},
		{1, 4, 6},
		{6, 1, 4},
		{1, 5, 8},
		{3, 4, 2},
		{27, 3, 7},
		{18, 7, 9},
		{5, 6, 3},
		{7, 7, 4},
		{1, 7, 8},
		{9, 3, 5},
		{5, 3, 6},
		{3, 4, 2},
		{1, 7, 2},
		{2, 8, 4},
		{2, 8, 6},
		{2, 8, 6},
		{8, 2, 1},
		{7, 5, 4},
		{1, 8, 9},
		{4, 1, 5},
		{1, 2, 9},
		{8, 6, 3},
		{3, 1, 8},
		{1, 1, 7},
		{8, 3, 6},
		{2, 8, 3},
		{1, 3, 6},
		{4, 6, 7},
		{16, 4, 2},
		{1, 3, 5},
		{2, 6, 4},
		{1, 2, 3},
		{2, 7, 3},
		{2, 7, 8},
		{3, 6, 7},
		{4, 5, 2},
		{2, 4, 2},
		{4, 9, 8},
		{3, 5, 1},
		{3, 1, 6},
		{6, 9, 1},
		{4, 7, 9},
		{8, 9, 5},
		{4, 5, 2},
		{7, 8, 6},
		{11, 6, 8},
		{4, 1, 2},
		{3, 8, 9},
		{5, 8, 7},
		{2, 1, 6},
		{4, 5, 6},
		{2, 7, 9},
		{2, 7, 3},
		{5, 6, 2},
		{4, 3, 1},
		{1, 7, 2},
		{1, 3, 2},
		{2, 6, 7},
		{1, 1, 6},
		{6, 9, 6},
		{1, 7, 6},
		{1, 7, 6},
		{2, 1, 7},
		{2, 8, 6},
		{4, 9, 2},
		{17, 2, 6},
		{1, 9, 4},
		{1, 1, 3},
		{1, 4, 1},
		{20, 2, 8},
		{2, 7, 6},
		{2, 2, 5},
		{1, 3, 1},
		{1, 2, 5},
		{6, 8, 6},
		{2, 5, 6},
		{3, 6, 4},
		{1, 1, 4},
		{15, 8, 2},
		{11, 2, 9},
		{1, 1, 3},
		{10, 9, 4},
		{1, 9, 8},
		{12, 6, 3},
		{1, 8, 7},
		{1, 5, 4},
		{8, 4, 7},
		{5, 3, 4},
		{7, 6, 4},
		{3, 3, 6},
		{3, 3, 2},
		{1, 3, 6},
		{17, 4, 3},
		{1, 3, 4},
		{2, 4, 9},
		{14, 3, 6},
		{2, 2, 7},
		{1, 4, 9},
		{8, 7, 6},
		{1, 3, 4},
		{9, 6, 2},
		{1, 4, 2},
		{26, 6, 2},
		{27, 2, 6},
		{10, 2, 4},
		{1, 7, 6},
		{28, 6, 2},
		{21, 2, 4},
		{2, 6, 7},
		{3, 2, 1},
		{5, 6, 5},
		{3, 5, 2},
		{1, 7, 4},
		{11, 2, 4},
		{21, 4, 9},
		{1, 5, 8},
		{1, 8, 6},
		{18, 9, 7},
		{1, 5, 7},
		{3, 9, 8},
		{1, 6, 7},
		{1, 3, 5},
		{1, 8, 3},
		{22, 7, 5},
		{13, 5, 1},
		{16, 4, 5},
		{3, 1, 4},
		{2, 3, 9},
		{3, 9, 7},
		{6, 4, 6},
		{1, 4, 2},
		{2, 7, 3},
	}
}
