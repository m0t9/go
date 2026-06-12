// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tuple

type Of0 struct {
}

func MakeOf0() Of0 {
	return Of0{}
}

func (t Of0) Unpack() Of0 {
	return t
}

type Of1[I1 any] struct {
	I1 I1
}

func MakeOf1[I1 any](i1 I1) Of1[I1] {
	return Of1[I1]{
		I1: i1,
	}
}

func (t Of1[I1]) Unpack() I1 {
	return t.I1
}

type Of2[I1, I2 any] struct {
	I1 I1
	I2 I2
}

func MakeOf2[I1, I2 any](i1 I1, i2 I2) Of2[I1, I2] {
	return Of2[I1, I2]{
		I1: i1,
		I2: i2,
	}
}

func (t Of2[I1, I2]) Unpack() (I1, I2) {
	return t.I1, t.I2
}

type Of3[I1, I2, I3 any] struct {
	I1 I1
	I2 I2
	I3 I3
}

func MakeOf3[I1, I2, I3 any](i1 I1, i2 I2, i3 I3) Of3[I1, I2, I3] {
	return Of3[I1, I2, I3]{
		I1: i1,
		I2: i2,
		I3: i3,
	}
}

func (t Of3[I1, I2, I3]) Unpack() (I1, I2, I3) {
	return t.I1, t.I2, t.I3
}

type Of4[I1, I2, I3, I4 any] struct {
	I1 I1
	I2 I2
	I3 I3
	I4 I4
}

func MakeOf4[I1, I2, I3, I4 any](i1 I1, i2 I2, i3 I3, i4 I4) Of4[I1, I2, I3, I4] {
	return Of4[I1, I2, I3, I4]{
		I1: i1,
		I2: i2,
		I3: i3,
		I4: i4,
	}
}

func (t Of4[I1, I2, I3, I4]) Unpack() (I1, I2, I3, I4) {
	return t.I1, t.I2, t.I3, t.I4
}

type Of5[I1, I2, I3, I4, I5 any] struct {
	I1 I1
	I2 I2
	I3 I3
	I4 I4
	I5 I5
}

func MakeOf5[I1, I2, I3, I4, I5 any](i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) Of5[I1, I2, I3, I4, I5] {
	return Of5[I1, I2, I3, I4, I5]{
		I1: i1,
		I2: i2,
		I3: i3,
		I4: i4,
		I5: i5,
	}
}

func (t Of5[I1, I2, I3, I4, I5]) Unpack() (I1, I2, I3, I4, I5) {
	return t.I1, t.I2, t.I3, t.I4, t.I5
}
