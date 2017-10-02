// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package f64

import (
	"fmt"
	"testing"
)

type DgemvCase struct {
	Name  string
	m     int
	n     int
	A     []float64
	trans bool
	x     []float64
	y     []float64
	xCopy []float64
	yCopy []float64

	NoTrans []DgemvSubcase
	Trans   []DgemvSubcase
}

type DgemvSubcase struct {
	mulXNeg1  bool
	mulYNeg1  bool
	alpha     float64
	beta      float64
	want      []float64
	wantRevX  []float64
	wantRevY  []float64
	wantRevXY []float64
}

var DgemvCases = []DgemvCase{
	{ // 3x3
		Name:  "M_eq_N",
		trans: false,
		m:     3,
		n:     3,
		A: []float64{
			4.1, 6.2, 8.1,
			9.6, 3.5, 9.1,
			10, 7, 3,
		},
		x: []float64{1, 2, 3},
		y: []float64{7, 2, 2},

		NoTrans: []DgemvSubcase{ // (2x2, 2x1, 1x2, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0},
				wantRevX:  []float64{0, 0, 0},
				wantRevY:  []float64{0, 0, 0},
				wantRevXY: []float64{0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{7, 2, 2},
				wantRevX:  []float64{7, 2, 2},
				wantRevY:  []float64{7, 2, 2},
				wantRevXY: []float64{7, 2, 2},
			},
			{alpha: 1, beta: 0,
				want:      []float64{40.8, 43.9, 33},
				wantRevX:  []float64{32.8, 44.9, 47},
				wantRevY:  []float64{33, 43.9, 40.8},
				wantRevXY: []float64{47, 44.9, 32.8},
			},
			{alpha: 8, beta: -6,
				want:      []float64{284.4, 339.2, 252},
				wantRevX:  []float64{220.4, 347.2, 364},
				wantRevY:  []float64{222, 339.2, 314.4},
				wantRevXY: []float64{334, 347.2, 250.4},
			},
		},

		Trans: []DgemvSubcase{ // (2x2, 1x2, 2x1, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0},
				wantRevX:  []float64{0, 0, 0},
				wantRevY:  []float64{0, 0, 0},
				wantRevXY: []float64{0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{1, 2, 3},
				wantRevX:  []float64{1, 2, 3},
				wantRevY:  []float64{1, 2, 3},
				wantRevXY: []float64{1, 2, 3},
			},
			{alpha: 1, beta: 0,
				want:      []float64{67.9, 64.4, 80.9},
				wantRevX:  []float64{97.4, 68.4, 55.4},
				wantRevY:  []float64{80.9, 64.4, 67.9},
				wantRevXY: []float64{55.4, 68.4, 97.4},
			},
			{alpha: 8, beta: -6,
				want:      []float64{537.2, 503.2, 629.2},
				wantRevX:  []float64{773.2, 535.2, 425.2},
				wantRevY:  []float64{641.2, 503.2, 525.2},
				wantRevXY: []float64{437.2, 535.2, 761.2},
			},
		},
	},

	{ // 5x3
		Name: "M_gt_N",
		m:    5,
		n:    3,
		A: []float64{
			4.1, 6.2, 8.1,
			9.6, 3.5, 9.1,
			10, 7, 3,
			1, 1, 2,
			9, 2, 5,
		},
		x: []float64{1, 2, 3},
		y: []float64{7, 8, 9, 10, 11},

		NoTrans: []DgemvSubcase{ //(4x2, 4x1, 1x2, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0, 0, 0},
				wantRevX:  []float64{0, 0, 0, 0, 0},
				wantRevY:  []float64{0, 0, 0, 0, 0},
				wantRevXY: []float64{0, 0, 0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{7, 8, 9, 10, 11},
				wantRevX:  []float64{7, 8, 9, 10, 11},
				wantRevY:  []float64{7, 8, 9, 10, 11},
				wantRevXY: []float64{7, 8, 9, 10, 11},
			},
			{alpha: 1, beta: 0,
				want:      []float64{40.8, 43.9, 33, 9, 28},
				wantRevX:  []float64{32.8, 44.9, 47, 7, 36},
				wantRevY:  []float64{28, 9, 33, 43.9, 40.8},
				wantRevXY: []float64{36, 7, 47, 44.9, 32.8},
			},
			{alpha: 8, beta: -6,
				want:      []float64{284.4, 303.2, 210, 12, 158},
				wantRevX:  []float64{220.4, 311.2, 322, -4, 222},
				wantRevY:  []float64{182, 24, 210, 291.2, 260.4},
				wantRevXY: []float64{246, 8, 322, 299.2, 196.4},
			},
		},

		Trans: []DgemvSubcase{ //( 2x4, 1x4, 2x1, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0},
				wantRevX:  []float64{0, 0, 0},
				wantRevY:  []float64{0, 0, 0},
				wantRevXY: []float64{0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{1, 2, 3},
				wantRevX:  []float64{1, 2, 3},
				wantRevY:  []float64{1, 2, 3},
				wantRevXY: []float64{1, 2, 3},
			},
			{alpha: 1, beta: 0,
				want:      []float64{304.5, 166.4, 231.5},
				wantRevX:  []float64{302.1, 188.2, 258.1},
				wantRevY:  []float64{231.5, 166.4, 304.5},
				wantRevXY: []float64{258.1, 188.2, 302.1},
			},
			{alpha: 8, beta: -6,
				want:      []float64{2430, 1319.2, 1834},
				wantRevX:  []float64{2410.8, 1493.6, 2046.8},
				wantRevY:  []float64{1846, 1319.2, 2418},
				wantRevXY: []float64{2058.8, 1493.6, 2398.8},
			},
		},
	},

	{ // 3x5
		Name:  "M_lt_N",
		trans: false,
		m:     3,
		n:     5,
		A: []float64{
			4.1, 6.2, 8.1, 10, 7,
			9.6, 3.5, 9.1, -2, 9,
			10, 7, 3, 1, -5,
		},
		x: []float64{1, 2, 3, -7.6, 8.1},
		y: []float64{7, 2, 2},

		NoTrans: []DgemvSubcase{ // (2x4, 2x1, 1x4, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0},
				wantRevX:  []float64{0, 0, 0},
				wantRevY:  []float64{0, 0, 0},
				wantRevXY: []float64{0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{7, 2, 2},
				wantRevX:  []float64{7, 2, 2},
				wantRevY:  []float64{7, 2, 2},
				wantRevXY: []float64{7, 2, 2},
			},
			{alpha: 1, beta: 0,
				want:      []float64{21.5, 132, -15.1},
				wantRevX:  []float64{37.39, 83.46, 33.8},
				wantRevY:  []float64{-15.1, 132, 21.5},
				wantRevXY: []float64{33.8, 83.46, 37.39},
			},
			{alpha: 8, beta: -6,
				want:      []float64{130, 1044, -132.8},
				wantRevX:  []float64{257.12, 655.68, 258.4},
				wantRevY:  []float64{-162.8, 1044, 160.},
				wantRevXY: []float64{228.4, 655.68, 287.12},
			},
		},

		Trans: []DgemvSubcase{ // (4x2, 1x2, 4x1, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0, 0, 0},
				wantRevX:  []float64{0, 0, 0, 0, 0},
				wantRevY:  []float64{0, 0, 0, 0, 0},
				wantRevXY: []float64{0, 0, 0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{1, 2, 3, -7.6, 8.1},
				wantRevX:  []float64{1, 2, 3, -7.6, 8.1},
				wantRevY:  []float64{1, 2, 3, -7.6, 8.1},
				wantRevXY: []float64{1, 2, 3, -7.6, 8.1},
			},
			{alpha: 1, beta: 0,
				want:      []float64{67.9, 64.4, 80.9, 68, 57},
				wantRevX:  []float64{97.4, 68.4, 55.4, 23, -3},
				wantRevY:  []float64{57, 68, 80.9, 64.4, 67.9},
				wantRevXY: []float64{-3, 23, 55.4, 68.4, 97.4},
			},
			{alpha: 8, beta: -6,
				want:      []float64{537.2, 503.2, 629.2, 589.6, 407.4},
				wantRevX:  []float64{773.2, 535.2, 425.2, 229.6, -72.6},
				wantRevY:  []float64{450, 532, 629.2, 560.8, 494.6},
				wantRevXY: []float64{-30, 172, 425.2, 592.8, 730.6},
			},
		},
	},

	{ // 7x7
		Name:  "M_eq_N_Lg",
		trans: true,
		m:     7,
		n:     7,
		A: []float64{
			4.1, 6.2, 8.1, 2.5, 3.3, 7.4, 9.3,
			9.6, 3.5, 9.1, 1.2, 5.4, 4.8, 8.7,
			10, 7, 3, 2, 4, 1, 12,
			9.6, 3.5, 9.1, 1.2, 5.4, 4.8, 8.7,
			4.1, 6.2, 8.1, 2.5, 3.3, 7.4, 9.3,
			10, 7, 3, 2, 4, 1, 12,
			9.6, 3.5, 9.1, 1.2, 5.4, 4.8, 8.7,
		},
		x: []float64{1, 2, 3, 4, 5, 6, 7},
		y: []float64{7, 8, 9, 10, 11, 12, 13},

		NoTrans: []DgemvSubcase{ // (4x4, 4x2, 4x1, 2x4, 2x2, 2x1, 1x4, 1x2, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevX:  []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevY:  []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevXY: []float64{0, 0, 0, 0, 0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{7, 8, 9, 10, 11, 12, 13},
				wantRevX:  []float64{7, 8, 9, 10, 11, 12, 13},
				wantRevY:  []float64{7, 8, 9, 10, 11, 12, 13},
				wantRevXY: []float64{7, 8, 9, 10, 11, 12, 13},
			},
			{alpha: 1, beta: 0,
				want:      []float64{176.8, 165.4, 151, 165.4, 176.8, 151, 165.4},
				wantRevX:  []float64{150.4, 173, 161, 173, 150.4, 161, 173},
				wantRevY:  []float64{165.4, 151, 176.8, 165.4, 151, 165.4, 176.8},
				wantRevXY: []float64{173, 161, 150.4, 173, 161, 173, 150.4},
			},
			{alpha: 8, beta: -6,
				want:      []float64{1372.4, 1275.2, 1154, 1263.2, 1348.4, 1136, 1245.2},
				wantRevX:  []float64{1161.2, 1336, 1234, 1324, 1137.2, 1216, 1306},
				wantRevY:  []float64{1281.2, 1160, 1360.4, 1263.2, 1142, 1251.2, 1336.4},
				wantRevXY: []float64{1342, 1240, 1149.2, 1324, 1222, 1312, 1125.2},
			},
		},

		Trans: []DgemvSubcase{ // (4x4, 2x4, 1x4, 4x2, 2x2, 1x2, 4x1, 2x1, 1x1)
			{alpha: 0, beta: 0,
				want:      []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevX:  []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevY:  []float64{0, 0, 0, 0, 0, 0, 0},
				wantRevXY: []float64{0, 0, 0, 0, 0, 0, 0},
			},
			{alpha: 0, beta: 1,
				want:      []float64{1, 2, 3, 4, 5, 6, 7},
				wantRevX:  []float64{1, 2, 3, 4, 5, 6, 7},
				wantRevY:  []float64{1, 2, 3, 4, 5, 6, 7},
				wantRevXY: []float64{1, 2, 3, 4, 5, 6, 7},
			},
			{alpha: 1, beta: 0,
				want:      []float64{581.4, 367.1, 490.9, 124.2, 310.8, 303, 689.1},
				wantRevX:  []float64{558.6, 370.9, 499.1, 127.8, 305.2, 321, 684.9},
				wantRevY:  []float64{689.1, 303, 310.8, 124.2, 490.9, 367.1, 581.4},
				wantRevXY: []float64{684.9, 321., 305.2, 127.8, 499.1, 370.9, 558.6},
			},
			{alpha: 8, beta: -6,
				want:      []float64{4645.2, 2924.8, 3909.2, 969.6, 2456.4, 2388, 5470.8},
				wantRevX:  []float64{4462.8, 2955.2, 3974.8, 998.4, 2411.6, 2532, 5437.2},
				wantRevY:  []float64{5506.8, 2412, 2468.4, 969.6, 3897.2, 2900.8, 4609.2},
				wantRevXY: []float64{5473.2, 2556., 2423.6, 998.4, 3962.8, 2931.2, 4426.8},
			},
		},
	},
}

func TestGemv(t *testing.T) {
	for _, test := range DgemvCases {
		t.Run(fmt.Sprintf("(%vx%v)", test.m, test.n), func(tt *testing.T) {
			for i, cas := range test.NoTrans {
				tt.Run(fmt.Sprintf("NoTrans case %v", i), func(st *testing.T) {
					// Test that it passes with row-major
					dgemvcomp(st, test, false, cas, i)
				})
			}
			for i, cas := range test.Trans {
				tt.Run(fmt.Sprintf("Trans case %v", i), func(st *testing.T) {
					// Test that it passes with row-major
					dgemvcomp(st, test, true, cas, i)
				})
			}
		})
	}
}

func dgemvcomp(t *testing.T, test DgemvCase, trans bool, cas DgemvSubcase, i int) {
	const (
		xGdVal, yGdVal, aGdVal = 0.5, 1.5, 10
		gdLn                   = 4
	)
	if trans {
		test.x, test.y = test.y, test.x
	}
	GemvT2 := GemvT
	prefix := fmt.Sprintf("%s - (%vx%v) t:%v (a:%v,b:%v)", test.Name, test.m, test.n, trans, cas.alpha, cas.beta)
	xg, yg := guardVector(test.x, xGdVal, gdLn), guardVector(test.y, yGdVal, gdLn)
	x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
	ag := guardVector(test.A, aGdVal, gdLn)
	a := ag[gdLn : len(ag)-gdLn]

	lda := uintptr(test.n)
	if trans {
		GemvT2(uintptr(test.m), uintptr(test.n), cas.alpha, a, lda, x, 1, cas.beta, y, 1)
	} else {
		GemvN(uintptr(test.m), uintptr(test.n), cas.alpha, a, lda, x, 1, cas.beta, y, 1)
	}
	for i := range cas.want {
		if !within(y[i], cas.want[i]) {
			t.Errorf(msgVal, prefix, i, y[i], cas.want[i])

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
		t.Errorf(msgReadOnly, prefix, "x")
	}
	if !equalStrided(test.A, a, 1) {
		t.Errorf(msgReadOnly, prefix, "a")
	}

	for _, inc := range newIncSet(-1, 1, 2, 3) {
		incPrefix := fmt.Sprintf("%s inc(x:%v, y:%v)", prefix, inc.x, inc.y)
		want, incY := cas.want, inc.y
		switch {
		case inc.x < 0 && inc.y < 0:
			want = cas.wantRevXY
			incY = -inc.y
		case inc.x < 0:
			want = cas.wantRevX
		case inc.y < 0:
			want = cas.wantRevY
			incY = -inc.y
		}
		xg, yg := guardIncVector(test.x, xGdVal, inc.x, gdLn), guardIncVector(test.y, yGdVal, inc.y, gdLn)
		x, y := xg[gdLn:len(xg)-gdLn], yg[gdLn:len(yg)-gdLn]
		ag := guardVector(test.A, aGdVal, gdLn)
		a := ag[gdLn : len(ag)-gdLn]

		if trans {
			GemvT2(uintptr(test.m), uintptr(test.n), cas.alpha,
				a, lda, x, uintptr(inc.x),
				cas.beta, y, uintptr(inc.y))
		} else {
			GemvN(uintptr(test.m), uintptr(test.n), cas.alpha,
				a, lda, x, uintptr(inc.x),
				cas.beta, y, uintptr(inc.y))
		}
		for i := range want {
			if !within(y[i*incY], want[i]) {
				t.Errorf(msgVal, incPrefix, i, y[i*incY], want[i])
			}
		}

		checkValidIncGuard(t, xg, xGdVal, inc.x, gdLn)
		checkValidIncGuard(t, yg, yGdVal, inc.y, gdLn)
		if !isValidGuard(ag, aGdVal, gdLn) {
			t.Errorf(msgGuard, incPrefix, "a", ag[:gdLn], ag[len(ag)-gdLn:])
		}
		if !equalStrided(test.x, x, inc.x) {
			t.Errorf(msgReadOnly, incPrefix, "x")
		}
		if !equalStrided(test.A, a, 1) {
			t.Errorf(msgReadOnly, incPrefix, "a")
		}
	}
}
