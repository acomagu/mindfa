package mindfa_test

import (
	"fmt"
	"sort"

	"github.com/acomagu/mindfa"
)

// Minimize the DFA as: https://commons.wikimedia.org/wiki/File%3ADFA_to_be_minimized.jpg#/media/File%3ADFA_to_be_minimized.jpg
// The result becomes: https://commons.wikimedia.org/wiki/File%3AMinimized_DFA.jpg#/media/File%3AMinimized_DFA.jpg .
func ExampleMinimize() {
	nState := 6
	nSymbol := 2
	finals := []int{2, 3, 4}
	transitions := [][]int{
		//  0  1
		0: {1, 2}, // a
		1: {0, 3}, // b
		2: {4, 5}, // c
		3: {4, 5}, // d
		4: {4, 5}, // e
		5: {5, 5}, // f
	}
	transitionFunc := func(state, symbol int) int { return transitions[state][symbol] }

	partitions := mindfa.Minimize(nState, nSymbol, finals, transitionFunc)
	fmt.Println(partitions) // Output: [[2 3 4] [0 1] [5]]
}

// This example creates the minimum DFA inputs the digits of year(like 2 -> 0 -> 2 -> 1)
// and accepts if the year is a leap year.
func ExampleMinimize_determiningLeapYear() {
	nSymbol := 10
	nState := 400

	// finals becomes the list of leap years up to 400.
	var finals []int
	for s := 0; s < nState; s++ {
		if s == 0 || (s%100 != 0 && s%4 == 0) {
			finals = append(finals, s)
		}
	}

	// e.g. 2 -> 20 -> 202 -> 21
	transitions := func(state, symbol int) int {
		return (state*10 + symbol) % nState
	}

	partitions := mindfa.Minimize(nState, nSymbol, finals, transitions)

	// Mark years belonging to the same partition with identical number.
	classes := make([]int, nState)
	for _, p := range partitions {
		for _, s := range p {
			classes[s] = p[0]
		}
	}

	checkLeapYear := func(year int) {
		state := classes[0]
		ds := digits(year) // digits(n) returns the slice of the digits of n.
		for i := len(ds) - 1; i >= 0; i-- {
			state = transitions(classes[state], ds[i])
		}

		// If the state is acceptable, it is a leap year.
		if u := sort.SearchInts(finals, state); u < len(finals) && finals[u] == state {
			fmt.Printf("%d is a leap year.\n", year)
		} else {
			fmt.Printf("%d is not a leap year.\n", year)
		}
	}

	checkLeapYear(2019)
	checkLeapYear(2020)
	checkLeapYear(2021)
	// Output:
	// 2019 is not a leap year.
	// 2020 is a leap year.
	// 2021 is not a leap year.
}

// digits returns the slice of the digits of n.
func digits(n int) []int {
	var ans []int
	for n > 0 {
		ans = append(ans, n%10)
		n /= 10
	}
	return ans
}
