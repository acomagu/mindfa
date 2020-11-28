# mindfa: The implementation of Hopcroft's algorithm in Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/acomagu/mindfa)

The implementation of [DFA minimization](https://en.wikipedia.org/wiki/DFA_minimization) using Hopcroft's algorithm in Go. The time complexity is O(n log n) and the memory complexity is O(n)(n is the number of the states of the input DFA).

## Example Usage

Consider minimizing this DFA:

> [![DFA to be minimized.jpg](https://upload.wikimedia.org/wikipedia/commons/c/cd/DFA_to_be_minimized.jpg)](https://commons.wikimedia.org/wiki/File:DFA_to_be_minimized.jpg#/media/File:DFA_to_be_minimized.jpg)\
> By Vevek - Own work, created with Inkscape, [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0)

It results as:

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

[Playground](https://play.golang.org/p/cwMl0qKIAp3)

## Benchmark

[BenchmarkMinimize](https://github.com/acomagu/mindfa/blob/master/dfa_test.go#LC81:~:text=func%20BenchmarkMinimize(b%20*testing.B)%20%7B)

```
goos: linux
goarch: amd64
pkg: github.com/acomagu/mindfa
BenchmarkMinimize-8         1172            933586 ns/op            6869 B/op          8 allocs/op
```
