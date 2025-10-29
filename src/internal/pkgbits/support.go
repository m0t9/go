// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

import (
	"fmt"
	"runtime/debug"
)

func assert(b bool) {
	if !b {
		panic(fmt.Sprintf("assertion failed: %s", debug.Stack()))
	}
}

func panicf(format string, args ...any) {
	panic(fmt.Errorf(format, args...))
}
