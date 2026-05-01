// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func membersAccess() {
	tp := (1, 2, 3)
	_, _, _ = tp.I1, tp.I2, tp.I3
}

func unpacking() {
	tp := ("string", 2, true)

	// In assignment.
	_, _, _ = tp...

	// Passing as function arguments.
	f := func(string, int, bool) {

	}
	f(tp...)
}

func sliceVariadicUnpacking() {
	sum := func(values ...int) int {
		res := 0
		for _, v := range values {
			res += v
		}
		return res
	}

	values := []int{1, 2, 3}
	if sum(values...) != 6 {
		panic("slice variadic unpacking does not work")
	}
}

func passingToStructures() {
	f := func() (int, error) {
		return (42, error(nil))
	}

	ch := make(chan (int, error), 1)
	var sl [](int, error)

	sl = append(sl, f())
	ch <- f()

	item := <- ch
	if x, err := item...; err == nil {
		fmt.Println("everything works", x)
	}
}

func basicPatternMatching() {
	v := (42, "s")

	switch v {
	case (42, "s"):
		fmt.Println("matched")
	default:
		fmt.Println("not matched")
	}
}

func withGenerics[T any, V any](t T, v V) {
	var tp (T, V)
	tp.I1 = t
	tp.I2 = v
	_ = tp
}

func main() {
	membersAccess()
	unpacking()
	sliceVariadicUnpacking()
	passingToStructures()
	basicPatternMatching()
	withGenerics("string", 12)
}
