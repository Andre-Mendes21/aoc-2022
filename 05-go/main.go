package main

import (
	"fmt"
	"strings"
)

func makeState(states []string) []Stack {
	var result []Stack
	for _, state := range states {
		var column Stack
		for _, rn := range state {
			column.push(string(rn))
		}
		result = append(result, column)
	}
	return result
}

func applyMoves(states []Stack, moves [][3]int) {
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			rn, ok := states[move[1]-1].pop()
			if ok {
				states[move[2]-1].push(rn)
			}
		}
	}
}

func applyMoves2(states []Stack, moves [][3]int) {
	for _, move := range moves {
		count := move[0]
		aux := Stack{}
		for i := 0; i < count; i++ {
			rn, ok := states[move[1]-1].pop()
			if ok {
				aux.push(rn)
			}
		}
		for !aux.isEmpty() {
			rn, ok := aux.pop()
			if ok {
				states[move[2]-1].push(rn)
			}
		}
	}
}

func solveBoth(initState []string, moves [][3]int, apply func([]Stack, [][3]int)) string {
	state := makeState(initState)
	apply(state, moves)
	var result = []string{}
	for _, st := range state {
		result = append(result, st.peek())
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Printf("Part One Sample: %s\n", solveBoth(SampleState(), SampleMoves(), applyMoves))
	fmt.Printf("Part One Test: %s\n", solveBoth(TestState(), TestMoves(), applyMoves))

	fmt.Printf("\nPart Two Sample: %s\n", solveBoth(SampleState(), SampleMoves(), applyMoves2))
	fmt.Printf("Part Two Test: %s\n", solveBoth(TestState(), TestMoves(), applyMoves2))
}
