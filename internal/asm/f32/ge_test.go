// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package f32

import (
	"fmt"
	"math"
	"testing"
)

var gerTests = []struct {
	name string
	a    []float32
	m    uintptr
	n    uintptr
	x    []float32
	y    []float32
	incX uintptr
	incY uintptr

	trueAns []float32
}{
	{
		name:    "Unit",
		m:       1,
		n:       1,
		a:       []float32{10},
		x:       []float32{2},
		y:       []float32{4.4},
		incX:    1,
		incY:    1,
		trueAns: []float32{18.8},
	},
	{
		name: "M gt N inc 1",
		m:    5,
		n:    3,
		a: []float32{
			1.3, 2.4, 3.5,
			2.6, 2.8, 3.3,
			-1.3, -4.3, -9.7,
			8, 9, -10,
			-12, -14, -6,
		},
		x:       []float32{-2, -3, 0, 1, 2},
		y:       []float32{-1.1, 5, 0},
		incX:    1,
		incY:    1,
		trueAns: []float32{3.5, -7.6, 3.5, 5.9, -12.2, 3.3, -1.3, -4.3, -9.7, 6.9, 14, -10, -14.2, -4, -6},
	},
	{
		name: "M eq N inc 1",
		m:    3,
		n:    3,
		a: []float32{
			1.3, 2.4, 3.5,
			2.6, 2.8, 3.3,
			-1.3, -4.3, -9.7,
		},
		x:       []float32{-2, -3, 0},
		y:       []float32{-1.1, 5, 0},
		incX:    1,
		incY:    1,
		trueAns: []float32{3.5, -7.6, 3.5, 5.9, -12.2, 3.3, -1.3, -4.3, -9.7},
	},
	{
		name: "M lt N inc 1",
		m:    3,
		n:    6,
		a: []float32{
			1.3, 2.4, 3.5, 4.8, 1.11, -9,
			2.6, 2.8, 3.3, -3.4, 6.2, -8.7,
			-1.3, -4.3, -9.7, -3.1, 8.9, 8.9,
		},
		x:       []float32{-2, -3, 0},
		y:       []float32{-1.1, 5, 0, 9, 19, 22},
		incX:    1,
		incY:    1,
		trueAns: []float32{3.5, -7.6, 3.5, -13.2, -36.89, -53, 5.9, -12.2, 3.3, -30.4, -50.8, -74.7, -1.3, -4.3, -9.7, -3.1, 8.9, 8.9},
	},
	{
		name: "M gt N inc not 1",
		m:    5,
		n:    3,
		a: []float32{
			1.3, 2.4, 3.5,
			2.6, 2.8, 3.3,
			-1.3, -4.3, -9.7,
			8, 9, -10,
			-12, -14, -6,
		},
		x:       []float32{-2, -3, 0, 1, 2, 6, 0, 9, 7},
		y:       []float32{-1.1, 5, 0, 8, 7, -5, 7},
		incX:    2,
		incY:    3,
		trueAns: []float32{3.5, -13.6, -10.5, 2.6, 2.8, 3.3, -3.5, 11.7, 4.3, 8, 9, -10, -19.700000000000003, 42, 43},
	},
	{
		name: "M eq N inc not 1",
		m:    3,
		n:    3,
		a: []float32{
			1.3, 2.4, 3.5,
			2.6, 2.8, 3.3,
			-1.3, -4.3, -9.7,
		},
		x:       []float32{-2, -3, 0, 8, 7, -9, 7, -6, 12, 6, 6, 6, -11},
		y:       []float32{-1.1, 5, 0, 0, 9, 8, 6},
		incX:    4,
		incY:    3,
		trueAns: []float32{3.5, 2.4, -8.5, -5.1, 2.8, 45.3, -14.5, -4.3, 62.3},
	},
	{
		name:    "M lt N inc not 1",
		m:       3,
		n:       6,
		a:       []float32{1.3, 2.4, 3.5, 4.8, 1.11, -9, 2.6, 2.8, 3.3, -3.4, 6.2, -8.7, -1.3, -4.3, -9.7, -3.1, 8.9, 8.9},
		x:       []float32{-2, -3, 0, 0, 8, 0, 9, -3},
		y:       []float32{-1.1, 5, 0, 9, 19, 22, 11, -8.11, -9.22, 9.87, 7},
		incX:    3,
		incY:    2,
		trueAns: []float32{3.5, 2.4, -34.5, -17.2, 19.55, -23, 2.6, 2.8, 3.3, -3.4, 6.2, -8.7, -11.2, -4.3, 161.3, 95.9, -74.08, 71.9},
	},
	{
		name:    "Y NaN element",
		m:       1,
		n:       1,
		a:       []float32{1.3},
		x:       []float32{1.3},
		y:       []float32{float32(math.NaN())},
		incX:    1,
		incY:    1,
		trueAns: []float32{float32(math.NaN())},
	},
}

func TestGer(t *testing.T) {
	// Ger(test.m, test.n, test.alpha, test.x, test.incX, test.y, test.incY, test.a, test.lda)
	const (
		xGdVal, yGdVal, aGdVal = -0.5, 1.5, -1
		gdLn                   = 4
	)

	for i, test := range gerTests {
		for _, align := range align2 {
			prefix := fmt.Sprintf("Test %v (x:%v y:%v a:%v)", i, align.x, align.y, align.x^align.y)
			// xg, yg := guardIncVector(test.x, xGdVal, inc.x, gdLn)
			xg, yg := guardVector(test.x, xGdVal, align.x+gdLn), guardVector(test.y, yGdVal, align.y+gdLn)
			x, y := xg[align.x+gdLn:len(xg)-(align.x+gdLn)], yg[align.y+gdLn:len(yg)-(align.y+gdLn)]
			ag := guardVector(test.a, aGdVal, align.x^align.y+gdLn)
			a := ag[(align.x^align.y)+gdLn : len(ag)-(align.x^align.y+gdLn)]

			// Test with row major
			var alpha float32 = 1.0
			Ger(test.m, test.n, alpha, x, test.incX, y, test.incY, a, test.n)
			for i := range test.trueAns {
				if !within(a[i], test.trueAns[i]) {
					t.Errorf(msgVal, prefix, i, a[i], test.trueAns[i])
				}
			}

			if !isValidGuard(xg, xGdVal, gdLn) {
				t.Errorf(msgGuard, prefix, "x", xg[:gdLn], xg[len(xg)-gdLn:])
			}
			if !isValidGuard(yg, yGdVal, gdLn) {
				t.Errorf(msgGuard, prefix, "y", yg[:gdLn], yg[len(yg)-gdLn:])
			}
			if !isValidGuard(ag, aGdVal, gdLn) {
				t.Errorf(msgGuard, prefix, "a", ag[:gdLn], ag[len(ag)-gdLn:])
			}

			if !equalStrided(test.x, x, 1) {
				t.Errorf("%v: modified read-only x argument", prefix)
			}
			if !equalStrided(test.y, y, 1) {
				t.Errorf("%v: modified read-only y argument", prefix)
			}
		}
	}
}

/* type sgerWrap struct{}

func (d sgerWrap) Sger(m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	Ger(uintptr(m), uintptr(n), alpha, x, uintptr(incX), y, uintptr(incY), a, uintptr(lda))
}

func TestBlasGer(t *testing.T) {
	testblas.SgerTest(t, sgerWrap{})
}

func BenchmarkBlasGer(t *testing.B) {
	for _, dims := range newIncSet(3, 10, 30, 100, 300, 1000, 1e4, 1e5) {
		m, n := dims.x, dims.y
		if m/n >= 100 || n/m >= 100 || (m == 1e5 && n == 1e5) {
			continue
		}
		for _, inc := range newIncSet(1, 2, 3, 4, 10) {
			incX, incY := inc.x, inc.y
			t.Run(fmt.Sprintf("Sger %dx%d (%d %d)", m, n, incX, incY), func(b *testing.B) {
				for i := 0; i < t.N; i++ {
					testblas.SgerBenchmark(b, sgerWrap{}, m, n, incX, incY)
				}
			})

		}
	}
}
*/
