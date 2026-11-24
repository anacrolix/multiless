//go:build go1.27

package multiless

import (
	"math"
	"testing"
)

func TestEagerOrderedMethod(t *testing.T) {
	cases := []struct {
		l, r int
		want int
	}{{1, 2, -1}, {2, 1, 1}, {3, 3, 0}}
	for _, c := range cases {
		if got := New().EagerOrdered(c.l, c.r).OrderingInt(); got != c.want {
			t.Errorf("EagerOrdered(%v, %v) = %v, want %v", c.l, c.r, got, c.want)
		}
		if got, old := New().EagerOrdered(c.l, c.r), EagerOrdered(New(), c.l, c.r); got != old {
			t.Errorf("method and function disagree for (%v, %v)", c.l, c.r)
		}
	}
	// An already-decided computation is left alone.
	if got := New().Int(1, 2).EagerOrdered("b", "a").OrderingInt(); got != -1 {
		t.Errorf("decided computation was overridden: %v", got)
	}
}

// The deprecated type-specific methods must agree with EagerOrdered, including on an
// already-decided computation and for NaN.
func TestTypeSpecificMatchEagerOrdered(t *testing.T) {
	decided := New().Int(1, 2)
	for _, start := range []Computation{New(), decided} {
		for _, p := range [][2]int{{1, 2}, {2, 1}, {3, 3}} {
			l, r := p[0], p[1]
			check := func(name string, old, new Computation) {
				t.Helper()
				if old != new {
					t.Errorf("%s(%v, %v) from %+v: %+v != EagerOrdered %+v", name, l, r, start, old, new)
				}
			}
			check("Int", start.Int(l, r), start.EagerOrdered(l, r))
			check("Int64", start.Int64(int64(l), int64(r)), start.EagerOrdered(int64(l), int64(r)))
			check("Uint32", start.Uint32(uint32(l), uint32(r)), start.EagerOrdered(uint32(l), uint32(r)))
			check("Uint64", start.Uint64(uint64(l), uint64(r)), start.EagerOrdered(uint64(l), uint64(r)))
			check("Uintptr", start.Uintptr(uintptr(l), uintptr(r)), start.EagerOrdered(uintptr(l), uintptr(r)))
			check("Float64", start.Float64(float64(l), float64(r)), start.EagerOrdered(float64(l), float64(r)))
		}
		nan := math.NaN()
		for _, p := range [][2]float64{{nan, 1}, {1, nan}, {nan, nan}} {
			if old, new := start.Float64(p[0], p[1]), start.EagerOrdered(p[0], p[1]); old != new {
				t.Errorf("Float64(%v, %v) from %+v: %+v != EagerOrdered %+v", p[0], p[1], start, old, new)
			}
		}
	}
}
