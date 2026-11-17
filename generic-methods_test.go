//go:build go1.27

package multiless

import "testing"

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
