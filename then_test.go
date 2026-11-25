package multiless

import "testing"

func TestThen(t *testing.T) {
	for _, c := range []struct{ i, want int }{{-5, -1}, {-1, -1}, {0, 0}, {1, 1}, {7, 1}} {
		got := New().Then(c.i)
		if o := got.OrderingInt(); o != c.want {
			t.Errorf("Then(%v).OrderingInt() = %v, want %v", c.i, o, c.want)
		}
		if old := New().Cmp(c.i); old != got {
			t.Errorf("Cmp(%v) = %+v, Then = %+v", c.i, old, got)
		}
	}
	// Decided computations are left alone; undecided ones continue.
	if got := New().Then(-1).Then(1).OrderingInt(); got != -1 {
		t.Errorf("decided computation was overridden: %v", got)
	}
	if got := New().Then(0).Then(1).OrderingInt(); got != 1 {
		t.Errorf("tie didn't fall through: %v", got)
	}
}
