// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_quota, _rem         int
	_parsed, _err        any
	_swapped1, _swapped2 string
	_min, _max, _sum     int
)

func divmod(a, b int) (int, int) {
	return (a / b, a % b)
}

func parseAndDouble(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return (0, err)
	}
	return (i * 2, error(nil))
}

func swap(x, y string) (string, string) {
	return (y, x)
}

func stats(nums []int) (int, int, int) {
	if len(nums) == 0 {
		return (0, 0, 0)
	}
	min, max, sum := nums[0], nums[0], 0
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	return (min, max, sum)
}

func benchmark(name string, iterations int, fn func()) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	elapsed := time.Since(start)
	nsPerOp := elapsed.Nanoseconds() / int64(iterations)
	fmt.Printf("%-20s %10d runs   %12s total   %8d ns/op\n",
		name, iterations, elapsed, nsPerOp)
}

func main() {
	const iterations = 10_000_000

	fmt.Printf("Benchmarking functions with tuples\n")
	fmt.Printf("%-20s %-12s %-18s %s\n", "Function", "Iterations", "Total Time", "ns/op")
	fmt.Println("----------------------------------------------------------------")

	benchmark("divmod", iterations, func() {
		_quota, _rem = divmod(123456, 789)...
	})

	benchmark("parseAndDouble", iterations, func() {
		var val int
		var err error
		val, err = parseAndDouble("12345")...
		// Store to globals to avoid compiler elimination.
		_parsed, _err = val, err
	})

	benchmark("swap", iterations, func() {
		_swapped1, _swapped2 = swap("hello", "world")...
	})

	data := []int{5, 2, 9, 1, 7, 3, 8, 4, 6, 0}
	benchmark("stats", iterations, func() {
		_min, _max, _sum = stats(data)...
	})
}
