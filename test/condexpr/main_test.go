
package main_test

import "testing"

var (
	sinkInt int
	valA, valB = 100, 200
)

func condition(i int) bool {
	return i % 3 > 1
}

func Bench(b *testing.B) {
	getA := func() int { return 100 }
	getB := func() int { return 200 }

	b.Run("IfAssign", func(sb *testing.B) {
		var res int
		for i := 0; i < sb.N; i++ {
			if condition(i) {
				res = valA
			} else {
				res = valB
			}
		}
		sinkInt = res
	})

	b.Run("CondExprAssign", func(sb *testing.B) {
		var res int
		for i := 0; i < sb.N; i++ {
			res = if condition(i) { valA } else { valB }
		}
		sinkInt = res
	})

	b.Run("IfFuncCall", func(sb *testing.B) {
		var res int
		for i := 0; i < sb.N; i++ {
			if condition(i) {
				res = getA()
			} else {
				res = getB()
			}
		}
		sinkInt = res
	})

	b.Run("CondExprFuncCall", func(sb *testing.B) {
		var res int
		for i := 0; i < sb.N; i++ {
			res = if condition(i) { getA() } else { getB() }
		}
		sinkInt = res
	})
}
