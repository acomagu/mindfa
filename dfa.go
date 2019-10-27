// Package mindfa implements DFA minimization using Hopcroft's algorithm.
package mindfa

import (
	"fmt"
	"sort"
)

// Minimize takes a DFA representation, minimize it and returns the groups of
// the states. The states in a group has same behavior.
//
// The all of arguments represents a single DFA(before minimization). It has
// 0..nState states, 0..nSymbol input symbols and transitions as transition function.
//
// transition is a function takes state and symbol and returns the destination
// state. The states which is included in finals are accepting state. The order
// of the resulting partitions is not defined, but the order of the states in
// a partition is ascending.
//
// It uses Hopcroft's algorithm. The memory complexity is O(nState).
//
// The all numbers in finals must be in [0, nState), and two same values must not
// appear.
func Minimize(nState, nSymbol int, finals []int, transition func(state, symbol int) int) [][]int {
	if len(finals) > nState {
		panic(fmt.Sprintf("len(finals) should be less than or equal to nState: len(finals) = %d, nState = %d", len(finals), nState))
	}

	whole := make([]int, nState+max(len(finals), nState-len(finals)))
	copy(whole, finals)
	sort.Ints(whole[:len(finals)])

	for i := 0; i < len(finals)-1; i++ {
		if whole[i] == whole[i+1] {
			panic(fmt.Sprintf("finals contains same two value: %d", whole[i]))
		}
	}

	cmpl(whole[len(finals):nState], whole[:len(finals)], nState)

	buf := whole[nState:]

	partitions := [][]int{whole[:len(finals)], whole[len(finals):nState]}
	// works is a set of the partition which has never tried to be split.
	works := [][]int{whole[:len(finals)], whole[len(finals):nState]}

	for len(works) > 0 {
		for c := 0; c < nSymbol; c++ {
			for ip, pFrom := range partitions {
				ip1, ip2 := 0, len(buf)-1
				for _, state := range pFrom {
					if includes(works[0], transition(state, c)) {
						buf[ip1] = state
						ip1++
					} else {
						buf[ip2] = state
						ip2--
					}
				}

				if ip1 == 0 || ip2 == len(buf)-1 {
					continue
				}

				p1 := pFrom[:ip1]
				copy(p1, buf[:ip1])

				p2 := pFrom[ip1:]
				for i := range p2 {
					p2[i] = buf[len(buf)-1-i]
				}

				var split bool
				for i, w := range works {
					if &w[0] != &pFrom[0] {
						continue
					}

					// Split works[i].
					works = append(works, p2)
					works[i] = p1
					split = true
					break
				}

				if !split {
					if len(p1) < len(p2) {
						works = append(works, p1)
					} else {
						works = append(works, p2)
					}
				}
				partitions[ip] = p1
				partitions = append(partitions, p2) // Don't worry, p2 is not iterated in the current loop.
			}
		}
		// pseudo-shift
		works[0] = works[len(works)-1]
		works = works[:len(works)-1]
	}
	return partitions
}

// cmpl returns the complement set of a in (0..upper).
func cmpl(dst, a []int, upper int) {
	var n, i int
	for _, u := range a {
		for ; n < u; n++ {
			dst[i] = n
			i++
		}
		n++
	}
	for ; n < upper; n++ {
		dst[i] = n
		i++
	}
}

func includes(a []int, e int) bool {
	if len(a) < 100 {
		var i int
		for ; i < len(a) && a[i] < e; i++ {
		}
		return i < len(a) && a[i] == e
	}
	i := sort.SearchInts(a, e)
	return i < len(a) && a[i] == e
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
