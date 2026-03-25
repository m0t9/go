package tuple

type T0 struct {
}

func Make0() T0 {
	return T0{}
}

func (t T0) Unpack() T0 {
	return t
}

type T1[I1 any] struct {
	I1 I1
}

func Make1[I1 any](i1 I1) T1[I1] {
	return T1[I1]{
		I1: i1,
	}
}

func (t T1[I1]) Unpack() I1 {
	return t.I1
}

type T2[I1, I2 any] struct {
	I1 I1
	I2 I2
}

func Make2[I1, I2 any](i1 I1, i2 I2) T2[I1, I2] {
	return T2[I1, I2]{
		I1: i1,
		I2: i2,
	}
}

func (t T2[I1, I2]) Unpack() (I1, I2) {
	return t.I1, t.I2
}

type T3[I1, I2, I3 any] struct {
	I1 I1
	I2 I2
	I3 I3
}

func Make3[I1, I2, I3 any](i1 I1, i2 I2, i3 I3) T3[I1, I2, I3] {
	return T3[I1, I2, I3]{
		I1: i1,
		I2: i2,
		I3: i3,
	}
}

func (t T3[I1, I2, I3]) Unpack() (I1, I2, I3) {
	return t.I1, t.I2, t.I3
}
