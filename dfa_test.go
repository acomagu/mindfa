package mindfa

import (
	"testing"

	"github.com/matryer/is"
)

func TestMinimize2(t *testing.T) {
	is := is.New(t)

	nState := 5
	nSymbol := 3
	finals := []int{1, 3}
	var transitions [5][3]int
	transitions[0][0] = 3
	transitions[0][1] = 4
	transitions[0][2] = 3
	transitions[1][0] = 4
	transitions[1][1] = 2
	transitions[1][2] = 3
	transitions[2][0] = 3
	transitions[2][1] = 4
	transitions[2][2] = 3
	transitions[3][0] = 2
	transitions[3][1] = 4
	transitions[3][2] = 4
	transitions[4][0] = 0
	transitions[4][1] = 3
	transitions[4][2] = 3

	partitions := Minimize(nState, nSymbol, finals, func(s, c int) int { return transitions[s][c] })
	is.Equal(len(partitions), 4)
}

func TestMinimize1(t *testing.T) {
	is := is.New(t)

	nSymbol := 10
	nState := 400

	var finals []int
	for s := 0; s < nState; s++ {
		if s == 0 || (s%100 != 0 && s%4 == 0) {
			finals = append(finals, s)
		}
	}

	transitions := make([][]int, 0, nState)
	for i := 0; i < nState; i++ {
		t := make([]int, 0, nSymbol)
		for c := 0; c < nSymbol; c++ {
			t = append(t, (i*nSymbol+c)%nState)
		}

		transitions = append(transitions, t)
	}

	partitions := Minimize(nState, nSymbol, finals, func(s, c int) int { return transitions[s][c] })
	is.Equal(len(partitions), 7)

	classes := make([]int, nState)
	for _, p := range partitions {
		for _, s := range p {
			classes[s] = p[0]
		}
	}

	for n := 0; n < 5000; n++ {
		cur := classes[0]
		ds := digits(n)
		for i := len(ds) - 1; i >= 0; i-- {
			cur = transitions[classes[cur]][ds[i]]
		}
		is.Equal(includes(finals, cur), n%4 == 0 && (n%100 != 0 || n%400 == 0))
	}
}

var _unused interface{}

func BenchmarkMinimize(b *testing.B) {
	nSymbol := 10
	nState := 400

	var finals []int
	for s := 0; s < nState; s++ {
		if s == 0 || (s%100 != 0 && s%4 == 0) {
			finals = append(finals, s)
		}
	}

	transitions := make([][]int, 0, nState)
	for i := 0; i < nState; i++ {
		t := make([]int, 0, nSymbol)
		for c := 0; c < nSymbol; c++ {
			t = append(t, (i*nSymbol+c)%nState)
		}

		transitions = append(transitions, t)
	}

	transition := func(s, c int) int { return transitions[s][c] }
	for i := 0; i < b.N; i++ {
		_unused = Minimize(nState, nSymbol, finals, transition)
	}
}

func digits(n int) []int {
	var ans []int
	for n > 0 {
		ans = append(ans, n%10)
		n /= 10
	}
	return ans
}
