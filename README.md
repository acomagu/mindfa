# mindfa: The implementation of Hopcroft's algorithm in Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/acomagu/mindfa)

The implementation of [DFA minimization](https://en.wikipedia.org/wiki/DFA_minimization) using Hopcroft's algorithm in Go. The time complexity is O(n log n) and the memory complexity is O(n)(n is the number of the states of the input DFA).

## Examples

### 1

The demonstration minimizing the DFA:

> [![DFA to be minimized.jpg](https://upload.wikimedia.org/wikipedia/commons/c/cd/DFA_to_be_minimized.jpg)](https://commons.wikimedia.org/wiki/File:DFA_to_be_minimized.jpg#/media/File:DFA_to_be_minimized.jpg)\
> By Vevek - Own work, created with Inkscape, [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0)

results as:

> [![Minimized DFA.jpg](https://upload.wikimedia.org/wikipedia/commons/6/66/Minimized_DFA.jpg)](https://commons.wikimedia.org/wiki/File:Minimized_DFA.jpg#/media/File:Minimized_DFA.jpg)\
> By Vevek - Own work, created with Inkscape, [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0)

```go
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
fmt.Println(partitions)
```

Output:

```
[[2 3 4] [0 1] [5]]
```

(means (c, d, e), (a, b), (f)).

### 2

This example creates the minimum DFA inputs the digits of year(like 2 -> 0 -> 2 -> 1) and accepts if it's a leap year.

```go
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
	return (state*10 + symbol) % 400
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
```

Output:

```
2019 is not a leap year.
2020 is a leap year.
2021 is not a leap year.
```

## Benchmark

The task is [Example 2](https://github.com/acomagu/mindfa#2) above.

```
goos: linux
goarch: amd64
pkg: github.com/acomagu/mindfa
BenchmarkMinimize-8         1172            933586 ns/op            6869 B/op          8 allocs/op
```
