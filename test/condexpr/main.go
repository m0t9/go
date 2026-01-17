// run

// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func constants() {
	const (
		constBool bool = 1 < 10
		constStringA string = if constBool { "stringA" } else { "stringB" }
		constStringC string = if false { "stringD" } else { "stringC" }
		constStringD string = if true { if true { "true" } else { "false" } } else { "false" } 
	)

	if constStringA != "stringA" {
		panic("Failed to evaluate ternary for constants in proper way")
	}

	if constStringC != "stringC" {
		panic("Failed to evaluate ternary for constants in proper way")
	}

	if constStringD != "true" {
		panic("Failed to evaluate ternary for constants in proper way")
	}
}

func sideEffects() {
	var (
		cond = false
		thenExpr = false
		elseExpr = false
	)

	reset := func() {
		cond = false
		thenExpr = false
		elseExpr = false
	}

	updateCond := func(v bool) bool {
		fmt.Println("cond")
		cond = true
		return v
	}
	
	updateElse := func() bool {
		fmt.Println("else")
		elseExpr = true
		return true
	}

	updateThen := func() bool {
		fmt.Println("then")
		thenExpr = true
		return true
	}
	
	_ = if updateCond(true) { updateThen() } else { updateElse() }
	if !cond || !thenExpr || elseExpr {
		panic("Something goes wrong with side effects")
	}

	reset()
}

func generic[T any](t T, e T) {
	c := if true { t } else { e }
	fmt.Println(c)
}

func logic() {
	f := func() bool {
		return true
	}

	res := if f() { "A" } else { "B" }
	if res != "A" {
		panic("Logic of ternary does not work")
	}

	res = if !f() { "A" } else { "B" }
	if res != "B" {
		panic("Logic of ternary does not work")
	}
}

func weird() {
	if if true { true } else { false } {
		fmt.Println("Weird")
	}
}


func main() {
	constants()
	sideEffects()
	generic("A", "B")
	logic()
	weird()
}
