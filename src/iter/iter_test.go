// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iter_test

import (
	. "iter"
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	nums := []int{2, 3, 4, 5}
	squared := slices.Collect(
		Values(Map(func(idx, x int) (int, int) {
			return idx, x * x
		}, slices.All(nums))),
	)

	for idx, i := range []int{2, 3, 4, 5} {
		wantSquared := i * i
		if gotSquared := squared[idx]; gotSquared != wantSquared {
			t.Errorf("iter.Map gives invalid %d-th item of squared nums. have %d, want %d", idx+1, gotSquared, wantSquared)
		}
	}

	if wantLen, gotLen := 4, len(squared); wantLen != gotLen {
		t.Errorf("iter.Map gives a sequence of invalid length. have %d, want %d", gotLen, wantLen)
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	nums := []int{2, 3, 4, 5}
	onEven := slices.Collect(
		Values(Filter(func(idx, _ int) bool {
			return idx%2 == 0
		}, slices.All(nums))),
	)

	for idx, want := range []int{2, 4} {
		if got := onEven[idx]; got != want {
			t.Errorf("iter.Filter gives invalid %d-th item of squared nums. have %d, want %d", idx+1, got, want)
		}
	}

	if wantLen, gotLen := 2, len(onEven); wantLen != gotLen {
		t.Errorf("iter.Filter gives a sequence of invalid length. have %d, want %d", gotLen, wantLen)
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}
	sum := Reduce(0, func(sum, _, item int) int {
		return sum + item
	}, slices.All(nums))

	if wantSum, gotSum := 15, sum; wantSum != gotSum {
		t.Errorf("iter.Reduce improperly calculates sum. have %d, want %d", gotSum, wantSum)
	}
}

func TestKeys(t *testing.T) {
	t.Parallel()

	s := []int{1, 4, 5, 2, 3}
	keys := slices.Collect(Keys(slices.All(s)))

	if wantLen, gotLen := 5, len(keys); wantLen != gotLen {
		t.Errorf("iter.Keys improperly gives count of keys of iter.Seq2. have %d, want %d", gotLen, wantLen)
	}
	for want, got := range keys {
		if want != got {
			t.Errorf("iter.Keys improperly provides %d-th key. have %d, want %d", want+1, got, want)
		}
	}
}

func TestValues(t *testing.T) {
	t.Parallel()

	s := []int{1, 4, 5, 2, 3}
	values := slices.Collect(Values(slices.All(s)))

	if wantLen, gotLen := 5, len(values); wantLen != gotLen {
		t.Errorf("iter.Values improperly gives count of values of iter.Seq2. have %d, want %d", gotLen, wantLen)
	}

	for idx, got := range values {
		if want := s[idx]; want != got {
			t.Errorf("iter.Values improperly provides %d-th value. have %d, want %d", idx+1, got, want)
		}
	}
}
