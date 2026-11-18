package fibonacci

import (
	"math"
)

func ClosedForm(n int) int {
	sqrt_5 := math.Sqrt(5)
	fn := float64(n)
	return int((1/sqrt_5)*math.Pow(((1+sqrt_5)/2), fn) -
		(1/sqrt_5)*math.Pow(((1-sqrt_5)/2), fn))
}

func Recursive(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Recursive(n-1) + Recursive(n-2)
	}
}

func Iterative(n int) int {
	a := 0
	b := 1
	for range n {
		acc := a + b
		a = b
		b = acc
	}
	return b
}

func MemoizedRecurssive(n int) int {
	MEMO := make(map[int]int)
	MEMO[0] = 0
	MEMO[1] = 1
	var recurse func(int) int
	recurse = func(n int) int {
		if fib, ok := MEMO[n]; ok {
			return fib
		} else {
			fib := recurse(n-1) + recurse(n-2)
			MEMO[n] = fib
			return fib
		}
	}
	return recurse(n)
}
