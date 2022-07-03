// Copyright 2022 The AC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/itsubaki/q"
)

func main() {
	qsim := q.New()

	// https://en.wikipedia.org/wiki/Algorithmic_cooling#Reversible_algorithmic_cooling_-_basic_compression_subroutine
	a := qsim.Zero()
	b := qsim.Zero()
	c := qsim.Zero()
	qsim.H(a)
	qsim.H(b)
	qsim.H(c)

	qsim.CNOT(c, b)
	// https://en.wikipedia.org/wiki/Fredkin_gate
	qsim.CNOT(a, c)
	s := qsim.Zero()
	qsim.CCNOT(b, c, s)
	qsim.CNOT(a, s)
	qsim.CNOT(c, s)

	for _, s := range qsim.State() {
		fmt.Println(s)
	}
}
